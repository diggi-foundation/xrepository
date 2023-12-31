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

// NewUpdate3Params creates a new Update3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdate3Params() *Update3Params {
	return &Update3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdate3ParamsWithTimeout creates a new Update3Params object
// with the ability to set a timeout on a request.
func NewUpdate3ParamsWithTimeout(timeout time.Duration) *Update3Params {
	return &Update3Params{
		timeout: timeout,
	}
}

// NewUpdate3ParamsWithContext creates a new Update3Params object
// with the ability to set a context for a request.
func NewUpdate3ParamsWithContext(ctx context.Context) *Update3Params {
	return &Update3Params{
		Context: ctx,
	}
}

// NewUpdate3ParamsWithHTTPClient creates a new Update3Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdate3ParamsWithHTTPClient(client *http.Client) *Update3Params {
	return &Update3Params{
		HTTPClient: client,
	}
}

/*
Update3Params contains all the parameters to send to the API endpoint

	for the update 3 operation.

	Typically these are written to a http.Request.
*/
type Update3Params struct {

	/* File.

	   Bild (Binärdaten)
	*/
	File runtime.NamedReadCloser

	/* Model.

	   Bild-Modell als JSON
	*/
	Model string

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Update3Params) WithDefaults() *Update3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Update3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update 3 params
func (o *Update3Params) WithTimeout(timeout time.Duration) *Update3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update 3 params
func (o *Update3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update 3 params
func (o *Update3Params) WithContext(ctx context.Context) *Update3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update 3 params
func (o *Update3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update 3 params
func (o *Update3Params) WithHTTPClient(client *http.Client) *Update3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update 3 params
func (o *Update3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the update 3 params
func (o *Update3Params) WithFile(file runtime.NamedReadCloser) *Update3Params {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the update 3 params
func (o *Update3Params) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithModel adds the model to the update 3 params
func (o *Update3Params) WithModel(model string) *Update3Params {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the update 3 params
func (o *Update3Params) SetModel(model string) {
	o.Model = model
}

// WithUrn adds the urn to the update 3 params
func (o *Update3Params) WithUrn(urn string) *Update3Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the update 3 params
func (o *Update3Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *Update3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	// form param model
	frModel := o.Model
	fModel := frModel
	if fModel != "" {
		if err := r.SetFormParam("model", fModel); err != nil {
			return err
		}
	}

	// path param urn
	if err := r.SetPathParam("urn", o.Urn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
