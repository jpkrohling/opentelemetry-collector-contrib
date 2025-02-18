// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/connector/connectortest"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"
)

func NewSettings(tt *componenttest.Telemetry) connector.Settings {
	set := connectortest.NewNopSettingsWithType(connectortest.NopType)
	set.ID = component.NewID(component.MustNewType("grafanacloud"))
	set.TelemetrySettings = tt.NewTelemetrySettings()
	return set
}

func AssertEqualGrafanacloudDatapointCount(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_grafanacloud_datapoint_count",
		Description: "Number of datapoints sent to Grafana Cloud",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_grafanacloud_datapoint_count")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualGrafanacloudFlushCount(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_grafanacloud_flush_count",
		Description: "Number of metrics flushes",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_grafanacloud_flush_count")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualGrafanacloudHostCount(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_grafanacloud_host_count",
		Description: "Number of unique hosts",
		Unit:        "1",
		Data: metricdata.Gauge[int64]{
			DataPoints: dps,
		},
	}
	got, err := tt.GetMetric("otelcol_grafanacloud_host_count")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}
