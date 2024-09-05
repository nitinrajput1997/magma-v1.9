// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SlackConfirmField slack confirm field
// swagger:model slack_confirm_field
type SlackConfirmField struct {

	// dismiss text
	// Required: true
	DismissText *string `json:"dismiss_text"`

	// ok text
	// Required: true
	OkText *string `json:"ok_text"`

	// text
	// Required: true
	Text *string `json:"text"`

	// title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this slack confirm field
func (m *SlackConfirmField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDismissText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOkText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SlackConfirmField) validateDismissText(formats strfmt.Registry) error {

	if err := validate.Required("dismiss_text", "body", m.DismissText); err != nil {
		return err
	}

	return nil
}

func (m *SlackConfirmField) validateOkText(formats strfmt.Registry) error {

	if err := validate.Required("ok_text", "body", m.OkText); err != nil {
		return err
	}

	return nil
}

func (m *SlackConfirmField) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

func (m *SlackConfirmField) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SlackConfirmField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SlackConfirmField) UnmarshalBinary(b []byte) error {
	var res SlackConfirmField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
