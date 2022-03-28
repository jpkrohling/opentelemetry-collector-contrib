// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/model/pdata"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for riakreceiver metrics.
type MetricsSettings struct {
	RiakMemoryLimit              MetricSettings `mapstructure:"riak.memory.limit"`
	RiakNodeOperationCount       MetricSettings `mapstructure:"riak.node.operation.count"`
	RiakNodeOperationTimeMean    MetricSettings `mapstructure:"riak.node.operation.time.mean"`
	RiakNodeReadRepairCount      MetricSettings `mapstructure:"riak.node.read_repair.count"`
	RiakVnodeIndexOperationCount MetricSettings `mapstructure:"riak.vnode.index.operation.count"`
	RiakVnodeOperationCount      MetricSettings `mapstructure:"riak.vnode.operation.count"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		RiakMemoryLimit: MetricSettings{
			Enabled: true,
		},
		RiakNodeOperationCount: MetricSettings{
			Enabled: true,
		},
		RiakNodeOperationTimeMean: MetricSettings{
			Enabled: true,
		},
		RiakNodeReadRepairCount: MetricSettings{
			Enabled: true,
		},
		RiakVnodeIndexOperationCount: MetricSettings{
			Enabled: true,
		},
		RiakVnodeOperationCount: MetricSettings{
			Enabled: true,
		},
	}
}

type metricRiakMemoryLimit struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.memory.limit metric with initial data.
func (m *metricRiakMemoryLimit) init() {
	m.data.SetName("riak.memory.limit")
	m.data.SetDescription("The amount of memory allocated to the node.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricRiakMemoryLimit) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakMemoryLimit) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakMemoryLimit) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakMemoryLimit(settings MetricSettings) metricRiakMemoryLimit {
	m := metricRiakMemoryLimit{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricRiakNodeOperationCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.node.operation.count metric with initial data.
func (m *metricRiakNodeOperationCount) init() {
	m.data.SetName("riak.node.operation.count")
	m.data.SetDescription("The number of operations performed by the node.")
	m.data.SetUnit("{operation}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricRiakNodeOperationCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, requestAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Request, pdata.NewAttributeValueString(requestAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakNodeOperationCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakNodeOperationCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakNodeOperationCount(settings MetricSettings) metricRiakNodeOperationCount {
	m := metricRiakNodeOperationCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricRiakNodeOperationTimeMean struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.node.operation.time.mean metric with initial data.
func (m *metricRiakNodeOperationTimeMean) init() {
	m.data.SetName("riak.node.operation.time.mean")
	m.data.SetDescription("The mean time between request and response for operations performed by the node over the last minute.")
	m.data.SetUnit("us")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricRiakNodeOperationTimeMean) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, requestAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Request, pdata.NewAttributeValueString(requestAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakNodeOperationTimeMean) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakNodeOperationTimeMean) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakNodeOperationTimeMean(settings MetricSettings) metricRiakNodeOperationTimeMean {
	m := metricRiakNodeOperationTimeMean{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricRiakNodeReadRepairCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.node.read_repair.count metric with initial data.
func (m *metricRiakNodeReadRepairCount) init() {
	m.data.SetName("riak.node.read_repair.count")
	m.data.SetDescription("The number of read repairs performed by the node.")
	m.data.SetUnit("{read_repair}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
}

func (m *metricRiakNodeReadRepairCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakNodeReadRepairCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakNodeReadRepairCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakNodeReadRepairCount(settings MetricSettings) metricRiakNodeReadRepairCount {
	m := metricRiakNodeReadRepairCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricRiakVnodeIndexOperationCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.vnode.index.operation.count metric with initial data.
func (m *metricRiakVnodeIndexOperationCount) init() {
	m.data.SetName("riak.vnode.index.operation.count")
	m.data.SetDescription("The number of index operations performed by vnodes on the node.")
	m.data.SetUnit("{operation}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricRiakVnodeIndexOperationCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, operationAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Operation, pdata.NewAttributeValueString(operationAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakVnodeIndexOperationCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakVnodeIndexOperationCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakVnodeIndexOperationCount(settings MetricSettings) metricRiakVnodeIndexOperationCount {
	m := metricRiakVnodeIndexOperationCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricRiakVnodeOperationCount struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.vnode.operation.count metric with initial data.
func (m *metricRiakVnodeOperationCount) init() {
	m.data.SetName("riak.vnode.operation.count")
	m.data.SetDescription("The number of operations performed by vnodes on the node.")
	m.data.SetUnit("{operation}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricRiakVnodeOperationCount) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, requestAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Request, pdata.NewAttributeValueString(requestAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakVnodeOperationCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakVnodeOperationCount) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakVnodeOperationCount(settings MetricSettings) metricRiakVnodeOperationCount {
	m := metricRiakVnodeOperationCount{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                          pdata.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity                    int             // maximum observed number of metrics per resource.
	resourceCapacity                   int             // maximum observed number of resource attributes.
	metricsBuffer                      pdata.Metrics   // accumulates metrics data before emitting.
	metricRiakMemoryLimit              metricRiakMemoryLimit
	metricRiakNodeOperationCount       metricRiakNodeOperationCount
	metricRiakNodeOperationTimeMean    metricRiakNodeOperationTimeMean
	metricRiakNodeReadRepairCount      metricRiakNodeReadRepairCount
	metricRiakVnodeIndexOperationCount metricRiakVnodeIndexOperationCount
	metricRiakVnodeOperationCount      metricRiakVnodeOperationCount
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pdata.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                          pdata.NewTimestampFromTime(time.Now()),
		metricsBuffer:                      pdata.NewMetrics(),
		metricRiakMemoryLimit:              newMetricRiakMemoryLimit(settings.RiakMemoryLimit),
		metricRiakNodeOperationCount:       newMetricRiakNodeOperationCount(settings.RiakNodeOperationCount),
		metricRiakNodeOperationTimeMean:    newMetricRiakNodeOperationTimeMean(settings.RiakNodeOperationTimeMean),
		metricRiakNodeReadRepairCount:      newMetricRiakNodeReadRepairCount(settings.RiakNodeReadRepairCount),
		metricRiakVnodeIndexOperationCount: newMetricRiakVnodeIndexOperationCount(settings.RiakVnodeIndexOperationCount),
		metricRiakVnodeOperationCount:      newMetricRiakVnodeOperationCount(settings.RiakVnodeOperationCount),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pdata.ResourceMetrics) {
	if mb.metricsCapacity < rm.InstrumentationLibraryMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.InstrumentationLibraryMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceOption applies changes to provided resource.
type ResourceOption func(pdata.Resource)

// WithRiakNodeName sets provided value as "riak.node.name" attribute for current resource.
func WithRiakNodeName(val string) ResourceOption {
	return func(r pdata.Resource) {
		r.Attributes().UpsertString("riak.node.name", val)
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead. Resource attributes should be provided as ResourceOption arguments.
func (mb *MetricsBuilder) EmitForResource(ro ...ResourceOption) {
	rm := pdata.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	for _, op := range ro {
		op(rm.Resource())
	}
	ils := rm.InstrumentationLibraryMetrics().AppendEmpty()
	ils.InstrumentationLibrary().SetName("otelcol/riakreceiver")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricRiakMemoryLimit.emit(ils.Metrics())
	mb.metricRiakNodeOperationCount.emit(ils.Metrics())
	mb.metricRiakNodeOperationTimeMean.emit(ils.Metrics())
	mb.metricRiakNodeReadRepairCount.emit(ils.Metrics())
	mb.metricRiakVnodeIndexOperationCount.emit(ils.Metrics())
	mb.metricRiakVnodeOperationCount.emit(ils.Metrics())
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(ro ...ResourceOption) pdata.Metrics {
	mb.EmitForResource(ro...)
	metrics := pdata.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordRiakMemoryLimitDataPoint adds a data point to riak.memory.limit metric.
func (mb *MetricsBuilder) RecordRiakMemoryLimitDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricRiakMemoryLimit.recordDataPoint(mb.startTime, ts, val)
}

// RecordRiakNodeOperationCountDataPoint adds a data point to riak.node.operation.count metric.
func (mb *MetricsBuilder) RecordRiakNodeOperationCountDataPoint(ts pdata.Timestamp, val int64, requestAttributeValue string) {
	mb.metricRiakNodeOperationCount.recordDataPoint(mb.startTime, ts, val, requestAttributeValue)
}

// RecordRiakNodeOperationTimeMeanDataPoint adds a data point to riak.node.operation.time.mean metric.
func (mb *MetricsBuilder) RecordRiakNodeOperationTimeMeanDataPoint(ts pdata.Timestamp, val int64, requestAttributeValue string) {
	mb.metricRiakNodeOperationTimeMean.recordDataPoint(mb.startTime, ts, val, requestAttributeValue)
}

// RecordRiakNodeReadRepairCountDataPoint adds a data point to riak.node.read_repair.count metric.
func (mb *MetricsBuilder) RecordRiakNodeReadRepairCountDataPoint(ts pdata.Timestamp, val int64) {
	mb.metricRiakNodeReadRepairCount.recordDataPoint(mb.startTime, ts, val)
}

// RecordRiakVnodeIndexOperationCountDataPoint adds a data point to riak.vnode.index.operation.count metric.
func (mb *MetricsBuilder) RecordRiakVnodeIndexOperationCountDataPoint(ts pdata.Timestamp, val int64, operationAttributeValue string) {
	mb.metricRiakVnodeIndexOperationCount.recordDataPoint(mb.startTime, ts, val, operationAttributeValue)
}

// RecordRiakVnodeOperationCountDataPoint adds a data point to riak.vnode.operation.count metric.
func (mb *MetricsBuilder) RecordRiakVnodeOperationCountDataPoint(ts pdata.Timestamp, val int64, requestAttributeValue string) {
	mb.metricRiakVnodeOperationCount.recordDataPoint(mb.startTime, ts, val, requestAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Operation (The operation type for index operations.)
	Operation string
	// Request (The request operation type.)
	Request string
}{
	"operation",
	"request",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeOperation are the possible values that the attribute "operation" can have.
var AttributeOperation = struct {
	Read   string
	Write  string
	Delete string
}{
	"read",
	"write",
	"delete",
}

// AttributeRequest are the possible values that the attribute "request" can have.
var AttributeRequest = struct {
	Put string
	Get string
}{
	"put",
	"get",
}
