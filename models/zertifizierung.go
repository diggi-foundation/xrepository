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

// Zertifizierung zertifizierung
//
// swagger:model Zertifizierung
type Zertifizierung struct {

	// letzte aenderung pflegekonzept
	// Required: true
	// Format: date-time
	LetzteAenderungPflegekonzept *strfmt.DateTime `json:"letzteAenderungPflegekonzept" xml:"standard.dokumentation.pflegekonzept.zeitpunktLetzteAenderungVorBeantragen"`

	// letzte anpassung fachmodell proprietaer
	// Required: true
	// Format: date-time
	LetzteAnpassungFachmodellProprietaer *strfmt.DateTime `json:"letzteAnpassungFachmodellProprietaer" xml:"versionStandard.originalFachmodellProprietaer.zeitpunktLetzteAenderungVorBeantragen"`

	// letzte anpassung spezifikation
	// Required: true
	// Format: date-time
	LetzteAnpassungSpezifikation *strfmt.DateTime `json:"letzteAnpassungSpezifikation" xml:"versionStandard.dokumentation.spezifikation.zeitpunktLetzteAenderungVorBeantragen"`

	// letzte anpassung x m i
	// Required: true
	// Format: date-time
	LetzteAnpassungXMI *strfmt.DateTime `json:"letzteAnpassungXMI" xml:"versionStandard.originalFachmodellXMI.zeitpunktLetzteAenderungVorBeantragen"`

	// letzte anpassung x s d
	// Required: true
	// Format: date-time
	LetzteAnpassungXSD *strfmt.DateTime `json:"letzteAnpassungXSD" xml:"versionStandard.xmlschema.zeitpunktLetzteAenderungVorBeantragen"`

	// zeitpunkt beantragung
	// Required: true
	// Format: date-time
	ZeitpunktBeantragung *strfmt.DateTime `json:"zeitpunktBeantragung" xml:"zeitpunktBeantragung"`

	// zeitpunkt zertifizierung
	// Format: date-time
	ZeitpunktZertifizierung strfmt.DateTime `json:"zeitpunktZertifizierung,omitempty" xml:"zeitpunktZertifizierung,omitempty"`
}

// Validate validates this zertifizierung
func (m *Zertifizierung) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLetzteAenderungPflegekonzept(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLetzteAnpassungFachmodellProprietaer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLetzteAnpassungSpezifikation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLetzteAnpassungXMI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLetzteAnpassungXSD(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZeitpunktBeantragung(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZeitpunktZertifizierung(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Zertifizierung) validateLetzteAenderungPflegekonzept(formats strfmt.Registry) error {

	if err := validate.Required("letzteAenderungPflegekonzept", "body", m.LetzteAenderungPflegekonzept); err != nil {
		return err
	}

	if err := validate.FormatOf("letzteAenderungPflegekonzept", "body", "date-time", m.LetzteAenderungPflegekonzept.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Zertifizierung) validateLetzteAnpassungFachmodellProprietaer(formats strfmt.Registry) error {

	if err := validate.Required("letzteAnpassungFachmodellProprietaer", "body", m.LetzteAnpassungFachmodellProprietaer); err != nil {
		return err
	}

	if err := validate.FormatOf("letzteAnpassungFachmodellProprietaer", "body", "date-time", m.LetzteAnpassungFachmodellProprietaer.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Zertifizierung) validateLetzteAnpassungSpezifikation(formats strfmt.Registry) error {

	if err := validate.Required("letzteAnpassungSpezifikation", "body", m.LetzteAnpassungSpezifikation); err != nil {
		return err
	}

	if err := validate.FormatOf("letzteAnpassungSpezifikation", "body", "date-time", m.LetzteAnpassungSpezifikation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Zertifizierung) validateLetzteAnpassungXMI(formats strfmt.Registry) error {

	if err := validate.Required("letzteAnpassungXMI", "body", m.LetzteAnpassungXMI); err != nil {
		return err
	}

	if err := validate.FormatOf("letzteAnpassungXMI", "body", "date-time", m.LetzteAnpassungXMI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Zertifizierung) validateLetzteAnpassungXSD(formats strfmt.Registry) error {

	if err := validate.Required("letzteAnpassungXSD", "body", m.LetzteAnpassungXSD); err != nil {
		return err
	}

	if err := validate.FormatOf("letzteAnpassungXSD", "body", "date-time", m.LetzteAnpassungXSD.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Zertifizierung) validateZeitpunktBeantragung(formats strfmt.Registry) error {

	if err := validate.Required("zeitpunktBeantragung", "body", m.ZeitpunktBeantragung); err != nil {
		return err
	}

	if err := validate.FormatOf("zeitpunktBeantragung", "body", "date-time", m.ZeitpunktBeantragung.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Zertifizierung) validateZeitpunktZertifizierung(formats strfmt.Registry) error {
	if swag.IsZero(m.ZeitpunktZertifizierung) { // not required
		return nil
	}

	if err := validate.FormatOf("zeitpunktZertifizierung", "body", "date-time", m.ZeitpunktZertifizierung.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this zertifizierung based on context it is used
func (m *Zertifizierung) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Zertifizierung) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Zertifizierung) UnmarshalBinary(b []byte) error {
	var res Zertifizierung
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
