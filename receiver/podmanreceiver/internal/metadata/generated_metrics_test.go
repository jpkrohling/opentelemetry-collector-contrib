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
			mb.RecordContainerBlockioIoServiceBytesRecursiveReadDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordContainerBlockioIoServiceBytesRecursiveWriteDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordContainerCPUPercentDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordContainerCPUUsagePercpuDataPoint(ts, 1, "core-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordContainerCPUUsageSystemDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordContainerCPUUsageTotalDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordContainerMemoryPercentDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordContainerMemoryUsageLimitDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordContainerMemoryUsageTotalDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordContainerNetworkIoUsageRxBytesDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordContainerNetworkIoUsageTxBytesDataPoint(ts, 1)

			rb := mb.NewResourceBuilder()
			rb.SetContainerID("container.id-val")
			rb.SetContainerImageName("container.image.name-val")
			rb.SetContainerName("container.name-val")
			rb.SetContainerRuntime("container.runtime-val")
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
				case "container.blockio.io_service_bytes_recursive.read":
					assert.False(t, validatedMetrics["container.blockio.io_service_bytes_recursive.read"], "Found a duplicate in the metrics slice: container.blockio.io_service_bytes_recursive.read")
					validatedMetrics["container.blockio.io_service_bytes_recursive.read"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of bytes transferred from the disk by the container", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "container.blockio.io_service_bytes_recursive.write":
					assert.False(t, validatedMetrics["container.blockio.io_service_bytes_recursive.write"], "Found a duplicate in the metrics slice: container.blockio.io_service_bytes_recursive.write")
					validatedMetrics["container.blockio.io_service_bytes_recursive.write"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of bytes transferred to the disk by the container", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "container.cpu.percent":
					assert.False(t, validatedMetrics["container.cpu.percent"], "Found a duplicate in the metrics slice: container.cpu.percent")
					validatedMetrics["container.cpu.percent"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Percent of CPU used by the container.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "container.cpu.usage.percpu":
					assert.False(t, validatedMetrics["container.cpu.usage.percpu"], "Found a duplicate in the metrics slice: container.cpu.usage.percpu")
					validatedMetrics["container.cpu.usage.percpu"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total CPU time consumed per CPU-core.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("core")
					assert.True(t, ok)
					assert.EqualValues(t, "core-val", attrVal.Str())
				case "container.cpu.usage.system":
					assert.False(t, validatedMetrics["container.cpu.usage.system"], "Found a duplicate in the metrics slice: container.cpu.usage.system")
					validatedMetrics["container.cpu.usage.system"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "System CPU usage.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "container.cpu.usage.total":
					assert.False(t, validatedMetrics["container.cpu.usage.total"], "Found a duplicate in the metrics slice: container.cpu.usage.total")
					validatedMetrics["container.cpu.usage.total"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total CPU time consumed.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "container.memory.percent":
					assert.False(t, validatedMetrics["container.memory.percent"], "Found a duplicate in the metrics slice: container.memory.percent")
					validatedMetrics["container.memory.percent"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Percentage of memory used.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "container.memory.usage.limit":
					assert.False(t, validatedMetrics["container.memory.usage.limit"], "Found a duplicate in the metrics slice: container.memory.usage.limit")
					validatedMetrics["container.memory.usage.limit"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Memory limit of the container.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "container.memory.usage.total":
					assert.False(t, validatedMetrics["container.memory.usage.total"], "Found a duplicate in the metrics slice: container.memory.usage.total")
					validatedMetrics["container.memory.usage.total"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Memory usage of the container.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "container.network.io.usage.rx_bytes":
					assert.False(t, validatedMetrics["container.network.io.usage.rx_bytes"], "Found a duplicate in the metrics slice: container.network.io.usage.rx_bytes")
					validatedMetrics["container.network.io.usage.rx_bytes"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes received by the container.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "container.network.io.usage.tx_bytes":
					assert.False(t, validatedMetrics["container.network.io.usage.tx_bytes"], "Found a duplicate in the metrics slice: container.network.io.usage.tx_bytes")
					validatedMetrics["container.network.io.usage.tx_bytes"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes sent by the container.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
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
