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

import "go.opentelemetry.io/collector/config/configauth"

var _ configauth.AuthContext = (*authContext)(nil)

type authContext struct {
	// The raw token.
	Raw string `json:"raw"`
	// The access token.
	Subject string `json:"subject"`
	// The refresh token.
	Membership []string `json:"membership"`
}

func (a *authContext) Equal(other interface{}) bool {
	if otherAuthContext, ok := other.(*authContext); ok {
		if a.Raw == otherAuthContext.Raw &&
			a.Subject == otherAuthContext.Subject {
			if len(a.Membership) == len(otherAuthContext.Membership) {
				for i := range a.Membership {
					if a.Membership[i] != otherAuthContext.Membership[i] {
						return false
					}
				}
				return true
			}
		}
	}
	return false
}
func (a *authContext) GetAttribute(attrName string) interface{} {
	switch attrName {
	case "raw":
		return a.Raw
	case "subject":
		return a.Subject
	case "membership":
		return a.Membership
	}
	return nil
}
func (a *authContext) GetAttributeNames() []string {
	return []string{"raw", "subject", "membership"}
}
