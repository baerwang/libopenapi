// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"github.com/pb33f/libopenapi/datamodel/high"
	low "github.com/pb33f/libopenapi/datamodel/low/2.0"
)

type SecurityScheme struct {
	Type             string
	Description      string
	Name             string
	In               string
	Flow             string
	AuthorizationUrl string
	TokenUrl         string
	Scopes           *Scopes
	Extensions       map[string]any
	low              *low.SecurityScheme
}

func NewSecurityScheme(securityScheme *low.SecurityScheme) *SecurityScheme {
	s := new(SecurityScheme)
	s.low = securityScheme
	s.Extensions = high.ExtractExtensions(securityScheme.Extensions)
	if !securityScheme.Type.IsEmpty() {
		s.Type = securityScheme.Type.Value
	}
	if !securityScheme.Description.IsEmpty() {
		s.Description = securityScheme.Description.Value
	}
	if !securityScheme.Name.IsEmpty() {
		s.Name = securityScheme.Name.Value
	}
	if !securityScheme.In.IsEmpty() {
		s.In = securityScheme.In.Value
	}
	if !securityScheme.Flow.IsEmpty() {
		s.Flow = securityScheme.Flow.Value
	}
	if !securityScheme.AuthorizationUrl.IsEmpty() {
		s.AuthorizationUrl = securityScheme.AuthorizationUrl.Value
	}
	if !securityScheme.TokenUrl.IsEmpty() {
		s.TokenUrl = securityScheme.TokenUrl.Value
	}
	if !securityScheme.Scopes.IsEmpty() {
		s.Scopes = NewScopes(securityScheme.Scopes.Value)
	}
	return s
}