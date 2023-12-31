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

	"github.com/diggi-foundation/xrepository/models"
)

// NewUpdate8Params creates a new Update8Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdate8Params() *Update8Params {
	return &Update8Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdate8ParamsWithTimeout creates a new Update8Params object
// with the ability to set a timeout on a request.
func NewUpdate8ParamsWithTimeout(timeout time.Duration) *Update8Params {
	return &Update8Params{
		timeout: timeout,
	}
}

// NewUpdate8ParamsWithContext creates a new Update8Params object
// with the ability to set a context for a request.
func NewUpdate8ParamsWithContext(ctx context.Context) *Update8Params {
	return &Update8Params{
		Context: ctx,
	}
}

// NewUpdate8ParamsWithHTTPClient creates a new Update8Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdate8ParamsWithHTTPClient(client *http.Client) *Update8Params {
	return &Update8Params{
		HTTPClient: client,
	}
}

/*
Update8Params contains all the parameters to send to the API endpoint

	for the update 8 operation.

	Typically these are written to a http.Request.
*/
type Update8Params struct {

	// Body.
	Body *models.UpdateData

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update 8 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Update8Params) WithDefaults() *Update8Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update 8 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Update8Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update 8 params
func (o *Update8Params) WithTimeout(timeout time.Duration) *Update8Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update 8 params
func (o *Update8Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update 8 params
func (o *Update8Params) WithContext(ctx context.Context) *Update8Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update 8 params
func (o *Update8Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update 8 params
func (o *Update8Params) WithHTTPClient(client *http.Client) *Update8Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update 8 params
func (o *Update8Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update 8 params
func (o *Update8Params) WithBody(body *models.UpdateData) *Update8Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update 8 params
func (o *Update8Params) SetBody(body *models.UpdateData) {
	o.Body = body
}

// WithUrn adds the urn to the update 8 params
func (o *Update8Params) WithUrn(urn string) *Update8Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the update 8 params
func (o *Update8Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *Update8Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
