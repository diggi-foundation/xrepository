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

	"github.com/diggi-foundation/xrepository/models"
)

// NewAddVersionCodelisteParams creates a new AddVersionCodelisteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddVersionCodelisteParams() *AddVersionCodelisteParams {
	return &AddVersionCodelisteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddVersionCodelisteParamsWithTimeout creates a new AddVersionCodelisteParams object
// with the ability to set a timeout on a request.
func NewAddVersionCodelisteParamsWithTimeout(timeout time.Duration) *AddVersionCodelisteParams {
	return &AddVersionCodelisteParams{
		timeout: timeout,
	}
}

// NewAddVersionCodelisteParamsWithContext creates a new AddVersionCodelisteParams object
// with the ability to set a context for a request.
func NewAddVersionCodelisteParamsWithContext(ctx context.Context) *AddVersionCodelisteParams {
	return &AddVersionCodelisteParams{
		Context: ctx,
	}
}

// NewAddVersionCodelisteParamsWithHTTPClient creates a new AddVersionCodelisteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddVersionCodelisteParamsWithHTTPClient(client *http.Client) *AddVersionCodelisteParams {
	return &AddVersionCodelisteParams{
		HTTPClient: client,
	}
}

/*
AddVersionCodelisteParams contains all the parameters to send to the API endpoint

	for the add version codeliste operation.

	Typically these are written to a http.Request.
*/
type AddVersionCodelisteParams struct {

	// Body.
	Body *models.MultipartFormDataInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add version codeliste params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVersionCodelisteParams) WithDefaults() *AddVersionCodelisteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add version codeliste params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVersionCodelisteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add version codeliste params
func (o *AddVersionCodelisteParams) WithTimeout(timeout time.Duration) *AddVersionCodelisteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add version codeliste params
func (o *AddVersionCodelisteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add version codeliste params
func (o *AddVersionCodelisteParams) WithContext(ctx context.Context) *AddVersionCodelisteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add version codeliste params
func (o *AddVersionCodelisteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add version codeliste params
func (o *AddVersionCodelisteParams) WithHTTPClient(client *http.Client) *AddVersionCodelisteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add version codeliste params
func (o *AddVersionCodelisteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add version codeliste params
func (o *AddVersionCodelisteParams) WithBody(body *models.MultipartFormDataInput) *AddVersionCodelisteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add version codeliste params
func (o *AddVersionCodelisteParams) SetBody(body *models.MultipartFormDataInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddVersionCodelisteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
