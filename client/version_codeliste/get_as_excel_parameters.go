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

// NewGetAsExcelParams creates a new GetAsExcelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAsExcelParams() *GetAsExcelParams {
	return &GetAsExcelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAsExcelParamsWithTimeout creates a new GetAsExcelParams object
// with the ability to set a timeout on a request.
func NewGetAsExcelParamsWithTimeout(timeout time.Duration) *GetAsExcelParams {
	return &GetAsExcelParams{
		timeout: timeout,
	}
}

// NewGetAsExcelParamsWithContext creates a new GetAsExcelParams object
// with the ability to set a context for a request.
func NewGetAsExcelParamsWithContext(ctx context.Context) *GetAsExcelParams {
	return &GetAsExcelParams{
		Context: ctx,
	}
}

// NewGetAsExcelParamsWithHTTPClient creates a new GetAsExcelParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAsExcelParamsWithHTTPClient(client *http.Client) *GetAsExcelParams {
	return &GetAsExcelParams{
		HTTPClient: client,
	}
}

/*
GetAsExcelParams contains all the parameters to send to the API endpoint

	for the get as excel operation.

	Typically these are written to a http.Request.
*/
type GetAsExcelParams struct {

	// Kennung.
	Kennung string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get as excel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAsExcelParams) WithDefaults() *GetAsExcelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get as excel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAsExcelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get as excel params
func (o *GetAsExcelParams) WithTimeout(timeout time.Duration) *GetAsExcelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get as excel params
func (o *GetAsExcelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get as excel params
func (o *GetAsExcelParams) WithContext(ctx context.Context) *GetAsExcelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get as excel params
func (o *GetAsExcelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get as excel params
func (o *GetAsExcelParams) WithHTTPClient(client *http.Client) *GetAsExcelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get as excel params
func (o *GetAsExcelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKennung adds the kennung to the get as excel params
func (o *GetAsExcelParams) WithKennung(kennung string) *GetAsExcelParams {
	o.SetKennung(kennung)
	return o
}

// SetKennung adds the kennung to the get as excel params
func (o *GetAsExcelParams) SetKennung(kennung string) {
	o.Kennung = kennung
}

// WithName adds the name to the get as excel params
func (o *GetAsExcelParams) WithName(name string) *GetAsExcelParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get as excel params
func (o *GetAsExcelParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetAsExcelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param kennung
	if err := r.SetPathParam("kennung", o.Kennung); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
