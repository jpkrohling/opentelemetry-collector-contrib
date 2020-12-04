// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package policysamplingprocessor

import (
	"context"
	"fmt"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenterror"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/policysamplingprocessor/sampling"
)

// Policy combines a sampling policy evaluator with the destinations to be
// used for that policy.
type Policy struct {
	// Name used to identify this policy instance.
	Name string
	// Evaluator that decides if a trace is sampled or not by this policy instance.
	Evaluator sampling.PolicyEvaluator
	// ctx used to carry metric tags of each policy.
	ctx context.Context
}

// processor handles the incoming trace data and uses the given sampling policy to sample traces.
type processor struct {
	ctx          context.Context
	logger       *zap.Logger
	nextConsumer consumer.Traces

	policies []*Policy
}

// newTraceProcessor returns a processor.TraceProcessor that will perform policy sampling according to the given
// configuration.
func newTraceProcessor(logger *zap.Logger, nextConsumer consumer.Traces, cfg Config) (*processor, error) {
	if nextConsumer == nil {
		return nil, componenterror.ErrNilNextConsumer
	}

	ctx := context.Background()
	var policies []*Policy
	for i := range cfg.PolicyCfgs {
		policyCfg := &cfg.PolicyCfgs[i]
		policyCtx, _ := tag.New(ctx, tag.Upsert(tag.MustNewKey("policy"), policyCfg.Name))

		eval, err := getPolicyEvaluator(logger, policyCfg)
		if err != nil {
			return nil, err
		}

		policy := &Policy{
			Name:      policyCfg.Name,
			Evaluator: eval,
			ctx:       policyCtx,
		}

		policies = append(policies, policy)
	}

	tsp := &processor{
		ctx:          ctx,
		nextConsumer: nextConsumer,
		logger:       logger,
		policies:     policies,
	}

	return tsp, nil
}

func getPolicyEvaluator(logger *zap.Logger, cfg *PolicyCfg) (sampling.PolicyEvaluator, error) {
	switch cfg.Type {
	case AlwaysSample:
		return sampling.NewAlwaysSample(logger), nil
	case NumericAttribute:
		nafCfg := cfg.NumericAttributeCfg
		return sampling.NewNumericAttributeFilter(logger, nafCfg.Key, nafCfg.MinValue, nafCfg.MaxValue), nil
	case StringAttribute:
		safCfg := cfg.StringAttributeCfg
		return sampling.NewStringAttributeFilter(logger, safCfg.Key, safCfg.Values), nil
	case RateLimiting:
		rlfCfg := cfg.RateLimitingCfg
		return sampling.NewRateLimiting(logger, rlfCfg.SpansPerSecond), nil
	default:
		return nil, fmt.Errorf("unknown sampling policy type %q", cfg.Type)
	}
}

// ConsumeTraceData is required by the SpanProcessor interface.
func (tsp *processor) ConsumeTraces(ctx context.Context, td pdata.Traces) error {
	resourceSpans := td.ResourceSpans()
	for i := 0; i < resourceSpans.Len(); i++ {
		tsp.processTraces(resourceSpans.At(i))
	}
	return nil
}

func (tsp *processor) processTraces(resourceSpans pdata.ResourceSpans) {
	finalDecision := sampling.Undecided

	// we evaluate all policies until we find one that tells us to drop the span
	for _, policy := range tsp.policies {
		startTime := time.Now()
		decision, err := policy.Evaluator.Evaluate(resourceSpans)
		duration := time.Since(startTime)
		if err != nil {
			tsp.logger.Warn("failed to evaluate policy for trace", zap.String("policy", policy.Name), zap.Error(err))
			policyCtx, _ := tag.New(policy.ctx, tag.Upsert(tag.MustNewKey("success"), "false"))
			stats.Record(policyCtx, mPolicyLatency.M(duration.Milliseconds()))
			continue
		}
		policyCtx, _ := tag.New(policy.ctx, tag.Upsert(tag.MustNewKey("success"), "true"))
		stats.Record(policyCtx, mPolicyLatency.M(duration.Milliseconds()))

		switch decision {
		case sampling.Drop:
			decisionCtx, _ := tag.New(policy.ctx, tag.Upsert(tag.MustNewKey("decision"), "dropped"))
			stats.Record(decisionCtx, mPolicyDecisions.M(1))
			finalDecision = sampling.Drop
		case sampling.Sample:
			finalDecision = sampling.Sample
		case sampling.Undecided:
			// there's nothing for us to do here at the moment, we might want to record a metric in the future if we see the need
		default:
			decisionCtx, _ := tag.New(policy.ctx, tag.Upsert(tag.MustNewKey("decision"), "unknown"))
			stats.Record(decisionCtx, mPolicyDecisions.M(1))
			tsp.logger.Warn("unexpected sampling decision", zap.String("policy", policy.Name), zap.Int("decision", int(decision)))
			return
		}
	}

	// at this point, we either got a "Sample" decision, or are still undecided
	if finalDecision == sampling.Sample {
		decisionCtx, _ := tag.New(context.Background(), tag.Upsert(tag.MustNewKey("decision"), "sampled"))
		stats.Record(decisionCtx, mPolicyDecisions.M(1))
		traceTd := pdata.NewTraces()
		traceTd.ResourceSpans().Append(resourceSpans)
		if err := tsp.nextConsumer.ConsumeTraces(decisionCtx, traceTd); err != nil {
			tsp.logger.Warn("error sending trace to the next consumer", zap.Error(err))
		}
	} else {
		// no definitive decisions were made -- drop the trace
		decisionCtx, _ := tag.New(context.Background(), tag.Upsert(tag.MustNewKey("decision"), "undecided"))
		stats.Record(decisionCtx, mPolicyDecisions.M(1))
	}
}

func (tsp *processor) GetCapabilities() component.ProcessorCapabilities {
	return component.ProcessorCapabilities{MutatesConsumedData: false}
}

// Start is invoked during service startup.
func (tsp *processor) Start(context.Context, component.Host) error {
	return nil
}

// Shutdown is invoked during service shutdown.
func (tsp *processor) Shutdown(context.Context) error {
	return nil
}
