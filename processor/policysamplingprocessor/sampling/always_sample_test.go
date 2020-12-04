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

func TestEvaluate_AlwaysSample(t *testing.T) {
	filter := NewAlwaysSample(zap.NewNop())
	trace := pdata.NewResourceSpans()
	trace.InstrumentationLibrarySpans().Resize(1)
	trace.InstrumentationLibrarySpans().At(0).Spans().Resize(1)
	trace.InstrumentationLibrarySpans().At(0).Spans().At(0).SetTraceID(pdata.NewTraceID([16]byte{1, 2, 3, 4}))
	decision, err := filter.Evaluate(trace)
	assert.Nil(t, err)
	assert.Equal(t, decision, Sample)
}
