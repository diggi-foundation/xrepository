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

// NewUploadDokument3Params creates a new UploadDokument3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadDokument3Params() *UploadDokument3Params {
	return &UploadDokument3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadDokument3ParamsWithTimeout creates a new UploadDokument3Params object
// with the ability to set a timeout on a request.
func NewUploadDokument3ParamsWithTimeout(timeout time.Duration) *UploadDokument3Params {
	return &UploadDokument3Params{
		timeout: timeout,
	}
}

// NewUploadDokument3ParamsWithContext creates a new UploadDokument3Params object
// with the ability to set a context for a request.
func NewUploadDokument3ParamsWithContext(ctx context.Context) *UploadDokument3Params {
	return &UploadDokument3Params{
		Context: ctx,
	}
}

// NewUploadDokument3ParamsWithHTTPClient creates a new UploadDokument3Params object
// with the ability to set a custom HTTPClient for a request.
func NewUploadDokument3ParamsWithHTTPClient(client *http.Client) *UploadDokument3Params {
	return &UploadDokument3Params{
		HTTPClient: client,
	}
}

/*
UploadDokument3Params contains all the parameters to send to the API endpoint

	for the upload dokument 3 operation.

	Typically these are written to a http.Request.
*/
type UploadDokument3Params struct {

	/* Document.

	   Dokumentinformatation
	*/
	Document string

	/* Files.

	   .uml Dateien mit Fachmodell
	*/
	Files runtime.NamedReadCloser

	// Urn.
	Urn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload dokument 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadDokument3Params) WithDefaults() *UploadDokument3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload dokument 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadDokument3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload dokument 3 params
func (o *UploadDokument3Params) WithTimeout(timeout time.Duration) *UploadDokument3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload dokument 3 params
func (o *UploadDokument3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload dokument 3 params
func (o *UploadDokument3Params) WithContext(ctx context.Context) *UploadDokument3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload dokument 3 params
func (o *UploadDokument3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload dokument 3 params
func (o *UploadDokument3Params) WithHTTPClient(client *http.Client) *UploadDokument3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload dokument 3 params
func (o *UploadDokument3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocument adds the document to the upload dokument 3 params
func (o *UploadDokument3Params) WithDocument(document string) *UploadDokument3Params {
	o.SetDocument(document)
	return o
}

// SetDocument adds the document to the upload dokument 3 params
func (o *UploadDokument3Params) SetDocument(document string) {
	o.Document = document
}

// WithFiles adds the files to the upload dokument 3 params
func (o *UploadDokument3Params) WithFiles(files runtime.NamedReadCloser) *UploadDokument3Params {
	o.SetFiles(files)
	return o
}

// SetFiles adds the files to the upload dokument 3 params
func (o *UploadDokument3Params) SetFiles(files runtime.NamedReadCloser) {
	o.Files = files
}

// WithUrn adds the urn to the upload dokument 3 params
func (o *UploadDokument3Params) WithUrn(urn string) *UploadDokument3Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the upload dokument 3 params
func (o *UploadDokument3Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *UploadDokument3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param document
	frDocument := o.Document
	fDocument := frDocument
	if fDocument != "" {
		if err := r.SetFormParam("document", fDocument); err != nil {
			return err
		}
	}
	// form file param files
	if err := r.SetFileParam("files", o.Files); err != nil {
		return err
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
