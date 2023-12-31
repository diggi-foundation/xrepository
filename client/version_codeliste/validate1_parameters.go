// Code generated by go-swagger; DO NOT EDIT.

package version_codeliste

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

	"github.com/diggi-foundation/xrepository/models"
)

// NewValidate1Params creates a new Validate1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidate1Params() *Validate1Params {
	return &Validate1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidate1ParamsWithTimeout creates a new Validate1Params object
// with the ability to set a timeout on a request.
func NewValidate1ParamsWithTimeout(timeout time.Duration) *Validate1Params {
	return &Validate1Params{
		timeout: timeout,
	}
}

// NewValidate1ParamsWithContext creates a new Validate1Params object
// with the ability to set a context for a request.
func NewValidate1ParamsWithContext(ctx context.Context) *Validate1Params {
	return &Validate1Params{
		Context: ctx,
	}
}

// NewValidate1ParamsWithHTTPClient creates a new Validate1Params object
// with the ability to set a custom HTTPClient for a request.
func NewValidate1ParamsWithHTTPClient(client *http.Client) *Validate1Params {
	return &Validate1Params{
		HTTPClient: client,
	}
}

/*
Validate1Params contains all the parameters to send to the API endpoint

	for the validate 1 operation.

	Typically these are written to a http.Request.
*/
type Validate1Params struct {

	// Body.
	Body *models.MultipartFormDataInput

	// Kennung.
	Kennung string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Validate1Params) WithDefaults() *Validate1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Validate1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate 1 params
func (o *Validate1Params) WithTimeout(timeout time.Duration) *Validate1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate 1 params
func (o *Validate1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate 1 params
func (o *Validate1Params) WithContext(ctx context.Context) *Validate1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate 1 params
func (o *Validate1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate 1 params
func (o *Validate1Params) WithHTTPClient(client *http.Client) *Validate1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate 1 params
func (o *Validate1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate 1 params
func (o *Validate1Params) WithBody(body *models.MultipartFormDataInput) *Validate1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate 1 params
func (o *Validate1Params) SetBody(body *models.MultipartFormDataInput) {
	o.Body = body
}

// WithKennung adds the kennung to the validate 1 params
func (o *Validate1Params) WithKennung(kennung string) *Validate1Params {
	o.SetKennung(kennung)
	return o
}

// SetKennung adds the kennung to the validate 1 params
func (o *Validate1Params) SetKennung(kennung string) {
	o.Kennung = kennung
}

// WriteToRequest writes these params to a swagger request
func (o *Validate1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
