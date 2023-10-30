// Code generated by go-swagger; DO NOT EDIT.

package version_standard

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

// NewGetJSON4Params creates a new GetJSON4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJSON4Params() *GetJSON4Params {
	return &GetJSON4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJSON4ParamsWithTimeout creates a new GetJSON4Params object
// with the ability to set a timeout on a request.
func NewGetJSON4ParamsWithTimeout(timeout time.Duration) *GetJSON4Params {
	return &GetJSON4Params{
		timeout: timeout,
	}
}

// NewGetJSON4ParamsWithContext creates a new GetJSON4Params object
// with the ability to set a context for a request.
func NewGetJSON4ParamsWithContext(ctx context.Context) *GetJSON4Params {
	return &GetJSON4Params{
		Context: ctx,
	}
}

// NewGetJSON4ParamsWithHTTPClient creates a new GetJSON4Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetJSON4ParamsWithHTTPClient(client *http.Client) *GetJSON4Params {
	return &GetJSON4Params{
		HTTPClient: client,
	}
}

/*
GetJSON4Params contains all the parameters to send to the API endpoint

	for the get Json 4 operation.

	Typically these are written to a http.Request.
*/
type GetJSON4Params struct {

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Json 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJSON4Params) WithDefaults() *GetJSON4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Json 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJSON4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Json 4 params
func (o *GetJSON4Params) WithTimeout(timeout time.Duration) *GetJSON4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Json 4 params
func (o *GetJSON4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Json 4 params
func (o *GetJSON4Params) WithContext(ctx context.Context) *GetJSON4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Json 4 params
func (o *GetJSON4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Json 4 params
func (o *GetJSON4Params) WithHTTPClient(client *http.Client) *GetJSON4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Json 4 params
func (o *GetJSON4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUrn adds the urn to the get Json 4 params
func (o *GetJSON4Params) WithUrn(urn string) *GetJSON4Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the get Json 4 params
func (o *GetJSON4Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *GetJSON4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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