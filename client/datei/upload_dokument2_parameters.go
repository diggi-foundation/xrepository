// Code generated by go-swagger; DO NOT EDIT.

package datei

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

// NewUploadDokument2Params creates a new UploadDokument2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadDokument2Params() *UploadDokument2Params {
	return &UploadDokument2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadDokument2ParamsWithTimeout creates a new UploadDokument2Params object
// with the ability to set a timeout on a request.
func NewUploadDokument2ParamsWithTimeout(timeout time.Duration) *UploadDokument2Params {
	return &UploadDokument2Params{
		timeout: timeout,
	}
}

// NewUploadDokument2ParamsWithContext creates a new UploadDokument2Params object
// with the ability to set a context for a request.
func NewUploadDokument2ParamsWithContext(ctx context.Context) *UploadDokument2Params {
	return &UploadDokument2Params{
		Context: ctx,
	}
}

// NewUploadDokument2ParamsWithHTTPClient creates a new UploadDokument2Params object
// with the ability to set a custom HTTPClient for a request.
func NewUploadDokument2ParamsWithHTTPClient(client *http.Client) *UploadDokument2Params {
	return &UploadDokument2Params{
		HTTPClient: client,
	}
}

/*
UploadDokument2Params contains all the parameters to send to the API endpoint

	for the upload dokument 2 operation.

	Typically these are written to a http.Request.
*/
type UploadDokument2Params struct {

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

// WithDefaults hydrates default values in the upload dokument 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadDokument2Params) WithDefaults() *UploadDokument2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload dokument 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadDokument2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload dokument 2 params
func (o *UploadDokument2Params) WithTimeout(timeout time.Duration) *UploadDokument2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload dokument 2 params
func (o *UploadDokument2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload dokument 2 params
func (o *UploadDokument2Params) WithContext(ctx context.Context) *UploadDokument2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload dokument 2 params
func (o *UploadDokument2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload dokument 2 params
func (o *UploadDokument2Params) WithHTTPClient(client *http.Client) *UploadDokument2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload dokument 2 params
func (o *UploadDokument2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocument adds the document to the upload dokument 2 params
func (o *UploadDokument2Params) WithDocument(document string) *UploadDokument2Params {
	o.SetDocument(document)
	return o
}

// SetDocument adds the document to the upload dokument 2 params
func (o *UploadDokument2Params) SetDocument(document string) {
	o.Document = document
}

// WithFiles adds the files to the upload dokument 2 params
func (o *UploadDokument2Params) WithFiles(files runtime.NamedReadCloser) *UploadDokument2Params {
	o.SetFiles(files)
	return o
}

// SetFiles adds the files to the upload dokument 2 params
func (o *UploadDokument2Params) SetFiles(files runtime.NamedReadCloser) {
	o.Files = files
}

// WithUrn adds the urn to the upload dokument 2 params
func (o *UploadDokument2Params) WithUrn(urn string) *UploadDokument2Params {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the upload dokument 2 params
func (o *UploadDokument2Params) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *UploadDokument2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
