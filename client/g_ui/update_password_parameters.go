// Code generated by go-swagger; DO NOT EDIT.

package g_ui

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/diggi.foundation/xrepository/models"
)

// NewUpdatePasswordParams creates a new UpdatePasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePasswordParams() *UpdatePasswordParams {
	return &UpdatePasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePasswordParamsWithTimeout creates a new UpdatePasswordParams object
// with the ability to set a timeout on a request.
func NewUpdatePasswordParamsWithTimeout(timeout time.Duration) *UpdatePasswordParams {
	return &UpdatePasswordParams{
		timeout: timeout,
	}
}

// NewUpdatePasswordParamsWithContext creates a new UpdatePasswordParams object
// with the ability to set a context for a request.
func NewUpdatePasswordParamsWithContext(ctx context.Context) *UpdatePasswordParams {
	return &UpdatePasswordParams{
		Context: ctx,
	}
}

// NewUpdatePasswordParamsWithHTTPClient creates a new UpdatePasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePasswordParamsWithHTTPClient(client *http.Client) *UpdatePasswordParams {
	return &UpdatePasswordParams{
		HTTPClient: client,
	}
}

/*
UpdatePasswordParams contains all the parameters to send to the API endpoint

	for the update password operation.

	Typically these are written to a http.Request.
*/
type UpdatePasswordParams struct {

	// Body.
	Body *models.ChangePasswordData

	// Kennung.
	Kennung string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePasswordParams) WithDefaults() *UpdatePasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update password params
func (o *UpdatePasswordParams) WithTimeout(timeout time.Duration) *UpdatePasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update password params
func (o *UpdatePasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update password params
func (o *UpdatePasswordParams) WithContext(ctx context.Context) *UpdatePasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update password params
func (o *UpdatePasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update password params
func (o *UpdatePasswordParams) WithHTTPClient(client *http.Client) *UpdatePasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update password params
func (o *UpdatePasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update password params
func (o *UpdatePasswordParams) WithBody(body *models.ChangePasswordData) *UpdatePasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update password params
func (o *UpdatePasswordParams) SetBody(body *models.ChangePasswordData) {
	o.Body = body
}

// WithKennung adds the kennung to the update password params
func (o *UpdatePasswordParams) WithKennung(kennung string) *UpdatePasswordParams {
	o.SetKennung(kennung)
	return o
}

// SetKennung adds the kennung to the update password params
func (o *UpdatePasswordParams) SetKennung(kennung string) {
	o.Kennung = kennung
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param kennung
	if err := r.SetPathParam("kennung", o.Kennung); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}