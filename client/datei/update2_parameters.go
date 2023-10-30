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

	"github.com/diggi-foundation/xrepository/models"
)

// NewUpdate2Params creates a new Update2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdate2Params() *Update2Params {
	return &Update2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdate2ParamsWithTimeout creates a new Update2Params object
// with the ability to set a timeout on a request.
func NewUpdate2ParamsWithTimeout(timeout time.Duration) *Update2Params {
	return &Update2Params{
		timeout: timeout,
	}
}

// NewUpdate2ParamsWithContext creates a new Update2Params object
// with the ability to set a context for a request.
func NewUpdate2ParamsWithContext(ctx context.Context) *Update2Params {
	return &Update2Params{
		Context: ctx,
	}
}

// NewUpdate2ParamsWithHTTPClient creates a new Update2Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdate2ParamsWithHTTPClient(client *http.Client) *Update2Params {
	return &Update2Params{
		HTTPClient: client,
	}
}

/*
Update2Params contains all the parameters to send to the API endpoint

	for the update 2 operation.

	Typically these are written to a http.Request.
*/
type Update2Params struct {

	// Body.
	Body *models.UpdateData

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Update2Params) WithDefaults() *Update2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Update2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update 2 params
func (o *Update2Params) WithTimeout(timeout time.Duration) *Update2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update 2 params
func (o *Update2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update 2 params
func (o *Update2Params) WithContext(ctx context.Context) *Update2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update 2 params
func (o *Update2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update 2 params
func (o *Update2Params) WithHTTPClient(client *http.Client) *Update2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update 2 params
func (o *Update2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update 2 params
func (o *Update2Params) WithBody(body *models.UpdateData) *Update2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update 2 params
func (o *Update2Params) SetBody(body *models.UpdateData) {
	o.Body = body
}

// WithUrn adds the urn to the update 2 params
func (o *Update2Params) WithUrn(urn string) *Update2Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the update 2 params
func (o *Update2Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *Update2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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
