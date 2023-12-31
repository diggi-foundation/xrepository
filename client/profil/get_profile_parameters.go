// Code generated by go-swagger; DO NOT EDIT.

package profil

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

// NewGetProfileParams creates a new GetProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProfileParams() *GetProfileParams {
	return &GetProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProfileParamsWithTimeout creates a new GetProfileParams object
// with the ability to set a timeout on a request.
func NewGetProfileParamsWithTimeout(timeout time.Duration) *GetProfileParams {
	return &GetProfileParams{
		timeout: timeout,
	}
}

// NewGetProfileParamsWithContext creates a new GetProfileParams object
// with the ability to set a context for a request.
func NewGetProfileParamsWithContext(ctx context.Context) *GetProfileParams {
	return &GetProfileParams{
		Context: ctx,
	}
}

// NewGetProfileParamsWithHTTPClient creates a new GetProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProfileParamsWithHTTPClient(client *http.Client) *GetProfileParams {
	return &GetProfileParams{
		HTTPClient: client,
	}
}

/*
GetProfileParams contains all the parameters to send to the API endpoint

	for the get profile operation.

	Typically these are written to a http.Request.
*/
type GetProfileParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProfileParams) WithDefaults() *GetProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get profile params
func (o *GetProfileParams) WithTimeout(timeout time.Duration) *GetProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get profile params
func (o *GetProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get profile params
func (o *GetProfileParams) WithContext(ctx context.Context) *GetProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get profile params
func (o *GetProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get profile params
func (o *GetProfileParams) WithHTTPClient(client *http.Client) *GetProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get profile params
func (o *GetProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
