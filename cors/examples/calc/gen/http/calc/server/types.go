// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen goa.design/plugins/cors/examples/calc/design

package server

import (
	calcsvc "goa.design/plugins/cors/examples/calc/gen/calc"
)

// NewAddAddPayload builds a calc service add endpoint payload.
func NewAddAddPayload(a int, b int) *calcsvc.AddPayload {
	return &calcsvc.AddPayload{
		A: a,
		B: b,
	}
}
