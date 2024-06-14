// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package wesleyexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/wesleyexporter"

type wesleyExporter struct {
}

func newExporter(*Config) (*wesleyExporter, error) {
	return &wesleyExporter{}, nil
}
