// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CodeStatusVeroeffentlichung code status veroeffentlichung
//
// swagger:model CodeStatusVeroeffentlichung
type CodeStatusVeroeffentlichung struct {

	// code
	// Required: true
	// Enum: [NICHT_VEROEFFENTLICHT VEROEFFENTLICHT]
	Code *string `json:"code" xml:"code"`
}

// Validate validates this code status veroeffentlichung
func (m *CodeStatusVeroeffentlichung) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var codeStatusVeroeffentlichungTypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NICHT_VEROEFFENTLICHT","VEROEFFENTLICHT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		codeStatusVeroeffentlichungTypeCodePropEnum = append(codeStatusVeroeffentlichungTypeCodePropEnum, v)
	}
}

const (

	// CodeStatusVeroeffentlichungCodeNICHTVEROEFFENTLICHT captures enum value "NICHT_VEROEFFENTLICHT"
	CodeStatusVeroeffentlichungCodeNICHTVEROEFFENTLICHT string = "NICHT_VEROEFFENTLICHT"

	// CodeStatusVeroeffentlichungCodeVEROEFFENTLICHT captures enum value "VEROEFFENTLICHT"
	CodeStatusVeroeffentlichungCodeVEROEFFENTLICHT string = "VEROEFFENTLICHT"
)

// prop value enum
func (m *CodeStatusVeroeffentlichung) validateCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, codeStatusVeroeffentlichungTypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CodeStatusVeroeffentlichung) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", *m.Code); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this code status veroeffentlichung based on context it is used
func (m *CodeStatusVeroeffentlichung) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CodeStatusVeroeffentlichung) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CodeStatusVeroeffentlichung) UnmarshalBinary(b []byte) error {
	var res CodeStatusVeroeffentlichung
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
