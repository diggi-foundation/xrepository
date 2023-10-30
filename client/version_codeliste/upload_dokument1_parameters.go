// Code generated by go-swagger; DO NOT EDIT.

package version_codeliste

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

// NewUploadDokument1Params creates a new UploadDokument1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadDokument1Params() *UploadDokument1Params {
	return &UploadDokument1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadDokument1ParamsWithTimeout creates a new UploadDokument1Params object
// with the ability to set a timeout on a request.
func NewUploadDokument1ParamsWithTimeout(timeout time.Duration) *UploadDokument1Params {
	return &UploadDokument1Params{
		timeout: timeout,
	}
}

// NewUploadDokument1ParamsWithContext creates a new UploadDokument1Params object
// with the ability to set a context for a request.
func NewUploadDokument1ParamsWithContext(ctx context.Context) *UploadDokument1Params {
	return &UploadDokument1Params{
		Context: ctx,
	}
}

// NewUploadDokument1ParamsWithHTTPClient creates a new UploadDokument1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUploadDokument1ParamsWithHTTPClient(client *http.Client) *UploadDokument1Params {
	return &UploadDokument1Params{
		HTTPClient: client,
	}
}

/*
UploadDokument1Params contains all the parameters to send to the API endpoint

	for the upload dokument 1 operation.

	Typically these are written to a http.Request.
*/
type UploadDokument1Params struct {

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

// WithDefaults hydrates default values in the upload dokument 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadDokument1Params) WithDefaults() *UploadDokument1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload dokument 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadDokument1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload dokument 1 params
func (o *UploadDokument1Params) WithTimeout(timeout time.Duration) *UploadDokument1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload dokument 1 params
func (o *UploadDokument1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload dokument 1 params
func (o *UploadDokument1Params) WithContext(ctx context.Context) *UploadDokument1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload dokument 1 params
func (o *UploadDokument1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload dokument 1 params
func (o *UploadDokument1Params) WithHTTPClient(client *http.Client) *UploadDokument1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload dokument 1 params
func (o *UploadDokument1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocument adds the document to the upload dokument 1 params
func (o *UploadDokument1Params) WithDocument(document string) *UploadDokument1Params {
	o.SetDocument(document)
	return o
}

// SetDocument adds the document to the upload dokument 1 params
func (o *UploadDokument1Params) SetDocument(document string) {
	o.Document = document
}

// WithFiles adds the files to the upload dokument 1 params
func (o *UploadDokument1Params) WithFiles(files runtime.NamedReadCloser) *UploadDokument1Params {
	o.SetFiles(files)
	return o
}

// SetFiles adds the files to the upload dokument 1 params
func (o *UploadDokument1Params) SetFiles(files runtime.NamedReadCloser) {
	o.Files = files
}

// WithUrn adds the urn to the upload dokument 1 params
func (o *UploadDokument1Params) WithUrn(urn string) *UploadDokument1Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the upload dokument 1 params
func (o *UploadDokument1Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *UploadDokument1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
