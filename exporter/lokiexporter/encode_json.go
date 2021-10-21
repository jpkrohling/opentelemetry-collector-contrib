// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lokiexporter

import (
	"encoding/json"

	"go.opentelemetry.io/collector/model/pdata"
)

// JSON representation of the LogRecord as described by https://developers.google.com/protocol-buffers/docs/proto3#json

type lokiEntry struct {
	Name       string                 `json:"name,omitempty"`
	Body       string                 `json:"body,omitempty"`
	TraceID    string                 `json:"traceid,omitempty"`
	SpanID     string                 `json:"spanid,omitempty"`
	Severity   string                 `json:"severity,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Resources  map[string]interface{} `json:"resources,omitempty"`
}

func encodeJSON(lr pdata.LogRecord, res pdata.Resource) (string, error) {
	var logRecord lokiEntry
	var jsonRecord []byte

	logRecord = lokiEntry{
		Name:       lr.Name(),
		Body:       lr.Body().StringVal(),
		TraceID:    lr.TraceID().HexString(),
		SpanID:     lr.SpanID().HexString(),
		Severity:   lr.SeverityText(),
		Attributes: lr.Attributes().AsRaw(),
		Resources:  res.Attributes().AsRaw(),
	}

	jsonRecord, err := json.Marshal(logRecord)
	if err != nil {
		return "", err
	}
	return string(jsonRecord), nil
}
