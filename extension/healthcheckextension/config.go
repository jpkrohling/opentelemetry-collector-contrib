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

package healthcheckextension

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confignet"
)

// Config has the configuration for the extension enabling the health check
// extension, used to report the health status of the service.
type Config struct {
	config.ExtensionSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct

	// Port is the port used to publish the health check status.
	// The default value is 13133.
	// Deprecated: use Endpoint instead.
	Port uint16 `mapstructure:"port"`

	// TCPAddr represents a tcp endpoint address that is to publish the health
	// check status.
	// The default endpoint is "0.0.0.0:13133".
	TCPAddr confignet.TCPAddr `mapstructure:",squash"`

	// metricsHealthCheck returns exporters functional health check settings
	MetricsHealthCheck metricsHealthCheckSettings `mapstructure:"metrics_health_check"`
}

var _ config.Extension = (*Config)(nil)

// Validate checks if the extension configuration is valid
func (cfg *Config) Validate() error {
	_, err := time.ParseDuration(cfg.MetricsHealthCheck.Interval)
	if err != nil {
		return err
	}
	if cfg.TCPAddr.Endpoint == "" {
		return errors.New("bad config: endpoint must be specified")
	}
	if cfg.MetricsHealthCheck.ExporterFailureThreshold <= 0 {
		return errors.New("bad config: exporter_failure_threshold expects a positive number")
	}
	return nil
}

type metricsHealthCheckSettings struct {
	// Enabled indicates whether to not enable metrics health check.
	Enabled bool `mapstructure:"enabled"`
	// Interval the time range to check healthy status based on the exporter failures
	Interval string `mapstructure:"interval"`
	// ExporterFailureThreshold is the threshold of failure numbers status during the Interval
	ExporterFailureThreshold int `mapstructure:"exporter_failure_threshold"`
}

// DefaultMetricsHealthCheckSettings returns the default settings for metricsHealthCheck.
func DefaultMetricsHealthCheckSettings() metricsHealthCheckSettings {
	return metricsHealthCheckSettings{
		Enabled:                  false,
		Interval:                 "5m",
		ExporterFailureThreshold: 5,
	}
}
