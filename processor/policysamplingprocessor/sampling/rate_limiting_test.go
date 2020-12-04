// Copyright  The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sampling

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

func TestRateLimiter(t *testing.T) {
	var empty = map[string]pdata.AttributeValue{}

	trace := newTraceStringAttrs(empty, "example", "value")
	rateLimiter := NewRateLimiting(zap.NewNop(), 3)

	// Trace span count greater than spans per second
	trace.InstrumentationLibrarySpans().At(0).Spans().Resize(10)
	decision, err := rateLimiter.Evaluate(trace)
	assert.Nil(t, err)
	assert.Equal(t, decision, Drop)

	// Trace span count equal to spans per second
	trace.InstrumentationLibrarySpans().At(0).Spans().Resize(3)
	decision, err = rateLimiter.Evaluate(trace)
	assert.Nil(t, err)
	assert.Equal(t, decision, Drop)

	// Trace span count less than spans per second
	trace.InstrumentationLibrarySpans().At(0).Spans().Resize(2)
	decision, err = rateLimiter.Evaluate(trace)
	assert.Nil(t, err)
	assert.Equal(t, decision, Undecided)

	// Trace span count less than spans per second
	trace.InstrumentationLibrarySpans().At(0).Spans().Resize(0)
	decision, err = rateLimiter.Evaluate(trace)
	assert.Nil(t, err)
	assert.Equal(t, decision, Undecided)
}
