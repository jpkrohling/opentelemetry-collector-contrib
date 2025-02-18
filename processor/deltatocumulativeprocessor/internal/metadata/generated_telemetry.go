// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                 metric.Meter
	mu                                    sync.Mutex
	registrations                         []metric.Registration
	DeltatocumulativeDatapointsDropped    metric.Int64Counter
	DeltatocumulativeDatapointsLinear     metric.Int64Counter
	DeltatocumulativeDatapointsProcessed  metric.Int64Counter
	DeltatocumulativeGapsLength           metric.Int64Counter
	DeltatocumulativeStreamsEvicted       metric.Int64Counter
	DeltatocumulativeStreamsLimit         metric.Int64Gauge
	DeltatocumulativeStreamsMaxStale      metric.Int64Gauge
	DeltatocumulativeStreamsTracked       metric.Int64UpDownCounter
	DeltatocumulativeStreamsTrackedLinear metric.Int64ObservableUpDownCounter
}

// TelemetryBuilderOption applies changes to default builder.
type TelemetryBuilderOption interface {
	apply(*TelemetryBuilder)
}

type telemetryBuilderOptionFunc func(mb *TelemetryBuilder)

func (tbof telemetryBuilderOptionFunc) apply(mb *TelemetryBuilder) {
	tbof(mb)
}

// RegisterDeltatocumulativeStreamsTrackedLinearCallback sets callback for observable DeltatocumulativeStreamsTrackedLinear metric.
func (builder *TelemetryBuilder) RegisterDeltatocumulativeStreamsTrackedLinearCallback(cb metric.Int64Callback) error {
	reg, err := builder.meter.RegisterCallback(func(ctx context.Context, o metric.Observer) error {
		cb(ctx, &observerInt64{inst: builder.DeltatocumulativeStreamsTrackedLinear, obs: o})
		return nil
	}, builder.DeltatocumulativeStreamsTrackedLinear)
	if err != nil {
		return err
	}
	builder.mu.Lock()
	defer builder.mu.Unlock()
	builder.registrations = append(builder.registrations, reg)
	return nil
}

type observerInt64 struct {
	embedded.Int64Observer
	inst metric.Int64Observable
	obs  metric.Observer
}

func (oi *observerInt64) Observe(value int64, opts ...metric.ObserveOption) {
	oi.obs.ObserveInt64(oi.inst, value, opts...)
}

// Shutdown unregister all registered callbacks for async instruments.
func (builder *TelemetryBuilder) Shutdown() {
	builder.mu.Lock()
	defer builder.mu.Unlock()
	for _, reg := range builder.registrations {
		reg.Unregister()
	}
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...TelemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{}
	for _, op := range options {
		op.apply(&builder)
	}
	builder.meter = Meter(settings)
	var err, errs error
	builder.DeltatocumulativeDatapointsDropped, err = builder.meter.Int64Counter(
		"otelcol_deltatocumulative.datapoints.dropped",
		metric.WithDescription("number of datapoints dropped due to given 'reason'"),
		metric.WithUnit("{datapoint}"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeDatapointsLinear, err = builder.meter.Int64Counter(
		"otelcol_deltatocumulative.datapoints.linear",
		metric.WithDescription("total number of datapoints processed. may have 'error' attribute, if processing failed"),
		metric.WithUnit("{datapoint}"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeDatapointsProcessed, err = builder.meter.Int64Counter(
		"otelcol_deltatocumulative.datapoints.processed",
		metric.WithDescription("number of datapoints processed"),
		metric.WithUnit("{datapoint}"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeGapsLength, err = builder.meter.Int64Counter(
		"otelcol_deltatocumulative.gaps.length",
		metric.WithDescription("total duration where data was expected but not received"),
		metric.WithUnit("s"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeStreamsEvicted, err = builder.meter.Int64Counter(
		"otelcol_deltatocumulative.streams.evicted",
		metric.WithDescription("number of streams evicted"),
		metric.WithUnit("{stream}"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeStreamsLimit, err = builder.meter.Int64Gauge(
		"otelcol_deltatocumulative.streams.limit",
		metric.WithDescription("upper limit of tracked streams"),
		metric.WithUnit("{stream}"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeStreamsMaxStale, err = builder.meter.Int64Gauge(
		"otelcol_deltatocumulative.streams.max_stale",
		metric.WithDescription("duration after which streams inactive streams are dropped"),
		metric.WithUnit("s"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeStreamsTracked, err = builder.meter.Int64UpDownCounter(
		"otelcol_deltatocumulative.streams.tracked",
		metric.WithDescription("number of streams tracked"),
		metric.WithUnit("{dps}"),
	)
	errs = errors.Join(errs, err)
	builder.DeltatocumulativeStreamsTrackedLinear, err = builder.meter.Int64ObservableUpDownCounter(
		"otelcol_deltatocumulative.streams.tracked.linear",
		metric.WithDescription("number of streams tracked"),
		metric.WithUnit("{dps}"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
