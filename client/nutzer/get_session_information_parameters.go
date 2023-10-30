// Code generated by go-swagger; DO NOT EDIT.

package nutzer

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

// NewGetSessionInformationParams creates a new GetSessionInformationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSessionInformationParams() *GetSessionInformationParams {
	return &GetSessionInformationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSessionInformationParamsWithTimeout creates a new GetSessionInformationParams object
// with the ability to set a timeout on a request.
func NewGetSessionInformationParamsWithTimeout(timeout time.Duration) *GetSessionInformationParams {
	return &GetSessionInformationParams{
		timeout: timeout,
	}
}

// NewGetSessionInformationParamsWithContext creates a new GetSessionInformationParams object
// with the ability to set a context for a request.
func NewGetSessionInformationParamsWithContext(ctx context.Context) *GetSessionInformationParams {
	return &GetSessionInformationParams{
		Context: ctx,
	}
}

// NewGetSessionInformationParamsWithHTTPClient creates a new GetSessionInformationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSessionInformationParamsWithHTTPClient(client *http.Client) *GetSessionInformationParams {
	return &GetSessionInformationParams{
		HTTPClient: client,
	}
}

/*
GetSessionInformationParams contains all the parameters to send to the API endpoint

	for the get session information operation.

	Typically these are written to a http.Request.
*/
type GetSessionInformationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get session information params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSessionInformationParams) WithDefaults() *GetSessionInformationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get session information params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSessionInformationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get session information params
func (o *GetSessionInformationParams) WithTimeout(timeout time.Duration) *GetSessionInformationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get session information params
func (o *GetSessionInformationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get session information params
func (o *GetSessionInformationParams) WithContext(ctx context.Context) *GetSessionInformationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get session information params
func (o *GetSessionInformationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get session information params
func (o *GetSessionInformationParams) WithHTTPClient(client *http.Client) *GetSessionInformationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get session information params
func (o *GetSessionInformationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSessionInformationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}