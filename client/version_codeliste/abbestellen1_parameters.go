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
)

// NewAbbestellen1Params creates a new Abbestellen1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAbbestellen1Params() *Abbestellen1Params {
	return &Abbestellen1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAbbestellen1ParamsWithTimeout creates a new Abbestellen1Params object
// with the ability to set a timeout on a request.
func NewAbbestellen1ParamsWithTimeout(timeout time.Duration) *Abbestellen1Params {
	return &Abbestellen1Params{
		timeout: timeout,
	}
}

// NewAbbestellen1ParamsWithContext creates a new Abbestellen1Params object
// with the ability to set a context for a request.
func NewAbbestellen1ParamsWithContext(ctx context.Context) *Abbestellen1Params {
	return &Abbestellen1Params{
		Context: ctx,
	}
}

// NewAbbestellen1ParamsWithHTTPClient creates a new Abbestellen1Params object
// with the ability to set a custom HTTPClient for a request.
func NewAbbestellen1ParamsWithHTTPClient(client *http.Client) *Abbestellen1Params {
	return &Abbestellen1Params{
		HTTPClient: client,
	}
}

/*
Abbestellen1Params contains all the parameters to send to the API endpoint

	for the abbestellen 1 operation.

	Typically these are written to a http.Request.
*/
type Abbestellen1Params struct {

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the abbestellen 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Abbestellen1Params) WithDefaults() *Abbestellen1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the abbestellen 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Abbestellen1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the abbestellen 1 params
func (o *Abbestellen1Params) WithTimeout(timeout time.Duration) *Abbestellen1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the abbestellen 1 params
func (o *Abbestellen1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the abbestellen 1 params
func (o *Abbestellen1Params) WithContext(ctx context.Context) *Abbestellen1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the abbestellen 1 params
func (o *Abbestellen1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the abbestellen 1 params
func (o *Abbestellen1Params) WithHTTPClient(client *http.Client) *Abbestellen1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the abbestellen 1 params
func (o *Abbestellen1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUrn adds the urn to the abbestellen 1 params
func (o *Abbestellen1Params) WithUrn(urn string) *Abbestellen1Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the abbestellen 1 params
func (o *Abbestellen1Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *Abbestellen1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param urn
	if err := r.SetPathParam("urn", o.Urn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
