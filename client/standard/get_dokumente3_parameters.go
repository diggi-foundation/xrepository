// Code generated by go-swagger; DO NOT EDIT.

package standard

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

// NewGetDokumente3Params creates a new GetDokumente3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDokumente3Params() *GetDokumente3Params {
	return &GetDokumente3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDokumente3ParamsWithTimeout creates a new GetDokumente3Params object
// with the ability to set a timeout on a request.
func NewGetDokumente3ParamsWithTimeout(timeout time.Duration) *GetDokumente3Params {
	return &GetDokumente3Params{
		timeout: timeout,
	}
}

// NewGetDokumente3ParamsWithContext creates a new GetDokumente3Params object
// with the ability to set a context for a request.
func NewGetDokumente3ParamsWithContext(ctx context.Context) *GetDokumente3Params {
	return &GetDokumente3Params{
		Context: ctx,
	}
}

// NewGetDokumente3ParamsWithHTTPClient creates a new GetDokumente3Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetDokumente3ParamsWithHTTPClient(client *http.Client) *GetDokumente3Params {
	return &GetDokumente3Params{
		HTTPClient: client,
	}
}

/*
GetDokumente3Params contains all the parameters to send to the API endpoint

	for the get dokumente 3 operation.

	Typically these are written to a http.Request.
*/
type GetDokumente3Params struct {

	// Kennung.
	Kennung string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dokumente 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDokumente3Params) WithDefaults() *GetDokumente3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dokumente 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDokumente3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dokumente 3 params
func (o *GetDokumente3Params) WithTimeout(timeout time.Duration) *GetDokumente3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dokumente 3 params
func (o *GetDokumente3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dokumente 3 params
func (o *GetDokumente3Params) WithContext(ctx context.Context) *GetDokumente3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dokumente 3 params
func (o *GetDokumente3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dokumente 3 params
func (o *GetDokumente3Params) WithHTTPClient(client *http.Client) *GetDokumente3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dokumente 3 params
func (o *GetDokumente3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKennung adds the kennung to the get dokumente 3 params
func (o *GetDokumente3Params) WithKennung(kennung string) *GetDokumente3Params {
	o.SetKennung(kennung)
	return o
}

// SetKennung adds the kennung to the get dokumente 3 params
func (o *GetDokumente3Params) SetKennung(kennung string) {
	o.Kennung = kennung
}

// WriteToRequest writes these params to a swagger request
func (o *GetDokumente3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
