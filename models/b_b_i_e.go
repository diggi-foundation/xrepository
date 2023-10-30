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

// BBIE b b i e
//
// swagger:model BBIE
type BBIE struct {

	// abweichung
	Abweichung []string `json:"abweichung" xml:"abweichung"`

	// beschreibung
	Beschreibung string `json:"beschreibung,omitempty" xml:"beschreibung,omitempty"`

	// genutzte codeliste kennung
	GenutzteCodelisteKennung string `json:"genutzteCodelisteKennung,omitempty" xml:"genutzteCodeliste.kennung,omitempty"`

	// genutzte version codeliste kennung
	GenutzteVersionCodelisteKennung string `json:"genutzteVersionCodelisteKennung,omitempty" xml:"genutzteVersionCodeliste.kennung,omitempty"`

	// kernkomponenteneigenschaft kennung
	// Required: true
	KernkomponenteneigenschaftKennung []string `json:"kernkomponenteneigenschaftKennung" xml:"kernkomponenteneigenschaft.kennung"`

	// motivation der abweichung
	MotivationDerAbweichung string `json:"motivationDerAbweichung,omitempty" xml:"motivationDerAbweichung,omitempty"`

	// multiplizitaet lower
	// Required: true
	MultiplizitaetLower *int64 `json:"multiplizitaetLower" xml:"multiplizitaetLower"`

	// multiplizitaet upper
	MultiplizitaetUpper int64 `json:"multiplizitaetUpper,omitempty" xml:"multiplizitaetUpper,omitempty"`

	// name technisch
	// Required: true
	NameTechnisch *string `json:"nameTechnisch" xml:"nameTechnisch"`

	// typ
	// Required: true
	Typ *string `json:"typ" xml:"typ"`
}

// Validate validates this b b i e
func (m *BBIE) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKernkomponenteneigenschaftKennung(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMultiplizitaetLower(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameTechnisch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTyp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BBIE) validateKernkomponenteneigenschaftKennung(formats strfmt.Registry) error {

	if err := validate.Required("kernkomponenteneigenschaftKennung", "body", m.KernkomponenteneigenschaftKennung); err != nil {
		return err
	}

	return nil
}

func (m *BBIE) validateMultiplizitaetLower(formats strfmt.Registry) error {

	if err := validate.Required("multiplizitaetLower", "body", m.MultiplizitaetLower); err != nil {
		return err
	}

	return nil
}

func (m *BBIE) validateNameTechnisch(formats strfmt.Registry) error {

	if err := validate.Required("nameTechnisch", "body", m.NameTechnisch); err != nil {
		return err
	}

	return nil
}

func (m *BBIE) validateTyp(formats strfmt.Registry) error {

	if err := validate.Required("typ", "body", m.Typ); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this b b i e based on context it is used
func (m *BBIE) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BBIE) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BBIE) UnmarshalBinary(b []byte) error {
	var res BBIE
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
