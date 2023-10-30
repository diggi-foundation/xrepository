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

	"github.com/diggi.foundation/xrepository/models"
)

// NewTransferOwnership3Params creates a new TransferOwnership3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTransferOwnership3Params() *TransferOwnership3Params {
	return &TransferOwnership3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewTransferOwnership3ParamsWithTimeout creates a new TransferOwnership3Params object
// with the ability to set a timeout on a request.
func NewTransferOwnership3ParamsWithTimeout(timeout time.Duration) *TransferOwnership3Params {
	return &TransferOwnership3Params{
		timeout: timeout,
	}
}

// NewTransferOwnership3ParamsWithContext creates a new TransferOwnership3Params object
// with the ability to set a context for a request.
func NewTransferOwnership3ParamsWithContext(ctx context.Context) *TransferOwnership3Params {
	return &TransferOwnership3Params{
		Context: ctx,
	}
}

// NewTransferOwnership3ParamsWithHTTPClient creates a new TransferOwnership3Params object
// with the ability to set a custom HTTPClient for a request.
func NewTransferOwnership3ParamsWithHTTPClient(client *http.Client) *TransferOwnership3Params {
	return &TransferOwnership3Params{
		HTTPClient: client,
	}
}

/*
TransferOwnership3Params contains all the parameters to send to the API endpoint

	for the transfer ownership 3 operation.

	Typically these are written to a http.Request.
*/
type TransferOwnership3Params struct {

	// Body.
	Body *models.TransferOwnerData

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the transfer ownership 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransferOwnership3Params) WithDefaults() *TransferOwnership3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the transfer ownership 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransferOwnership3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the transfer ownership 3 params
func (o *TransferOwnership3Params) WithTimeout(timeout time.Duration) *TransferOwnership3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transfer ownership 3 params
func (o *TransferOwnership3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transfer ownership 3 params
func (o *TransferOwnership3Params) WithContext(ctx context.Context) *TransferOwnership3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transfer ownership 3 params
func (o *TransferOwnership3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transfer ownership 3 params
func (o *TransferOwnership3Params) WithHTTPClient(client *http.Client) *TransferOwnership3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transfer ownership 3 params
func (o *TransferOwnership3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the transfer ownership 3 params
func (o *TransferOwnership3Params) WithBody(body *models.TransferOwnerData) *TransferOwnership3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the transfer ownership 3 params
func (o *TransferOwnership3Params) SetBody(body *models.TransferOwnerData) {
	o.Body = body
}

// WithUrn adds the urn to the transfer ownership 3 params
func (o *TransferOwnership3Params) WithUrn(urn string) *TransferOwnership3Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the transfer ownership 3 params
func (o *TransferOwnership3Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *TransferOwnership3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
