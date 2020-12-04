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
)

// Decision gives the status of sampling decision.
type Decision int32

const (
	// Undecided represents a decision that isn't definitive
	Undecided Decision = iota
	// Sampled represents the decision to sample the trace
	Sample
	// Drop represents the decision to not sample the trace
	Drop
)

// PolicyEvaluator implements a policy-based sampling policy evaluator,
// which makes a sampling decision for a given trace when requested.
type PolicyEvaluator interface {
	// Evaluate looks at the given trace and returns a corresponding SamplingDecision.
	Evaluate(trace pdata.ResourceSpans) (Decision, error)
}
