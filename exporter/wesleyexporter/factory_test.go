// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package wesleyexporter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExporter_new(t *testing.T) {
	t.Run("with valid config", func(t *testing.T) {
		cfg := &Config{}
		exp, err := newExporter(cfg)
		require.NoError(t, err)
		require.NotNil(t, exp)
	})
}
