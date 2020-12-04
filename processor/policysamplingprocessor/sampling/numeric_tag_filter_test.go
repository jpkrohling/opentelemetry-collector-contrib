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
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

func TestNumericTagFilter(t *testing.T) {
	filter := NewNumericAttributeFilter(zap.NewNop(), "example", math.MinInt32, math.MaxInt32)

	cases := []struct {
		Desc     string
		Trace    pdata.ResourceSpans
		Decision Decision
	}{
		{
			Desc:     "nonmatching span attribute",
			Trace:    newTraceIntAttrs("non_matching", math.MinInt32),
			Decision: Undecided,
		},
		{
			Desc:     "span attribute with lower limit",
			Trace:    newTraceIntAttrs("example", math.MinInt32),
			Decision: Sample,
		},
		{
			Desc:     "span attribute with upper limit",
			Trace:    newTraceIntAttrs("example", math.MaxInt32),
			Decision: Sample,
		},
		{
			Desc:     "span attribute below min limit",
			Trace:    newTraceIntAttrs("example", math.MinInt32-1),
			Decision: Undecided,
		},
		{
			Desc:     "span attribute above max limit",
			Trace:    newTraceIntAttrs("example", math.MaxInt32+1),
			Decision: Undecided,
		},
	}

	for _, c := range cases {
		t.Run(c.Desc, func(t *testing.T) {
			decision, err := filter.Evaluate(c.Trace)
			assert.NoError(t, err)
			assert.Equal(t, decision, c.Decision)
		})
	}
}

func newTraceIntAttrs(spanAttrKey string, spanAttrValue int64) pdata.ResourceSpans {
	rs := pdata.NewResourceSpans()
	rs.InstrumentationLibrarySpans().Resize(1)

	ils := rs.InstrumentationLibrarySpans().At(0)
	ils.Spans().Resize(1)

	span := ils.Spans().At(0)
	span.SetTraceID(pdata.NewTraceID([16]byte{1, 2, 3, 4}))
	span.SetSpanID(pdata.NewSpanID([8]byte{2, 2, 3, 4}))

	attributes := make(map[string]pdata.AttributeValue)
	attributes[spanAttrKey] = pdata.NewAttributeValueInt(spanAttrValue)
	span.Attributes().InitFromMap(attributes)

	return rs
}
