// Code generated by go-swagger; DO NOT EDIT.

package codeliste

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

// NewGetDebugInfoParams creates a new GetDebugInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDebugInfoParams() *GetDebugInfoParams {
	return &GetDebugInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDebugInfoParamsWithTimeout creates a new GetDebugInfoParams object
// with the ability to set a timeout on a request.
func NewGetDebugInfoParamsWithTimeout(timeout time.Duration) *GetDebugInfoParams {
	return &GetDebugInfoParams{
		timeout: timeout,
	}
}

// NewGetDebugInfoParamsWithContext creates a new GetDebugInfoParams object
// with the ability to set a context for a request.
func NewGetDebugInfoParamsWithContext(ctx context.Context) *GetDebugInfoParams {
	return &GetDebugInfoParams{
		Context: ctx,
	}
}

// NewGetDebugInfoParamsWithHTTPClient creates a new GetDebugInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDebugInfoParamsWithHTTPClient(client *http.Client) *GetDebugInfoParams {
	return &GetDebugInfoParams{
		HTTPClient: client,
	}
}

/*
GetDebugInfoParams contains all the parameters to send to the API endpoint

	for the get debug info operation.

	Typically these are written to a http.Request.
*/
type GetDebugInfoParams struct {

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get debug info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDebugInfoParams) WithDefaults() *GetDebugInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get debug info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDebugInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get debug info params
func (o *GetDebugInfoParams) WithTimeout(timeout time.Duration) *GetDebugInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get debug info params
func (o *GetDebugInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get debug info params
func (o *GetDebugInfoParams) WithContext(ctx context.Context) *GetDebugInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get debug info params
func (o *GetDebugInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get debug info params
func (o *GetDebugInfoParams) WithHTTPClient(client *http.Client) *GetDebugInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get debug info params
func (o *GetDebugInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUrn adds the urn to the get debug info params
func (o *GetDebugInfoParams) WithUrn(urn string) *GetDebugInfoParams {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the get debug info params
func (o *GetDebugInfoParams) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *GetDebugInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
