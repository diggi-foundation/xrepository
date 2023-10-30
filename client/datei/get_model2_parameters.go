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

// NewGetModel2Params creates a new GetModel2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetModel2Params() *GetModel2Params {
	return &GetModel2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetModel2ParamsWithTimeout creates a new GetModel2Params object
// with the ability to set a timeout on a request.
func NewGetModel2ParamsWithTimeout(timeout time.Duration) *GetModel2Params {
	return &GetModel2Params{
		timeout: timeout,
	}
}

// NewGetModel2ParamsWithContext creates a new GetModel2Params object
// with the ability to set a context for a request.
func NewGetModel2ParamsWithContext(ctx context.Context) *GetModel2Params {
	return &GetModel2Params{
		Context: ctx,
	}
}

// NewGetModel2ParamsWithHTTPClient creates a new GetModel2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetModel2ParamsWithHTTPClient(client *http.Client) *GetModel2Params {
	return &GetModel2Params{
		HTTPClient: client,
	}
}

/*
GetModel2Params contains all the parameters to send to the API endpoint

	for the get model 2 operation.

	Typically these are written to a http.Request.
*/
type GetModel2Params struct {

	// Kennung.
	Kennung string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get model 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetModel2Params) WithDefaults() *GetModel2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get model 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetModel2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get model 2 params
func (o *GetModel2Params) WithTimeout(timeout time.Duration) *GetModel2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get model 2 params
func (o *GetModel2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get model 2 params
func (o *GetModel2Params) WithContext(ctx context.Context) *GetModel2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get model 2 params
func (o *GetModel2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get model 2 params
func (o *GetModel2Params) WithHTTPClient(client *http.Client) *GetModel2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get model 2 params
func (o *GetModel2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKennung adds the kennung to the get model 2 params
func (o *GetModel2Params) WithKennung(kennung string) *GetModel2Params {
	o.SetKennung(kennung)
	return o
}

// SetKennung adds the kennung to the get model 2 params
func (o *GetModel2Params) SetKennung(kennung string) {
	o.Kennung = kennung
}

// WriteToRequest writes these params to a swagger request
func (o *GetModel2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param kennung
	if err := r.SetPathParam("kennung", o.Kennung); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}