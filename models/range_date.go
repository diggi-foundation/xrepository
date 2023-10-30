// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RangeDate range date
//
// swagger:model RangeDate
type RangeDate struct {

	// max
	// Format: date-time
	Max strfmt.DateTime `json:"max,omitempty"`

	// min
	// Format: date-time
	Min strfmt.DateTime `json:"min,omitempty"`
}

// Validate validates this range date
func (m *RangeDate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RangeDate) validateMax(formats strfmt.Registry) error {
	if swag.IsZero(m.Max) { // not required
		return nil
	}

	if err := validate.FormatOf("max", "body", "date-time", m.Max.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RangeDate) validateMin(formats strfmt.Registry) error {
	if swag.IsZero(m.Min) { // not required
		return nil
	}

	if err := validate.FormatOf("min", "body", "date-time", m.Min.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this range date based on context it is used
func (m *RangeDate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RangeDate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RangeDate) UnmarshalBinary(b []byte) error {
	var res RangeDate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
