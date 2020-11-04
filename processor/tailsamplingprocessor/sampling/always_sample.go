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

package sampling

import (
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

type alwaysSample struct {
	logger *zap.Logger
}

var _ PolicyEvaluator = (*alwaysSample)(nil)

// NewAlwaysSample creates a policy evaluator the samples all traces.
func NewAlwaysSample(logger *zap.Logger) PolicyEvaluator {
	return &alwaysSample{
		logger: logger,
	}
}

// OnLateArrivingSpans notifies the evaluator that the given list of spans arrived
// after the sampling decision was already taken for the trace.
// This gives the evaluator a chance to log any message/metrics and/or update any
// related internal state.
func (as *alwaysSample) OnLateArrivingSpans(Decision, []*pdata.Span) error {
	as.logger.Debug("Triggering action for late arriving spans in always-sample filter")
	return nil
}

// EvaluateSecondChance looks at the trace again and if it can/cannot be fit, returns a SamplingDecision
func (as *alwaysSample) EvaluateSecondChance(_ pdata.TraceID, trace *TraceData) (Decision, error) {
	return NotSampled, nil
}

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (as *alwaysSample) Evaluate(pdata.TraceID, *TraceData) (Decision, error) {
	as.logger.Debug("Evaluating spans in always-sample filter")
	return Sampled, nil
}
