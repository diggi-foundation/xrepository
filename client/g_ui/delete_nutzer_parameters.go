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

// NewDeleteNutzerParams creates a new DeleteNutzerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNutzerParams() *DeleteNutzerParams {
	return &DeleteNutzerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNutzerParamsWithTimeout creates a new DeleteNutzerParams object
// with the ability to set a timeout on a request.
func NewDeleteNutzerParamsWithTimeout(timeout time.Duration) *DeleteNutzerParams {
	return &DeleteNutzerParams{
		timeout: timeout,
	}
}

// NewDeleteNutzerParamsWithContext creates a new DeleteNutzerParams object
// with the ability to set a context for a request.
func NewDeleteNutzerParamsWithContext(ctx context.Context) *DeleteNutzerParams {
	return &DeleteNutzerParams{
		Context: ctx,
	}
}

// NewDeleteNutzerParamsWithHTTPClient creates a new DeleteNutzerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNutzerParamsWithHTTPClient(client *http.Client) *DeleteNutzerParams {
	return &DeleteNutzerParams{
		HTTPClient: client,
	}
}

/*
DeleteNutzerParams contains all the parameters to send to the API endpoint

	for the delete nutzer operation.

	Typically these are written to a http.Request.
*/
type DeleteNutzerParams struct {

	// Kennung.
	Kennung string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete nutzer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNutzerParams) WithDefaults() *DeleteNutzerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete nutzer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNutzerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete nutzer params
func (o *DeleteNutzerParams) WithTimeout(timeout time.Duration) *DeleteNutzerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete nutzer params
func (o *DeleteNutzerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete nutzer params
func (o *DeleteNutzerParams) WithContext(ctx context.Context) *DeleteNutzerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete nutzer params
func (o *DeleteNutzerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete nutzer params
func (o *DeleteNutzerParams) WithHTTPClient(client *http.Client) *DeleteNutzerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete nutzer params
func (o *DeleteNutzerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKennung adds the kennung to the delete nutzer params
func (o *DeleteNutzerParams) WithKennung(kennung string) *DeleteNutzerParams {
	o.SetKennung(kennung)
	return o
}

// SetKennung adds the kennung to the delete nutzer params
func (o *DeleteNutzerParams) SetKennung(kennung string) {
	o.Kennung = kennung
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNutzerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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