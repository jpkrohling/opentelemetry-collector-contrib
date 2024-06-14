// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package wesleyexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/wesleyexporter"

// Config defines configuration for Loki exporter.
type Config struct {
}

func (c *Config) Validate() error {
	return nil
}
