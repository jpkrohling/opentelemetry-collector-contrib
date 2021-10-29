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

package oidcauthextension

import "go.opentelemetry.io/collector/client"

var _ client.AuthData = (*authData)(nil)

type authData struct {
	raw        string
	subject    string
	membership []string
}

func (a *authData) Equal(other client.AuthData) bool {
	if other == nil {
		return false
	}

	otherAuthData, ok := other.(*authData)
	if !ok {
		return false
	}

	if a.subject != otherAuthData.subject {
		return false
	}

	if a.raw != otherAuthData.raw {
		return false
	}

	otherMembership := otherAuthData.membership
	if len(a.membership) != len(otherMembership) {
		return false
	}

	for i := range a.membership {
		if a.membership[i] != otherMembership[i] {
			return false
		}
	}

	return true
}

func (a *authData) GetAttribute(name string) interface{} {
	switch name {
	case "subject":
		return a.subject
	case "membership":
		return a.membership
	case "raw":
		return a.raw
	default:
		return nil
	}
}

func (*authData) GetAttributeNames() []string {
	return []string{"subject", "membership", "raw"}
}
