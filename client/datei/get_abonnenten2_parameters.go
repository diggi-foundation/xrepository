// Code generated by go-swagger; DO NOT EDIT.

package datei

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

// NewGetAbonnenten2Params creates a new GetAbonnenten2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAbonnenten2Params() *GetAbonnenten2Params {
	return &GetAbonnenten2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAbonnenten2ParamsWithTimeout creates a new GetAbonnenten2Params object
// with the ability to set a timeout on a request.
func NewGetAbonnenten2ParamsWithTimeout(timeout time.Duration) *GetAbonnenten2Params {
	return &GetAbonnenten2Params{
		timeout: timeout,
	}
}

// NewGetAbonnenten2ParamsWithContext creates a new GetAbonnenten2Params object
// with the ability to set a context for a request.
func NewGetAbonnenten2ParamsWithContext(ctx context.Context) *GetAbonnenten2Params {
	return &GetAbonnenten2Params{
		Context: ctx,
	}
}

// NewGetAbonnenten2ParamsWithHTTPClient creates a new GetAbonnenten2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAbonnenten2ParamsWithHTTPClient(client *http.Client) *GetAbonnenten2Params {
	return &GetAbonnenten2Params{
		HTTPClient: client,
	}
}

/*
GetAbonnenten2Params contains all the parameters to send to the API endpoint

	for the get abonnenten 2 operation.

	Typically these are written to a http.Request.
*/
type GetAbonnenten2Params struct {

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get abonnenten 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAbonnenten2Params) WithDefaults() *GetAbonnenten2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get abonnenten 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAbonnenten2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get abonnenten 2 params
func (o *GetAbonnenten2Params) WithTimeout(timeout time.Duration) *GetAbonnenten2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get abonnenten 2 params
func (o *GetAbonnenten2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get abonnenten 2 params
func (o *GetAbonnenten2Params) WithContext(ctx context.Context) *GetAbonnenten2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get abonnenten 2 params
func (o *GetAbonnenten2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get abonnenten 2 params
func (o *GetAbonnenten2Params) WithHTTPClient(client *http.Client) *GetAbonnenten2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get abonnenten 2 params
func (o *GetAbonnenten2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUrn adds the urn to the get abonnenten 2 params
func (o *GetAbonnenten2Params) WithUrn(urn string) *GetAbonnenten2Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the get abonnenten 2 params
func (o *GetAbonnenten2Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *GetAbonnenten2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
