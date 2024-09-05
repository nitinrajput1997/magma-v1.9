// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// PageToken Base64-encoded page token for subsequent paginated API requests
//
// swagger:model page_token
type PageToken string

// Validate validates this page token
func (m PageToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this page token based on context it is used
func (m PageToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
