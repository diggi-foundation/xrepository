// Code generated by go-swagger; DO NOT EDIT.

package profil

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

// NewInitResetPasswordParams creates a new InitResetPasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitResetPasswordParams() *InitResetPasswordParams {
	return &InitResetPasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitResetPasswordParamsWithTimeout creates a new InitResetPasswordParams object
// with the ability to set a timeout on a request.
func NewInitResetPasswordParamsWithTimeout(timeout time.Duration) *InitResetPasswordParams {
	return &InitResetPasswordParams{
		timeout: timeout,
	}
}

// NewInitResetPasswordParamsWithContext creates a new InitResetPasswordParams object
// with the ability to set a context for a request.
func NewInitResetPasswordParamsWithContext(ctx context.Context) *InitResetPasswordParams {
	return &InitResetPasswordParams{
		Context: ctx,
	}
}

// NewInitResetPasswordParamsWithHTTPClient creates a new InitResetPasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitResetPasswordParamsWithHTTPClient(client *http.Client) *InitResetPasswordParams {
	return &InitResetPasswordParams{
		HTTPClient: client,
	}
}

/*
InitResetPasswordParams contains all the parameters to send to the API endpoint

	for the init reset password operation.

	Typically these are written to a http.Request.
*/
type InitResetPasswordParams struct {

	// Body.
	Body *models.PasswordResetRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the init reset password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitResetPasswordParams) WithDefaults() *InitResetPasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the init reset password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitResetPasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the init reset password params
func (o *InitResetPasswordParams) WithTimeout(timeout time.Duration) *InitResetPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the init reset password params
func (o *InitResetPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the init reset password params
func (o *InitResetPasswordParams) WithContext(ctx context.Context) *InitResetPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the init reset password params
func (o *InitResetPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the init reset password params
func (o *InitResetPasswordParams) WithHTTPClient(client *http.Client) *InitResetPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the init reset password params
func (o *InitResetPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the init reset password params
func (o *InitResetPasswordParams) WithBody(body *models.PasswordResetRequest) *InitResetPasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the init reset password params
func (o *InitResetPasswordParams) SetBody(body *models.PasswordResetRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InitResetPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
