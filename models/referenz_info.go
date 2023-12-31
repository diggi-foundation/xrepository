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

// ReferenzInfo referenz info
//
// swagger:model ReferenzInfo
type ReferenzInfo struct {

	// binary
	Binary bool `json:"binary,omitempty" xml:"binary,attr,omitempty"`

	// kategorie
	// Required: true
	Kategorie *string `json:"kategorie" xml:"kategorie"`

	// kennung
	Kennung string `json:"kennung,omitempty" xml:"kennung,attr,omitempty"`

	// owner
	// Required: true
	Owner *string `json:"owner" xml:"owner"`

	// status
	Status string `json:"status,omitempty" xml:"status,attr,omitempty"`

	// type
	Type string `json:"type,omitempty" xml:"type,attr,omitempty"`
}

// Validate validates this referenz info
func (m *ReferenzInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKategorie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReferenzInfo) validateKategorie(formats strfmt.Registry) error {

	if err := validate.Required("kategorie", "body", m.Kategorie); err != nil {
		return err
	}

	return nil
}

func (m *ReferenzInfo) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this referenz info based on context it is used
func (m *ReferenzInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReferenzInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReferenzInfo) UnmarshalBinary(b []byte) error {
	var res ReferenzInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
