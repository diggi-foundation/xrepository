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

// NewZertifizierungUpdateParams creates a new ZertifizierungUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewZertifizierungUpdateParams() *ZertifizierungUpdateParams {
	return &ZertifizierungUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewZertifizierungUpdateParamsWithTimeout creates a new ZertifizierungUpdateParams object
// with the ability to set a timeout on a request.
func NewZertifizierungUpdateParamsWithTimeout(timeout time.Duration) *ZertifizierungUpdateParams {
	return &ZertifizierungUpdateParams{
		timeout: timeout,
	}
}

// NewZertifizierungUpdateParamsWithContext creates a new ZertifizierungUpdateParams object
// with the ability to set a context for a request.
func NewZertifizierungUpdateParamsWithContext(ctx context.Context) *ZertifizierungUpdateParams {
	return &ZertifizierungUpdateParams{
		Context: ctx,
	}
}

// NewZertifizierungUpdateParamsWithHTTPClient creates a new ZertifizierungUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewZertifizierungUpdateParamsWithHTTPClient(client *http.Client) *ZertifizierungUpdateParams {
	return &ZertifizierungUpdateParams{
		HTTPClient: client,
	}
}

/*
ZertifizierungUpdateParams contains all the parameters to send to the API endpoint

	for the zertifizierung update operation.

	Typically these are written to a http.Request.
*/
type ZertifizierungUpdateParams struct {

	// Body.
	Body *models.ZertifizierungData

	// Kennung.
	Kennung string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the zertifizierung update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ZertifizierungUpdateParams) WithDefaults() *ZertifizierungUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the zertifizierung update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ZertifizierungUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the zertifizierung update params
func (o *ZertifizierungUpdateParams) WithTimeout(timeout time.Duration) *ZertifizierungUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the zertifizierung update params
func (o *ZertifizierungUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the zertifizierung update params
func (o *ZertifizierungUpdateParams) WithContext(ctx context.Context) *ZertifizierungUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the zertifizierung update params
func (o *ZertifizierungUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the zertifizierung update params
func (o *ZertifizierungUpdateParams) WithHTTPClient(client *http.Client) *ZertifizierungUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the zertifizierung update params
func (o *ZertifizierungUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the zertifizierung update params
func (o *ZertifizierungUpdateParams) WithBody(body *models.ZertifizierungData) *ZertifizierungUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the zertifizierung update params
func (o *ZertifizierungUpdateParams) SetBody(body *models.ZertifizierungData) {
	o.Body = body
}

// WithKennung adds the kennung to the zertifizierung update params
func (o *ZertifizierungUpdateParams) WithKennung(kennung string) *ZertifizierungUpdateParams {
	o.SetKennung(kennung)
	return o
}

// SetKennung adds the kennung to the zertifizierung update params
func (o *ZertifizierungUpdateParams) SetKennung(kennung string) {
	o.Kennung = kennung
}

// WriteToRequest writes these params to a swagger request
func (o *ZertifizierungUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
