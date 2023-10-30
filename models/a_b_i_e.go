// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ABIE a b i e
//
// swagger:model ABIE
type ABIE struct {

	// abweichung
	Abweichung []string `json:"abweichung" xml:"abweichung"`

	// asbie
	Asbie []*ASBIE `json:"asbie" xml:"asbie"`

	// basistyp
	Basistyp *KomplexerBasistyp `json:"basistyp,omitempty" xml:"basistyp,omitempty"`

	// bbie
	Bbie []*BBIE `json:"bbie" xml:"bbie"`

	// beschreibung
	Beschreibung string `json:"beschreibung,omitempty" xml:"beschreibung,omitempty"`

	// ergaenzte eigenschaft
	ErgaenzteEigenschaft []*ErgaenzteEigenschaft `json:"ergaenzteEigenschaft" xml:"ergaenzteEigenschaft"`

	// genutzte version kernkomponente kennung
	// Required: true
	GenutzteVersionKernkomponenteKennung *string `json:"genutzteVersionKernkomponenteKennung" xml:"genutzteVersionKernkomponente.kennung"`

	// kennung
	// Required: true
	Kennung *string `json:"kennung" xml:"kennung"`

	// motivation der abweichung
	MotivationDerAbweichung string `json:"motivationDerAbweichung,omitempty" xml:"motivationDerAbweichung,omitempty"`

	// technischer name
	// Required: true
	TechnischerName *string `json:"technischerName" xml:"technischerName"`
}

// Validate validates this a b i e
func (m *ABIE) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsbie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBasistyp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBbie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErgaenzteEigenschaft(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenutzteVersionKernkomponenteKennung(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKennung(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTechnischerName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ABIE) validateAsbie(formats strfmt.Registry) error {
	if swag.IsZero(m.Asbie) { // not required
		return nil
	}

	for i := 0; i < len(m.Asbie); i++ {
		if swag.IsZero(m.Asbie[i]) { // not required
			continue
		}

		if m.Asbie[i] != nil {
			if err := m.Asbie[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("asbie" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("asbie" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ABIE) validateBasistyp(formats strfmt.Registry) error {
	if swag.IsZero(m.Basistyp) { // not required
		return nil
	}

	if m.Basistyp != nil {
		if err := m.Basistyp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basistyp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basistyp")
			}
			return err
		}
	}

	return nil
}

func (m *ABIE) validateBbie(formats strfmt.Registry) error {
	if swag.IsZero(m.Bbie) { // not required
		return nil
	}

	for i := 0; i < len(m.Bbie); i++ {
		if swag.IsZero(m.Bbie[i]) { // not required
			continue
		}

		if m.Bbie[i] != nil {
			if err := m.Bbie[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bbie" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bbie" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ABIE) validateErgaenzteEigenschaft(formats strfmt.Registry) error {
	if swag.IsZero(m.ErgaenzteEigenschaft) { // not required
		return nil
	}

	for i := 0; i < len(m.ErgaenzteEigenschaft); i++ {
		if swag.IsZero(m.ErgaenzteEigenschaft[i]) { // not required
			continue
		}

		if m.ErgaenzteEigenschaft[i] != nil {
			if err := m.ErgaenzteEigenschaft[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ergaenzteEigenschaft" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ergaenzteEigenschaft" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ABIE) validateGenutzteVersionKernkomponenteKennung(formats strfmt.Registry) error {

	if err := validate.Required("genutzteVersionKernkomponenteKennung", "body", m.GenutzteVersionKernkomponenteKennung); err != nil {
		return err
	}

	return nil
}

func (m *ABIE) validateKennung(formats strfmt.Registry) error {

	if err := validate.Required("kennung", "body", m.Kennung); err != nil {
		return err
	}

	return nil
}

func (m *ABIE) validateTechnischerName(formats strfmt.Registry) error {

	if err := validate.Required("technischerName", "body", m.TechnischerName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this a b i e based on the context it is used
func (m *ABIE) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAsbie(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBasistyp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBbie(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErgaenzteEigenschaft(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ABIE) contextValidateAsbie(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Asbie); i++ {

		if m.Asbie[i] != nil {

			if swag.IsZero(m.Asbie[i]) { // not required
				return nil
			}

			if err := m.Asbie[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("asbie" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("asbie" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ABIE) contextValidateBasistyp(ctx context.Context, formats strfmt.Registry) error {

	if m.Basistyp != nil {

		if swag.IsZero(m.Basistyp) { // not required
			return nil
		}

		if err := m.Basistyp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basistyp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basistyp")
			}
			return err
		}
	}

	return nil
}

func (m *ABIE) contextValidateBbie(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Bbie); i++ {

		if m.Bbie[i] != nil {

			if swag.IsZero(m.Bbie[i]) { // not required
				return nil
			}

			if err := m.Bbie[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bbie" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bbie" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ABIE) contextValidateErgaenzteEigenschaft(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ErgaenzteEigenschaft); i++ {

		if m.ErgaenzteEigenschaft[i] != nil {

			if swag.IsZero(m.ErgaenzteEigenschaft[i]) { // not required
				return nil
			}

			if err := m.ErgaenzteEigenschaft[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ergaenzteEigenschaft" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ergaenzteEigenschaft" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ABIE) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ABIE) UnmarshalBinary(b []byte) error {
	var res ABIE
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}