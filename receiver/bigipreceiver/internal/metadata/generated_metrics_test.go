// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
		{
			name:        "filter_set_include",
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "filter_set_exclude",
			resAttrsSet: testDataSetAll,
			expectEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopSettings(receivertest.NopType)
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, tt.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipNodeAvailabilityDataPoint(ts, 1, AttributeAvailabilityStatusOffline)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipNodeConnectionCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipNodeDataTransmittedDataPoint(ts, 1, AttributeDirectionSent)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipNodeEnabledDataPoint(ts, 1, AttributeEnabledStatusDisabled)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipNodePacketCountDataPoint(ts, 1, AttributeDirectionSent)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipNodeRequestCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipNodeSessionCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolAvailabilityDataPoint(ts, 1, AttributeAvailabilityStatusOffline)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolConnectionCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolDataTransmittedDataPoint(ts, 1, AttributeDirectionSent)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolEnabledDataPoint(ts, 1, AttributeEnabledStatusDisabled)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolMemberCountDataPoint(ts, 1, AttributeActiveStatusActive)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolPacketCountDataPoint(ts, 1, AttributeDirectionSent)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolRequestCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolMemberAvailabilityDataPoint(ts, 1, AttributeAvailabilityStatusOffline)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolMemberConnectionCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolMemberDataTransmittedDataPoint(ts, 1, AttributeDirectionSent)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolMemberEnabledDataPoint(ts, 1, AttributeEnabledStatusDisabled)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolMemberPacketCountDataPoint(ts, 1, AttributeDirectionSent)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolMemberRequestCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipPoolMemberSessionCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipVirtualServerAvailabilityDataPoint(ts, 1, AttributeAvailabilityStatusOffline)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipVirtualServerConnectionCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipVirtualServerDataTransmittedDataPoint(ts, 1, AttributeDirectionSent)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipVirtualServerEnabledDataPoint(ts, 1, AttributeEnabledStatusDisabled)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipVirtualServerPacketCountDataPoint(ts, 1, AttributeDirectionSent)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordBigipVirtualServerRequestCountDataPoint(ts, 1)

			rb := mb.NewResourceBuilder()
			rb.SetBigipNodeIPAddress("bigip.node.ip_address-val")
			rb.SetBigipNodeName("bigip.node.name-val")
			rb.SetBigipPoolName("bigip.pool.name-val")
			rb.SetBigipPoolMemberIPAddress("bigip.pool_member.ip_address-val")
			rb.SetBigipPoolMemberName("bigip.pool_member.name-val")
			rb.SetBigipVirtualServerDestination("bigip.virtual_server.destination-val")
			rb.SetBigipVirtualServerName("bigip.virtual_server.name-val")
			res := rb.Emit()
			metrics := mb.Emit(WithResource(res))

			if tt.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if tt.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if tt.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "bigip.node.availability":
					assert.False(t, validatedMetrics["bigip.node.availability"], "Found a duplicate in the metrics slice: bigip.node.availability")
					validatedMetrics["bigip.node.availability"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Availability of the node.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("status")
					assert.True(t, ok)
					assert.EqualValues(t, "offline", attrVal.Str())
				case "bigip.node.connection.count":
					assert.False(t, validatedMetrics["bigip.node.connection.count"], "Found a duplicate in the metrics slice: bigip.node.connection.count")
					validatedMetrics["bigip.node.connection.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Current number of connections to the node.", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "bigip.node.data.transmitted":
					assert.False(t, validatedMetrics["bigip.node.data.transmitted"], "Found a duplicate in the metrics slice: bigip.node.data.transmitted")
					validatedMetrics["bigip.node.data.transmitted"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Amount of data transmitted to and from the node.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
				case "bigip.node.enabled":
					assert.False(t, validatedMetrics["bigip.node.enabled"], "Found a duplicate in the metrics slice: bigip.node.enabled")
					validatedMetrics["bigip.node.enabled"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Enabled state of of the node.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("status")
					assert.True(t, ok)
					assert.EqualValues(t, "disabled", attrVal.Str())
				case "bigip.node.packet.count":
					assert.False(t, validatedMetrics["bigip.node.packet.count"], "Found a duplicate in the metrics slice: bigip.node.packet.count")
					validatedMetrics["bigip.node.packet.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of packets transmitted to and from the node.", ms.At(i).Description())
					assert.Equal(t, "{packets}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
				case "bigip.node.request.count":
					assert.False(t, validatedMetrics["bigip.node.request.count"], "Found a duplicate in the metrics slice: bigip.node.request.count")
					validatedMetrics["bigip.node.request.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of requests to the node.", ms.At(i).Description())
					assert.Equal(t, "{requests}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "bigip.node.session.count":
					assert.False(t, validatedMetrics["bigip.node.session.count"], "Found a duplicate in the metrics slice: bigip.node.session.count")
					validatedMetrics["bigip.node.session.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Current number of sessions for the node.", ms.At(i).Description())
					assert.Equal(t, "{sessions}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "bigip.pool.availability":
					assert.False(t, validatedMetrics["bigip.pool.availability"], "Found a duplicate in the metrics slice: bigip.pool.availability")
					validatedMetrics["bigip.pool.availability"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Availability of the pool.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("status")
					assert.True(t, ok)
					assert.EqualValues(t, "offline", attrVal.Str())
				case "bigip.pool.connection.count":
					assert.False(t, validatedMetrics["bigip.pool.connection.count"], "Found a duplicate in the metrics slice: bigip.pool.connection.count")
					validatedMetrics["bigip.pool.connection.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Current number of connections to the pool.", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "bigip.pool.data.transmitted":
					assert.False(t, validatedMetrics["bigip.pool.data.transmitted"], "Found a duplicate in the metrics slice: bigip.pool.data.transmitted")
					validatedMetrics["bigip.pool.data.transmitted"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Amount of data transmitted to and from the pool.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
				case "bigip.pool.enabled":
					assert.False(t, validatedMetrics["bigip.pool.enabled"], "Found a duplicate in the metrics slice: bigip.pool.enabled")
					validatedMetrics["bigip.pool.enabled"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Enabled state of of the pool.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("status")
					assert.True(t, ok)
					assert.EqualValues(t, "disabled", attrVal.Str())
				case "bigip.pool.member.count":
					assert.False(t, validatedMetrics["bigip.pool.member.count"], "Found a duplicate in the metrics slice: bigip.pool.member.count")
					validatedMetrics["bigip.pool.member.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of pool members.", ms.At(i).Description())
					assert.Equal(t, "{members}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("status")
					assert.True(t, ok)
					assert.EqualValues(t, "active", attrVal.Str())
				case "bigip.pool.packet.count":
					assert.False(t, validatedMetrics["bigip.pool.packet.count"], "Found a duplicate in the metrics slice: bigip.pool.packet.count")
					validatedMetrics["bigip.pool.packet.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of packets transmitted to and from the pool.", ms.At(i).Description())
					assert.Equal(t, "{packets}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
				case "bigip.pool.request.count":
					assert.False(t, validatedMetrics["bigip.pool.request.count"], "Found a duplicate in the metrics slice: bigip.pool.request.count")
					validatedMetrics["bigip.pool.request.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of requests to the pool.", ms.At(i).Description())
					assert.Equal(t, "{requests}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "bigip.pool_member.availability":
					assert.False(t, validatedMetrics["bigip.pool_member.availability"], "Found a duplicate in the metrics slice: bigip.pool_member.availability")
					validatedMetrics["bigip.pool_member.availability"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Availability of the pool member.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("status")
					assert.True(t, ok)
					assert.EqualValues(t, "offline", attrVal.Str())
				case "bigip.pool_member.connection.count":
					assert.False(t, validatedMetrics["bigip.pool_member.connection.count"], "Found a duplicate in the metrics slice: bigip.pool_member.connection.count")
					validatedMetrics["bigip.pool_member.connection.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Current number of connections to the pool member.", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "bigip.pool_member.data.transmitted":
					assert.False(t, validatedMetrics["bigip.pool_member.data.transmitted"], "Found a duplicate in the metrics slice: bigip.pool_member.data.transmitted")
					validatedMetrics["bigip.pool_member.data.transmitted"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Amount of data transmitted to and from the pool member.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
				case "bigip.pool_member.enabled":
					assert.False(t, validatedMetrics["bigip.pool_member.enabled"], "Found a duplicate in the metrics slice: bigip.pool_member.enabled")
					validatedMetrics["bigip.pool_member.enabled"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Enabled state of of the pool member.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("status")
					assert.True(t, ok)
					assert.EqualValues(t, "disabled", attrVal.Str())
				case "bigip.pool_member.packet.count":
					assert.False(t, validatedMetrics["bigip.pool_member.packet.count"], "Found a duplicate in the metrics slice: bigip.pool_member.packet.count")
					validatedMetrics["bigip.pool_member.packet.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of packets transmitted to and from the pool member.", ms.At(i).Description())
					assert.Equal(t, "{packets}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
				case "bigip.pool_member.request.count":
					assert.False(t, validatedMetrics["bigip.pool_member.request.count"], "Found a duplicate in the metrics slice: bigip.pool_member.request.count")
					validatedMetrics["bigip.pool_member.request.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of requests to the pool member.", ms.At(i).Description())
					assert.Equal(t, "{requests}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "bigip.pool_member.session.count":
					assert.False(t, validatedMetrics["bigip.pool_member.session.count"], "Found a duplicate in the metrics slice: bigip.pool_member.session.count")
					validatedMetrics["bigip.pool_member.session.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Current number of sessions for the pool member.", ms.At(i).Description())
					assert.Equal(t, "{sessions}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "bigip.virtual_server.availability":
					assert.False(t, validatedMetrics["bigip.virtual_server.availability"], "Found a duplicate in the metrics slice: bigip.virtual_server.availability")
					validatedMetrics["bigip.virtual_server.availability"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Availability of the virtual server.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("status")
					assert.True(t, ok)
					assert.EqualValues(t, "offline", attrVal.Str())
				case "bigip.virtual_server.connection.count":
					assert.False(t, validatedMetrics["bigip.virtual_server.connection.count"], "Found a duplicate in the metrics slice: bigip.virtual_server.connection.count")
					validatedMetrics["bigip.virtual_server.connection.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Current number of connections to the virtual server.", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "bigip.virtual_server.data.transmitted":
					assert.False(t, validatedMetrics["bigip.virtual_server.data.transmitted"], "Found a duplicate in the metrics slice: bigip.virtual_server.data.transmitted")
					validatedMetrics["bigip.virtual_server.data.transmitted"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Amount of data transmitted to and from the virtual server.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
				case "bigip.virtual_server.enabled":
					assert.False(t, validatedMetrics["bigip.virtual_server.enabled"], "Found a duplicate in the metrics slice: bigip.virtual_server.enabled")
					validatedMetrics["bigip.virtual_server.enabled"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Enabled state of of the virtual server.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("status")
					assert.True(t, ok)
					assert.EqualValues(t, "disabled", attrVal.Str())
				case "bigip.virtual_server.packet.count":
					assert.False(t, validatedMetrics["bigip.virtual_server.packet.count"], "Found a duplicate in the metrics slice: bigip.virtual_server.packet.count")
					validatedMetrics["bigip.virtual_server.packet.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of packets transmitted to and from the virtual server.", ms.At(i).Description())
					assert.Equal(t, "{packets}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
				case "bigip.virtual_server.request.count":
					assert.False(t, validatedMetrics["bigip.virtual_server.request.count"], "Found a duplicate in the metrics slice: bigip.virtual_server.request.count")
					validatedMetrics["bigip.virtual_server.request.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of requests to the virtual server.", ms.At(i).Description())
					assert.Equal(t, "{requests}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}
