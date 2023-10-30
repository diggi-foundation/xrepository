// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NsPrefixInAttributeValues ns prefix in attribute values
//
// swagger:model NsPrefixInAttributeValues
type NsPrefixInAttributeValues struct {

	// prefix
	Prefix string `json:"prefix,omitempty" xml:"prefix,attr,omitempty"`

	// uri
	URI string `json:"uri,omitempty" xml:"uri,attr,omitempty"`
}

// Validate validates this ns prefix in attribute values
func (m *NsPrefixInAttributeValues) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ns prefix in attribute values based on context it is used
func (m *NsPrefixInAttributeValues) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NsPrefixInAttributeValues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NsPrefixInAttributeValues) UnmarshalBinary(b []byte) error {
	var res NsPrefixInAttributeValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
