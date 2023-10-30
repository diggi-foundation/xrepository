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

// KomplexerBasistyp komplexer basistyp
//
// swagger:model KomplexerBasistyp
type KomplexerBasistyp struct {

	// art der ableitung
	// Required: true
	ArtDerAbleitung *string `json:"artDerAbleitung" xml:"artDerAbleitung"`

	// typ kennung
	TypKennung string `json:"typKennung,omitempty" xml:"typ.kennung,omitempty"`

	// typ name
	// Required: true
	TypName *string `json:"typName" xml:"typ.name"`
}

// Validate validates this komplexer basistyp
func (m *KomplexerBasistyp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtDerAbleitung(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KomplexerBasistyp) validateArtDerAbleitung(formats strfmt.Registry) error {

	if err := validate.Required("artDerAbleitung", "body", m.ArtDerAbleitung); err != nil {
		return err
	}

	return nil
}

func (m *KomplexerBasistyp) validateTypName(formats strfmt.Registry) error {

	if err := validate.Required("typName", "body", m.TypName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this komplexer basistyp based on context it is used
func (m *KomplexerBasistyp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KomplexerBasistyp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KomplexerBasistyp) UnmarshalBinary(b []byte) error {
	var res KomplexerBasistyp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
