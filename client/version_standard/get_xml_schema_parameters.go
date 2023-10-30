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
)

// NewGetXMLSchemaParams creates a new GetXMLSchemaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetXMLSchemaParams() *GetXMLSchemaParams {
	return &GetXMLSchemaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetXMLSchemaParamsWithTimeout creates a new GetXMLSchemaParams object
// with the ability to set a timeout on a request.
func NewGetXMLSchemaParamsWithTimeout(timeout time.Duration) *GetXMLSchemaParams {
	return &GetXMLSchemaParams{
		timeout: timeout,
	}
}

// NewGetXMLSchemaParamsWithContext creates a new GetXMLSchemaParams object
// with the ability to set a context for a request.
func NewGetXMLSchemaParamsWithContext(ctx context.Context) *GetXMLSchemaParams {
	return &GetXMLSchemaParams{
		Context: ctx,
	}
}

// NewGetXMLSchemaParamsWithHTTPClient creates a new GetXMLSchemaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetXMLSchemaParamsWithHTTPClient(client *http.Client) *GetXMLSchemaParams {
	return &GetXMLSchemaParams{
		HTTPClient: client,
	}
}

/*
GetXMLSchemaParams contains all the parameters to send to the API endpoint

	for the get Xml schema operation.

	Typically these are written to a http.Request.
*/
type GetXMLSchemaParams struct {

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Xml schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetXMLSchemaParams) WithDefaults() *GetXMLSchemaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Xml schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetXMLSchemaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Xml schema params
func (o *GetXMLSchemaParams) WithTimeout(timeout time.Duration) *GetXMLSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Xml schema params
func (o *GetXMLSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Xml schema params
func (o *GetXMLSchemaParams) WithContext(ctx context.Context) *GetXMLSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Xml schema params
func (o *GetXMLSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Xml schema params
func (o *GetXMLSchemaParams) WithHTTPClient(client *http.Client) *GetXMLSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Xml schema params
func (o *GetXMLSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUrn adds the urn to the get Xml schema params
func (o *GetXMLSchemaParams) WithUrn(urn string) *GetXMLSchemaParams {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the get Xml schema params
func (o *GetXMLSchemaParams) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *GetXMLSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param urn
	if err := r.SetPathParam("urn", o.Urn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
