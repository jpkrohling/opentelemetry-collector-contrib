// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"
)

func NewSettings(tt *componenttest.Telemetry) exporter.Settings {
	set := exportertest.NewNopSettingsWithType(exportertest.NopType)
	set.ID = component.NewID(component.MustNewType("loadbalancing"))
	set.TelemetrySettings = tt.NewTelemetrySettings()
	return set
}

func AssertEqualLoadbalancerBackendLatency(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.HistogramDataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_loadbalancer_backend_latency",
		Description: "Response latency in ms for the backends.",
		Unit:        "ms",
		Data: metricdata.Histogram[int64]{
			Temporality: metricdata.CumulativeTemporality,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_loadbalancer_backend_latency")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualLoadbalancerBackendOutcome(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_loadbalancer_backend_outcome",
		Description: "Number of successes and failures for each endpoint.",
		Unit:        "{outcomes}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_loadbalancer_backend_outcome")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualLoadbalancerNumBackendUpdates(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_loadbalancer_num_backend_updates",
		Description: "Number of times the list of backends was updated.",
		Unit:        "{updates}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_loadbalancer_num_backend_updates")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualLoadbalancerNumBackends(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_loadbalancer_num_backends",
		Description: "Current number of backends in use.",
		Unit:        "{backends}",
		Data: metricdata.Gauge[int64]{
			DataPoints: dps,
		},
	}
	got, err := tt.GetMetric("otelcol_loadbalancer_num_backends")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualLoadbalancerNumResolutions(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_loadbalancer_num_resolutions",
		Description: "Number of times the resolver has triggered new resolutions.",
		Unit:        "{resolutions}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_loadbalancer_num_resolutions")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}
