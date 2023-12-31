// Code generated by go-swagger; DO NOT EDIT.

package codeliste

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

// NewTransferOwnershipParams creates a new TransferOwnershipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTransferOwnershipParams() *TransferOwnershipParams {
	return &TransferOwnershipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTransferOwnershipParamsWithTimeout creates a new TransferOwnershipParams object
// with the ability to set a timeout on a request.
func NewTransferOwnershipParamsWithTimeout(timeout time.Duration) *TransferOwnershipParams {
	return &TransferOwnershipParams{
		timeout: timeout,
	}
}

// NewTransferOwnershipParamsWithContext creates a new TransferOwnershipParams object
// with the ability to set a context for a request.
func NewTransferOwnershipParamsWithContext(ctx context.Context) *TransferOwnershipParams {
	return &TransferOwnershipParams{
		Context: ctx,
	}
}

// NewTransferOwnershipParamsWithHTTPClient creates a new TransferOwnershipParams object
// with the ability to set a custom HTTPClient for a request.
func NewTransferOwnershipParamsWithHTTPClient(client *http.Client) *TransferOwnershipParams {
	return &TransferOwnershipParams{
		HTTPClient: client,
	}
}

/*
TransferOwnershipParams contains all the parameters to send to the API endpoint

	for the transfer ownership operation.

	Typically these are written to a http.Request.
*/
type TransferOwnershipParams struct {

	// Body.
	Body *models.TransferOwnerData

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the transfer ownership params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransferOwnershipParams) WithDefaults() *TransferOwnershipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the transfer ownership params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransferOwnershipParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the transfer ownership params
func (o *TransferOwnershipParams) WithTimeout(timeout time.Duration) *TransferOwnershipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transfer ownership params
func (o *TransferOwnershipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transfer ownership params
func (o *TransferOwnershipParams) WithContext(ctx context.Context) *TransferOwnershipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transfer ownership params
func (o *TransferOwnershipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transfer ownership params
func (o *TransferOwnershipParams) WithHTTPClient(client *http.Client) *TransferOwnershipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transfer ownership params
func (o *TransferOwnershipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the transfer ownership params
func (o *TransferOwnershipParams) WithBody(body *models.TransferOwnerData) *TransferOwnershipParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the transfer ownership params
func (o *TransferOwnershipParams) SetBody(body *models.TransferOwnerData) {
	o.Body = body
}

// WithUrn adds the urn to the transfer ownership params
func (o *TransferOwnershipParams) WithUrn(urn string) *TransferOwnershipParams {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the transfer ownership params
func (o *TransferOwnershipParams) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *TransferOwnershipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
