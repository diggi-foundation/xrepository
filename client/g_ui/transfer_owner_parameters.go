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

// NewTransferOwnerParams creates a new TransferOwnerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTransferOwnerParams() *TransferOwnerParams {
	return &TransferOwnerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTransferOwnerParamsWithTimeout creates a new TransferOwnerParams object
// with the ability to set a timeout on a request.
func NewTransferOwnerParamsWithTimeout(timeout time.Duration) *TransferOwnerParams {
	return &TransferOwnerParams{
		timeout: timeout,
	}
}

// NewTransferOwnerParamsWithContext creates a new TransferOwnerParams object
// with the ability to set a context for a request.
func NewTransferOwnerParamsWithContext(ctx context.Context) *TransferOwnerParams {
	return &TransferOwnerParams{
		Context: ctx,
	}
}

// NewTransferOwnerParamsWithHTTPClient creates a new TransferOwnerParams object
// with the ability to set a custom HTTPClient for a request.
func NewTransferOwnerParamsWithHTTPClient(client *http.Client) *TransferOwnerParams {
	return &TransferOwnerParams{
		HTTPClient: client,
	}
}

/*
TransferOwnerParams contains all the parameters to send to the API endpoint

	for the transfer owner operation.

	Typically these are written to a http.Request.
*/
type TransferOwnerParams struct {

	// Body.
	Body *models.TransferInhaltInfo

	// Kennung.
	Kennung string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the transfer owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransferOwnerParams) WithDefaults() *TransferOwnerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the transfer owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransferOwnerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the transfer owner params
func (o *TransferOwnerParams) WithTimeout(timeout time.Duration) *TransferOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transfer owner params
func (o *TransferOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transfer owner params
func (o *TransferOwnerParams) WithContext(ctx context.Context) *TransferOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transfer owner params
func (o *TransferOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transfer owner params
func (o *TransferOwnerParams) WithHTTPClient(client *http.Client) *TransferOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transfer owner params
func (o *TransferOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the transfer owner params
func (o *TransferOwnerParams) WithBody(body *models.TransferInhaltInfo) *TransferOwnerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the transfer owner params
func (o *TransferOwnerParams) SetBody(body *models.TransferInhaltInfo) {
	o.Body = body
}

// WithKennung adds the kennung to the transfer owner params
func (o *TransferOwnerParams) WithKennung(kennung string) *TransferOwnerParams {
	o.SetKennung(kennung)
	return o
}

// SetKennung adds the kennung to the transfer owner params
func (o *TransferOwnerParams) SetKennung(kennung string) {
	o.Kennung = kennung
}

// WriteToRequest writes these params to a swagger request
func (o *TransferOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param kennung
	if err := r.SetPathParam("kennung", o.Kennung); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
