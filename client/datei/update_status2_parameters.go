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

// NewUpdateStatus2Params creates a new UpdateStatus2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateStatus2Params() *UpdateStatus2Params {
	return &UpdateStatus2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateStatus2ParamsWithTimeout creates a new UpdateStatus2Params object
// with the ability to set a timeout on a request.
func NewUpdateStatus2ParamsWithTimeout(timeout time.Duration) *UpdateStatus2Params {
	return &UpdateStatus2Params{
		timeout: timeout,
	}
}

// NewUpdateStatus2ParamsWithContext creates a new UpdateStatus2Params object
// with the ability to set a context for a request.
func NewUpdateStatus2ParamsWithContext(ctx context.Context) *UpdateStatus2Params {
	return &UpdateStatus2Params{
		Context: ctx,
	}
}

// NewUpdateStatus2ParamsWithHTTPClient creates a new UpdateStatus2Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateStatus2ParamsWithHTTPClient(client *http.Client) *UpdateStatus2Params {
	return &UpdateStatus2Params{
		HTTPClient: client,
	}
}

/*
UpdateStatus2Params contains all the parameters to send to the API endpoint

	for the update status 2 operation.

	Typically these are written to a http.Request.
*/
type UpdateStatus2Params struct {

	// Body.
	Body *models.StatusData

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update status 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateStatus2Params) WithDefaults() *UpdateStatus2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update status 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateStatus2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update status 2 params
func (o *UpdateStatus2Params) WithTimeout(timeout time.Duration) *UpdateStatus2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update status 2 params
func (o *UpdateStatus2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update status 2 params
func (o *UpdateStatus2Params) WithContext(ctx context.Context) *UpdateStatus2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update status 2 params
func (o *UpdateStatus2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update status 2 params
func (o *UpdateStatus2Params) WithHTTPClient(client *http.Client) *UpdateStatus2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update status 2 params
func (o *UpdateStatus2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update status 2 params
func (o *UpdateStatus2Params) WithBody(body *models.StatusData) *UpdateStatus2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update status 2 params
func (o *UpdateStatus2Params) SetBody(body *models.StatusData) {
	o.Body = body
}

// WithUrn adds the urn to the update status 2 params
func (o *UpdateStatus2Params) WithUrn(urn string) *UpdateStatus2Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the update status 2 params
func (o *UpdateStatus2Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateStatus2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
