// Code generated by go-swagger; DO NOT EDIT.

package g_ui

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

// NewCreateTokenParams creates a new CreateTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTokenParams() *CreateTokenParams {
	return &CreateTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTokenParamsWithTimeout creates a new CreateTokenParams object
// with the ability to set a timeout on a request.
func NewCreateTokenParamsWithTimeout(timeout time.Duration) *CreateTokenParams {
	return &CreateTokenParams{
		timeout: timeout,
	}
}

// NewCreateTokenParamsWithContext creates a new CreateTokenParams object
// with the ability to set a context for a request.
func NewCreateTokenParamsWithContext(ctx context.Context) *CreateTokenParams {
	return &CreateTokenParams{
		Context: ctx,
	}
}

// NewCreateTokenParamsWithHTTPClient creates a new CreateTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTokenParamsWithHTTPClient(client *http.Client) *CreateTokenParams {
	return &CreateTokenParams{
		HTTPClient: client,
	}
}

/*
CreateTokenParams contains all the parameters to send to the API endpoint

	for the create token operation.

	Typically these are written to a http.Request.
*/
type CreateTokenParams struct {

	// Kennung.
	Kennung string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTokenParams) WithDefaults() *CreateTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create token params
func (o *CreateTokenParams) WithTimeout(timeout time.Duration) *CreateTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create token params
func (o *CreateTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create token params
func (o *CreateTokenParams) WithContext(ctx context.Context) *CreateTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create token params
func (o *CreateTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create token params
func (o *CreateTokenParams) WithHTTPClient(client *http.Client) *CreateTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create token params
func (o *CreateTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKennung adds the kennung to the create token params
func (o *CreateTokenParams) WithKennung(kennung string) *CreateTokenParams {
	o.SetKennung(kennung)
	return o
}

// SetKennung adds the kennung to the create token params
func (o *CreateTokenParams) SetKennung(kennung string) {
	o.Kennung = kennung
}

// WithName adds the name to the create token params
func (o *CreateTokenParams) WithName(name string) *CreateTokenParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the create token params
func (o *CreateTokenParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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