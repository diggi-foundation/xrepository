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

	"github.com/diggi-foundation/xrepository/models"
)

// NewSearch1Params creates a new Search1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearch1Params() *Search1Params {
	return &Search1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearch1ParamsWithTimeout creates a new Search1Params object
// with the ability to set a timeout on a request.
func NewSearch1ParamsWithTimeout(timeout time.Duration) *Search1Params {
	return &Search1Params{
		timeout: timeout,
	}
}

// NewSearch1ParamsWithContext creates a new Search1Params object
// with the ability to set a context for a request.
func NewSearch1ParamsWithContext(ctx context.Context) *Search1Params {
	return &Search1Params{
		Context: ctx,
	}
}

// NewSearch1ParamsWithHTTPClient creates a new Search1Params object
// with the ability to set a custom HTTPClient for a request.
func NewSearch1ParamsWithHTTPClient(client *http.Client) *Search1Params {
	return &Search1Params{
		HTTPClient: client,
	}
}

/*
Search1Params contains all the parameters to send to the API endpoint

	for the search 1 operation.

	Typically these are written to a http.Request.
*/
type Search1Params struct {

	// Body.
	Body *models.ObjektSearch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Search1Params) WithDefaults() *Search1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Search1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search 1 params
func (o *Search1Params) WithTimeout(timeout time.Duration) *Search1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search 1 params
func (o *Search1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search 1 params
func (o *Search1Params) WithContext(ctx context.Context) *Search1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search 1 params
func (o *Search1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search 1 params
func (o *Search1Params) WithHTTPClient(client *http.Client) *Search1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search 1 params
func (o *Search1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search 1 params
func (o *Search1Params) WithBody(body *models.ObjektSearch) *Search1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search 1 params
func (o *Search1Params) SetBody(body *models.ObjektSearch) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *Search1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
