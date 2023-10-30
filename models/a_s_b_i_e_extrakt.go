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

// ASBIEExtrakt a s b i e extrakt
//
// swagger:model ASBIEExtrakt
type ASBIEExtrakt struct {

	// abweichung
	Abweichung []string `json:"abweichung" xml:"abweichung"`

	// beschreibung
	Beschreibung string `json:"beschreibung,omitempty" xml:"beschreibung,omitempty"`

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
	Typ *KomplexerTypExtrakt `json:"typ" xml:"typ"`
}

// Validate validates this a s b i e extrakt
func (m *ASBIEExtrakt) Validate(formats strfmt.Registry) error {
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

func (m *ASBIEExtrakt) validateKernkomponenteneigenschaftKennung(formats strfmt.Registry) error {

	if err := validate.Required("kernkomponenteneigenschaftKennung", "body", m.KernkomponenteneigenschaftKennung); err != nil {
		return err
	}

	return nil
}

func (m *ASBIEExtrakt) validateMultiplizitaetLower(formats strfmt.Registry) error {

	if err := validate.Required("multiplizitaetLower", "body", m.MultiplizitaetLower); err != nil {
		return err
	}

	return nil
}

func (m *ASBIEExtrakt) validateNameTechnisch(formats strfmt.Registry) error {

	if err := validate.Required("nameTechnisch", "body", m.NameTechnisch); err != nil {
		return err
	}

	return nil
}

func (m *ASBIEExtrakt) validateTyp(formats strfmt.Registry) error {

	if err := validate.Required("typ", "body", m.Typ); err != nil {
		return err
	}

	if m.Typ != nil {
		if err := m.Typ.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typ")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("typ")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a s b i e extrakt based on the context it is used
func (m *ASBIEExtrakt) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTyp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ASBIEExtrakt) contextValidateTyp(ctx context.Context, formats strfmt.Registry) error {

	if m.Typ != nil {

		if err := m.Typ.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typ")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("typ")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ASBIEExtrakt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ASBIEExtrakt) UnmarshalBinary(b []byte) error {
	var res ASBIEExtrakt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
