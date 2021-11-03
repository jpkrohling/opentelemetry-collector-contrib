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

package metadata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/model/pdata"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/datasource"
)

const (
	// Value was generated using the same library. Intent is to detect that something changed in library implementation
	// in case we received different value here. For more details inspect tests where this value is used.
	expectedHashValue = "b519d020edc95bf"
)

func TestMetricsDataPoint_GroupingKey(t *testing.T) {
	dataPoint := metricsDataPointForTests()

	groupingKey := dataPoint.GroupingKey()

	assert.NotNil(t, groupingKey)
	assert.Equal(t, dataPoint.metricName, groupingKey.MetricName)
	assert.Equal(t, dataPoint.metricValue.Unit(), groupingKey.MetricUnit)
	assert.Equal(t, dataPoint.metricValue.DataType(), groupingKey.MetricDataType)
}

func TestMetricsDataPoint_ToItem(t *testing.T) {
	dataPoint := metricsDataPointForTests()

	item, err := dataPoint.ToItem()
	require.NoError(t, err)

	assert.Equal(t, expectedHashValue, item.SeriesKey)
	assert.Equal(t, dataPoint.timestamp, item.Timestamp)
}

func TestMetricsDataPoint_ToDataForHashing(t *testing.T) {
	dataPoint := metricsDataPointForTests()

	actual := dataPoint.toDataForHashing()

	assert.Equal(t, metricName, actual.MetricName)

	assertLabel(t, actual.Labels[0], projectIDLabelName, dataPoint.databaseID.ProjectID())
	assertLabel(t, actual.Labels[1], instanceIDLabelName, dataPoint.databaseID.InstanceID())
	assertLabel(t, actual.Labels[2], databaseLabelName, dataPoint.databaseID.DatabaseName())

	labelsIndex := 3
	for _, labelValue := range dataPoint.labelValues {
		assertLabel(t, actual.Labels[labelsIndex], labelValue.Name(), labelValue.Value())
		labelsIndex++
	}
}

func TestMetricsDataPoint_Hash(t *testing.T) {
	dataPoint := metricsDataPointForTests()

	hashValue, err := dataPoint.hash()
	require.NoError(t, err)

	assert.Equal(t, expectedHashValue, hashValue)
}

func TestMetricsDataPoint_CopyTo(t *testing.T) {
	timestamp := time.Now().UTC()
	labelValues := allPossibleLabelValues()
	metricValues := allPossibleMetricValues(metricDataType)
	databaseID := databaseID()

	for _, metricValue := range metricValues {
		dataPoint := pdata.NewNumberDataPoint()
		metricsDataPoint := &MetricsDataPoint{
			metricName:  metricName,
			timestamp:   timestamp,
			databaseID:  databaseID,
			labelValues: labelValues,
			metricValue: metricValue,
		}

		metricsDataPoint.CopyTo(dataPoint)

		assertMetricValue(t, metricValue, dataPoint)

		assert.Equal(t, pdata.NewTimestampFromTime(timestamp), dataPoint.Timestamp())
		// Adding +3 here because we'll always have 3 labels added for each metric: project_id, instance_id, database
		assert.Equal(t, 3+len(labelValues), dataPoint.Attributes().Len())

		attributesMap := dataPoint.Attributes()

		assertDefaultLabels(t, attributesMap, databaseID)
		assertNonDefaultLabels(t, attributesMap, labelValues)
	}
}

func allPossibleLabelValues() []LabelValue {
	strLabelValue := stringLabelValue{
		StringLabelValueMetadata: StringLabelValueMetadata{
			queryLabelValueMetadata: newQueryLabelValueMetadata("stringLabelName", "stringLabelColumnName"),
		},
		value: stringValue,
	}
	bLabelValue := boolLabelValue{
		BoolLabelValueMetadata: BoolLabelValueMetadata{
			queryLabelValueMetadata: newQueryLabelValueMetadata("boolLabelName", "boolLabelColumnName"),
		},
		value: boolValue,
	}
	i64LabelValue := int64LabelValue{
		Int64LabelValueMetadata: Int64LabelValueMetadata{
			queryLabelValueMetadata: newQueryLabelValueMetadata("int64LabelName", "int64LabelColumnName"),
		},
		value: int64Value,
	}
	strSliceLabelValue := stringSliceLabelValue{
		StringSliceLabelValueMetadata: StringSliceLabelValueMetadata{
			queryLabelValueMetadata: newQueryLabelValueMetadata("stringSliceLabelName", "stringSliceLabelColumnName"),
		},
		value: stringValue,
	}
	btSliceLabelValue := byteSliceLabelValue{
		ByteSliceLabelValueMetadata: ByteSliceLabelValueMetadata{
			queryLabelValueMetadata: newQueryLabelValueMetadata("byteSliceLabelName", "byteSliceLabelColumnName"),
		},
		value: stringValue,
	}

	return []LabelValue{
		strLabelValue,
		bLabelValue,
		i64LabelValue,
		strSliceLabelValue,
		btSliceLabelValue,
	}
}

func allPossibleMetricValues(metricDataType pdata.MetricDataType) []MetricValue {
	dataType := NewMetricDataType(metricDataType, pdata.MetricAggregationTemporalityDelta, true)
	return []MetricValue{
		int64MetricValue{
			Int64MetricValueMetadata: Int64MetricValueMetadata{
				queryMetricValueMetadata: newQueryMetricValueMetadata("int64MetricName",
					"int64MetricColumnName", dataType, metricUnit),
			},
			value: int64Value,
		},
		float64MetricValue{
			Float64MetricValueMetadata: Float64MetricValueMetadata{
				queryMetricValueMetadata: newQueryMetricValueMetadata("float64MetricName",
					"float64MetricColumnName", dataType, metricUnit),
			},
			value: float64Value,
		},
	}
}

func assertDefaultLabels(t *testing.T, attributesMap pdata.AttributeMap, databaseID *datasource.DatabaseID) {
	assertStringLabelValue(t, attributesMap, projectIDLabelName, databaseID.ProjectID())
	assertStringLabelValue(t, attributesMap, instanceIDLabelName, databaseID.InstanceID())
	assertStringLabelValue(t, attributesMap, databaseLabelName, databaseID.DatabaseName())
}

func assertNonDefaultLabels(t *testing.T, attributesMap pdata.AttributeMap, labelValues []LabelValue) {
	for _, labelValue := range labelValues {
		assertLabelValue(t, attributesMap, labelValue)
	}
}

func assertLabelValue(t *testing.T, attributesMap pdata.AttributeMap, labelValue LabelValue) {
	value, exists := attributesMap.Get(labelValue.Name())

	assert.True(t, exists)
	switch labelValue.(type) {
	case stringLabelValue, stringSliceLabelValue, byteSliceLabelValue:
		assert.Equal(t, labelValue.Value(), value.StringVal())
	case boolLabelValue:
		assert.Equal(t, labelValue.Value(), value.BoolVal())
	case int64LabelValue:
		assert.Equal(t, labelValue.Value(), value.IntVal())
	default:
		assert.Fail(t, "Unknown label value type received")
	}
}

func assertStringLabelValue(t *testing.T, attributesMap pdata.AttributeMap, labelName string, expectedValue interface{}) {
	value, exists := attributesMap.Get(labelName)

	assert.True(t, exists)
	assert.Equal(t, expectedValue, value.StringVal())
}

func assertMetricValue(t *testing.T, metricValue MetricValue, dataPoint pdata.NumberDataPoint) {
	switch metricValue.(type) {
	case int64MetricValue:
		assert.Equal(t, metricValue.Value(), dataPoint.IntVal())
	case float64MetricValue:
		assert.Equal(t, metricValue.Value(), dataPoint.DoubleVal())
	}
}

func assertLabel(t *testing.T, lbl label, expectedName string, expectedValue interface{}) {
	assert.Equal(t, expectedName, lbl.Name)
	assert.Equal(t, expectedValue, lbl.Value)
}

func metricsDataPointForTests() *MetricsDataPoint {
	timestamp := time.Now().UTC()
	labelValues := allPossibleLabelValues()
	databaseID := databaseID()

	return &MetricsDataPoint{
		metricName:  metricName,
		timestamp:   timestamp,
		databaseID:  databaseID,
		labelValues: labelValues,
		metricValue: allPossibleMetricValues(metricDataType)[0],
	}
}
