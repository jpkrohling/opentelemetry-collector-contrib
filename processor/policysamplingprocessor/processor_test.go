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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenterror"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/policysamplingprocessor/sampling"
)

func TestNilNextConsumer(t *testing.T) {
	// test
	processor, err := newTraceProcessor(nil, nil, Config{})

	// verify
	assert.Equal(t, componenterror.ErrNilNextConsumer, err)
	assert.Nil(t, processor)
}

func TestPolicyEvaluator(t *testing.T) {
	// we can't test the evaluator directly, so, we cross test the individual evaluators with the consume traces
	// we have only one case per policy, as each policy should be tested individually on their own tests
	testCases := []struct {
		desc  string
		cfg   PolicyCfg
		trace pdata.ResourceSpans
		fail  bool // whether we expect an error from the GetPolicyEvaluator
		drop  bool
	}{
		{
			desc: "always sample",
			cfg: PolicyCfg{
				Type: AlwaysSample,
			},
			trace: simpleTrace(),
		},
		{
			desc: "numeric attribute",
			cfg: PolicyCfg{
				Type: NumericAttribute,
				NumericAttributeCfg: NumericAttributeCfg{
					Key:      "my-key",
					MinValue: 10,
					MaxValue: 20,
				},
			},
			trace: func() pdata.ResourceSpans {
				trace := simpleTrace()
				trace.InstrumentationLibrarySpans().At(0).Spans().At(0).Attributes().InsertInt("my-key", 15)
				return trace
			}(),
		},
		{
			desc: "string attribute",
			cfg: PolicyCfg{
				Type: StringAttribute,
				StringAttributeCfg: StringAttributeCfg{
					Key:    "my-key",
					Values: []string{"some-value"},
				},
			},
			trace: func() pdata.ResourceSpans {
				trace := simpleTrace()
				trace.InstrumentationLibrarySpans().At(0).Spans().At(0).Attributes().InsertString("my-key", "some-value")
				return trace
			}(),
		},
		{
			desc: "rate limiting",
			cfg: PolicyCfg{
				Type: RateLimiting,
				RateLimitingCfg: RateLimitingCfg{
					SpansPerSecond: 1000,
				},
			},
			trace: simpleTrace(),
			drop:  true,
		},
		{
			desc: "unknown",
			cfg: PolicyCfg{
				Type: PolicyType("unknown"),
			},
			trace: simpleTrace(),
			fail:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// prepare
			next := &consumertest.TracesSink{}
			cfg := Config{
				PolicyCfgs: []PolicyCfg{tC.cfg},
			}

			// test
			processor, err := newTraceProcessor(zap.NewNop(), next, cfg)
			if tC.fail {
				// verify
				assert.Error(t, err)
				return
			}
			require.NotNil(t, processor)
			require.NoError(t, err)

			traces := pdata.NewTraces()
			traces.ResourceSpans().Append(tC.trace)

			err = processor.ConsumeTraces(context.Background(), traces)

			// verify
			assert.NoError(t, err)

			if tC.drop {
				assert.Empty(t, next.AllTraces())
			} else {
				assert.Len(t, next.AllTraces(), 1)
			}
		})
	}
}

func TestProcessTraces(t *testing.T) {
	testCases := []struct {
		desc     string
		rs       pdata.ResourceSpans
		next     *mockNextConsumer
		policies []*Policy
		drop     bool
	}{
		{
			desc: "err on evaluate",
			rs:   simpleTrace(),
			next: &mockNextConsumer{},
			policies: []*Policy{{
				ctx: context.Background(),
				Evaluator: &mockEvaluator{
					evaluator: func(trace pdata.ResourceSpans) (sampling.Decision, error) {
						return sampling.Undecided, errors.New("some expected error")
					},
				},
			}},
			drop: true, // we don't fail the processing of the trace when a policy fails
		},
		{
			desc: "err on next consumer",
			rs:   simpleTrace(),
			next: &mockNextConsumer{
				failOnConsume: true,
			},
			policies: []*Policy{{
				ctx:       context.Background(),
				Evaluator: sampling.NewAlwaysSample(zap.NewNop()),
			}},
		},
		{
			desc: "sample trace",
			rs:   simpleTrace(),
			next: &mockNextConsumer{},
			policies: []*Policy{{
				ctx:       context.Background(),
				Evaluator: sampling.NewAlwaysSample(zap.NewNop()),
			}},
		},
		{
			desc: "drop trace",
			rs:   simpleTrace(),
			next: &mockNextConsumer{},
			policies: []*Policy{{
				ctx: context.Background(),
				Evaluator: &mockEvaluator{
					evaluator: func(trace pdata.ResourceSpans) (sampling.Decision, error) {
						return sampling.Drop, nil
					},
				},
			}},
			drop: true,
		},
		{
			desc: "undecided",
			rs:   simpleTrace(),
			next: &mockNextConsumer{},
			policies: []*Policy{{
				ctx: context.Background(),
				Evaluator: &mockEvaluator{
					evaluator: func(trace pdata.ResourceSpans) (sampling.Decision, error) {
						return sampling.Undecided, nil
					},
				},
			}},
			drop: true,
		},
		{
			desc: "invalid decision",
			rs:   simpleTrace(),
			next: &mockNextConsumer{},
			policies: []*Policy{{
				ctx: context.Background(),
				Evaluator: &mockEvaluator{
					evaluator: func(trace pdata.ResourceSpans) (sampling.Decision, error) {
						return sampling.Decision(-1), nil
					},
				},
			}},
			drop: true,
		},
		{
			desc: "sample and drop",
			rs:   simpleTrace(),
			next: &mockNextConsumer{},
			policies: []*Policy{
				{ // this first policy returns a sample decision
					ctx: context.Background(),
					Evaluator: &mockEvaluator{
						evaluator: func(trace pdata.ResourceSpans) (sampling.Decision, error) {
							return sampling.Sample, nil
						},
					},
				},
				{ // this second policy returns a drop decision, making the trace be dropped
					ctx: context.Background(),
					Evaluator: &mockEvaluator{
						evaluator: func(trace pdata.ResourceSpans) (sampling.Decision, error) {
							return sampling.Drop, nil
						},
					},
				},
			},
			drop: true,
		},
		{
			desc: "sample and undecided",
			rs:   simpleTrace(),
			next: &mockNextConsumer{},
			policies: []*Policy{
				{ // this first policy returns a sample decision
					ctx: context.Background(),
					Evaluator: &mockEvaluator{
						evaluator: func(trace pdata.ResourceSpans) (sampling.Decision, error) {
							return sampling.Sample, nil
						},
					},
				},
				{ // this second policy returns an undecided decision, making the trace be sampled
					ctx: context.Background(),
					Evaluator: &mockEvaluator{
						evaluator: func(trace pdata.ResourceSpans) (sampling.Decision, error) {
							return sampling.Undecided, nil
						},
					},
				},
			},
		},
		{
			desc: "no decisions made",
			rs:   simpleTrace(),
			next: &mockNextConsumer{},
			drop: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// prepare
			cfg := Config{
				PolicyCfgs: []PolicyCfg{},
			}

			processor, err := newTraceProcessor(zap.NewNop(), tC.next, cfg)
			require.NoError(t, err)
			require.NotNil(t, processor)
			processor.policies = tC.policies

			// test
			processor.processTraces(tC.rs)

			// verify
			if tC.drop {
				assert.Empty(t, tC.next.AllTraces())
			} else {
				assert.Len(t, tC.next.AllTraces(), 1)
			}
		})
	}
}

func TestCapabilities(t *testing.T) {
	// prepare
	processor, err := newTraceProcessor(nil, consumertest.NewNop(), Config{})
	require.NotNil(t, processor)
	require.NoError(t, err)

	// test
	caps := processor.GetCapabilities()

	// verify
	assert.False(t, caps.MutatesConsumedData)
}

func TestStart(t *testing.T) {
	// prepare
	processor, err := newTraceProcessor(nil, consumertest.NewNop(), Config{})
	require.NotNil(t, processor)
	require.NoError(t, err)

	// test
	err = processor.Start(context.Background(), componenttest.NewNopHost())

	// verify
	assert.NoError(t, err)
}

func TestShutdown(t *testing.T) {
	// prepare
	processor, err := newTraceProcessor(nil, consumertest.NewNop(), Config{})
	require.NotNil(t, processor)
	require.NoError(t, err)

	// test
	err = processor.Shutdown(context.Background())

	// verify
	assert.NoError(t, err)
}

func simpleTrace() pdata.ResourceSpans {
	rs := pdata.NewResourceSpans()
	rs.InstrumentationLibrarySpans().Resize(1)
	rs.InstrumentationLibrarySpans().At(0).Spans().Resize(1)
	rs.InstrumentationLibrarySpans().At(0).Spans().At(0).SetTraceID(pdata.NewTraceID([16]byte{0, 1, 2, 3}))
	return rs
}

type mockNextConsumer struct {
	consumertest.TracesSink
	failOnConsume bool
}

func (m *mockNextConsumer) ConsumeTraces(ctx context.Context, td pdata.Traces) error {
	m.TracesSink.ConsumeTraces(ctx, td) // record the trace in the mock
	if m.failOnConsume {
		return errors.New("expected error on next consumer")
	}
	return nil
}

type mockEvaluator struct {
	evaluator func(trace pdata.ResourceSpans) (sampling.Decision, error)
}

func (m *mockEvaluator) Evaluate(trace pdata.ResourceSpans) (sampling.Decision, error) {
	if m.evaluator != nil {
		return m.evaluator(trace)
	}
	return sampling.Drop, nil
}
