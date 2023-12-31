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

// MultipartFormDataInput multipart form data input
//
// swagger:model MultipartFormDataInput
type MultipartFormDataInput struct {

	// form data
	FormData map[string]InputPart `json:"formData,omitempty"`

	// form data map
	FormDataMap map[string][]InputPart `json:"formDataMap,omitempty"`

	// parts
	Parts []*InputPart `json:"parts"`

	// preamble
	Preamble string `json:"preamble,omitempty"`
}

// Validate validates this multipart form data input
func (m *MultipartFormDataInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormDataMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultipartFormDataInput) validateFormData(formats strfmt.Registry) error {
	if swag.IsZero(m.FormData) { // not required
		return nil
	}

	for k := range m.FormData {

		if err := validate.Required("formData"+"."+k, "body", m.FormData[k]); err != nil {
			return err
		}
		if val, ok := m.FormData[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("formData" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("formData" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultipartFormDataInput) validateFormDataMap(formats strfmt.Registry) error {
	if swag.IsZero(m.FormDataMap) { // not required
		return nil
	}

	for k := range m.FormDataMap {

		if err := validate.Required("formDataMap"+"."+k, "body", m.FormDataMap[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.FormDataMap[k]); i++ {

			if err := m.FormDataMap[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("formDataMap" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("formDataMap" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *MultipartFormDataInput) validateParts(formats strfmt.Registry) error {
	if swag.IsZero(m.Parts) { // not required
		return nil
	}

	for i := 0; i < len(m.Parts); i++ {
		if swag.IsZero(m.Parts[i]) { // not required
			continue
		}

		if m.Parts[i] != nil {
			if err := m.Parts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this multipart form data input based on the context it is used
func (m *MultipartFormDataInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFormData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFormDataMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultipartFormDataInput) contextValidateFormData(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.FormData {

		if val, ok := m.FormData[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *MultipartFormDataInput) contextValidateFormDataMap(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.FormDataMap {

		for i := 0; i < len(m.FormDataMap[k]); i++ {

			if swag.IsZero(m.FormDataMap[k][i]) { // not required
				return nil
			}

			if err := m.FormDataMap[k][i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("formDataMap" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("formDataMap" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *MultipartFormDataInput) contextValidateParts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parts); i++ {

		if m.Parts[i] != nil {

			if swag.IsZero(m.Parts[i]) { // not required
				return nil
			}

			if err := m.Parts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MultipartFormDataInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultipartFormDataInput) UnmarshalBinary(b []byte) error {
	var res MultipartFormDataInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
