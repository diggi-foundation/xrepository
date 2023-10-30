// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Text text
//
// swagger:model Text
type Text struct {

	// content
	Content []interface{} `json:"content"`

	// fpi
	Fpi string `json:"fpi,omitempty" xml:"fpi,attr,omitempty"`

	// icon
	Icon string `json:"icon,omitempty" xml:"icon,attr,omitempty"`

	// lang
	Lang string `json:"lang,omitempty" xml:"lang,attr,omitempty"`

	// other attributes
	OtherAttributes map[string]string `json:"otherAttributes,omitempty"`

	// see
	See string `json:"see,omitempty" xml:"see,attr,omitempty"`

	// space
	Space string `json:"space,omitempty" xml:"space,attr,omitempty"`
}

// Validate validates this text
func (m *Text) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this text based on context it is used
func (m *Text) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Text) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Text) UnmarshalBinary(b []byte) error {
	var res Text
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
