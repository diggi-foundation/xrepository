// Code generated by go-swagger; DO NOT EDIT.

package redaktion

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

// NewFetchAllParams creates a new FetchAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFetchAllParams() *FetchAllParams {
	return &FetchAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFetchAllParamsWithTimeout creates a new FetchAllParams object
// with the ability to set a timeout on a request.
func NewFetchAllParamsWithTimeout(timeout time.Duration) *FetchAllParams {
	return &FetchAllParams{
		timeout: timeout,
	}
}

// NewFetchAllParamsWithContext creates a new FetchAllParams object
// with the ability to set a context for a request.
func NewFetchAllParamsWithContext(ctx context.Context) *FetchAllParams {
	return &FetchAllParams{
		Context: ctx,
	}
}

// NewFetchAllParamsWithHTTPClient creates a new FetchAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewFetchAllParamsWithHTTPClient(client *http.Client) *FetchAllParams {
	return &FetchAllParams{
		HTTPClient: client,
	}
}

/*
FetchAllParams contains all the parameters to send to the API endpoint

	for the fetch all operation.

	Typically these are written to a http.Request.
*/
type FetchAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fetch all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FetchAllParams) WithDefaults() *FetchAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fetch all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FetchAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fetch all params
func (o *FetchAllParams) WithTimeout(timeout time.Duration) *FetchAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fetch all params
func (o *FetchAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fetch all params
func (o *FetchAllParams) WithContext(ctx context.Context) *FetchAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fetch all params
func (o *FetchAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fetch all params
func (o *FetchAllParams) WithHTTPClient(client *http.Client) *FetchAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fetch all params
func (o *FetchAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FetchAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
