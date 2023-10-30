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

// NewGetBeziehungen1Params creates a new GetBeziehungen1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBeziehungen1Params() *GetBeziehungen1Params {
	return &GetBeziehungen1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBeziehungen1ParamsWithTimeout creates a new GetBeziehungen1Params object
// with the ability to set a timeout on a request.
func NewGetBeziehungen1ParamsWithTimeout(timeout time.Duration) *GetBeziehungen1Params {
	return &GetBeziehungen1Params{
		timeout: timeout,
	}
}

// NewGetBeziehungen1ParamsWithContext creates a new GetBeziehungen1Params object
// with the ability to set a context for a request.
func NewGetBeziehungen1ParamsWithContext(ctx context.Context) *GetBeziehungen1Params {
	return &GetBeziehungen1Params{
		Context: ctx,
	}
}

// NewGetBeziehungen1ParamsWithHTTPClient creates a new GetBeziehungen1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetBeziehungen1ParamsWithHTTPClient(client *http.Client) *GetBeziehungen1Params {
	return &GetBeziehungen1Params{
		HTTPClient: client,
	}
}

/*
GetBeziehungen1Params contains all the parameters to send to the API endpoint

	for the get beziehungen 1 operation.

	Typically these are written to a http.Request.
*/
type GetBeziehungen1Params struct {

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get beziehungen 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBeziehungen1Params) WithDefaults() *GetBeziehungen1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get beziehungen 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBeziehungen1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get beziehungen 1 params
func (o *GetBeziehungen1Params) WithTimeout(timeout time.Duration) *GetBeziehungen1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get beziehungen 1 params
func (o *GetBeziehungen1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get beziehungen 1 params
func (o *GetBeziehungen1Params) WithContext(ctx context.Context) *GetBeziehungen1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get beziehungen 1 params
func (o *GetBeziehungen1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get beziehungen 1 params
func (o *GetBeziehungen1Params) WithHTTPClient(client *http.Client) *GetBeziehungen1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get beziehungen 1 params
func (o *GetBeziehungen1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUrn adds the urn to the get beziehungen 1 params
func (o *GetBeziehungen1Params) WithUrn(urn string) *GetBeziehungen1Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the get beziehungen 1 params
func (o *GetBeziehungen1Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *GetBeziehungen1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
