// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CallTraceState Full state object of a call trace
//
// swagger:model call_trace_state
type CallTraceState struct {

	// True if the trace is available for download
	CallTraceAvailable bool `json:"call_trace_available,omitempty"`

	// True if trace has been requested to end
	CallTraceEnding bool `json:"call_trace_ending,omitempty"`
}

// Validate validates this call trace state
func (m *CallTraceState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this call trace state based on context it is used
func (m *CallTraceState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CallTraceState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallTraceState) UnmarshalBinary(b []byte) error {
	var res CallTraceState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
