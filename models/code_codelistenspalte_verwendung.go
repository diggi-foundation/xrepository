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

// CodeCodelistenspalteVerwendung code codelistenspalte verwendung
//
// swagger:model CodeCodelistenspalteVerwendung
type CodeCodelistenspalteVerwendung struct {

	// code
	// Required: true
	// Enum: [OPTIONAL REQUIRED]
	Code *string `json:"code" xml:"code"`
}

// Validate validates this code codelistenspalte verwendung
func (m *CodeCodelistenspalteVerwendung) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var codeCodelistenspalteVerwendungTypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPTIONAL","REQUIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		codeCodelistenspalteVerwendungTypeCodePropEnum = append(codeCodelistenspalteVerwendungTypeCodePropEnum, v)
	}
}

const (

	// CodeCodelistenspalteVerwendungCodeOPTIONAL captures enum value "OPTIONAL"
	CodeCodelistenspalteVerwendungCodeOPTIONAL string = "OPTIONAL"

	// CodeCodelistenspalteVerwendungCodeREQUIRED captures enum value "REQUIRED"
	CodeCodelistenspalteVerwendungCodeREQUIRED string = "REQUIRED"
)

// prop value enum
func (m *CodeCodelistenspalteVerwendung) validateCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, codeCodelistenspalteVerwendungTypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CodeCodelistenspalteVerwendung) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", *m.Code); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this code codelistenspalte verwendung based on context it is used
func (m *CodeCodelistenspalteVerwendung) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CodeCodelistenspalteVerwendung) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CodeCodelistenspalteVerwendung) UnmarshalBinary(b []byte) error {
	var res CodeCodelistenspalteVerwendung
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}