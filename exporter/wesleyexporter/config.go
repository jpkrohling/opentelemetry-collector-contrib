// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package wesleyexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/wesleyexporter"

// Config defines configuration for Loki exporter.
type Config struct {
	Foo int `mapstructure:"foo"`
	Bar int `mapstructure:"bar"`
	Bla int `mapstructure:"bla"`
}

func (c *Config) Validate() error {
	return nil
}
