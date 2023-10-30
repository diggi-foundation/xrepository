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
)

// NewUploadDokumentParams creates a new UploadDokumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadDokumentParams() *UploadDokumentParams {
	return &UploadDokumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadDokumentParamsWithTimeout creates a new UploadDokumentParams object
// with the ability to set a timeout on a request.
func NewUploadDokumentParamsWithTimeout(timeout time.Duration) *UploadDokumentParams {
	return &UploadDokumentParams{
		timeout: timeout,
	}
}

// NewUploadDokumentParamsWithContext creates a new UploadDokumentParams object
// with the ability to set a context for a request.
func NewUploadDokumentParamsWithContext(ctx context.Context) *UploadDokumentParams {
	return &UploadDokumentParams{
		Context: ctx,
	}
}

// NewUploadDokumentParamsWithHTTPClient creates a new UploadDokumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadDokumentParamsWithHTTPClient(client *http.Client) *UploadDokumentParams {
	return &UploadDokumentParams{
		HTTPClient: client,
	}
}

/*
UploadDokumentParams contains all the parameters to send to the API endpoint

	for the upload dokument operation.

	Typically these are written to a http.Request.
*/
type UploadDokumentParams struct {

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

// WithDefaults hydrates default values in the upload dokument params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadDokumentParams) WithDefaults() *UploadDokumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload dokument params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadDokumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload dokument params
func (o *UploadDokumentParams) WithTimeout(timeout time.Duration) *UploadDokumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload dokument params
func (o *UploadDokumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload dokument params
func (o *UploadDokumentParams) WithContext(ctx context.Context) *UploadDokumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload dokument params
func (o *UploadDokumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload dokument params
func (o *UploadDokumentParams) WithHTTPClient(client *http.Client) *UploadDokumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload dokument params
func (o *UploadDokumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocument adds the document to the upload dokument params
func (o *UploadDokumentParams) WithDocument(document string) *UploadDokumentParams {
	o.SetDocument(document)
	return o
}

// SetDocument adds the document to the upload dokument params
func (o *UploadDokumentParams) SetDocument(document string) {
	o.Document = document
}

// WithFiles adds the files to the upload dokument params
func (o *UploadDokumentParams) WithFiles(files runtime.NamedReadCloser) *UploadDokumentParams {
	o.SetFiles(files)
	return o
}

// SetFiles adds the files to the upload dokument params
func (o *UploadDokumentParams) SetFiles(files runtime.NamedReadCloser) {
	o.Files = files
}

// WithUrn adds the urn to the upload dokument params
func (o *UploadDokumentParams) WithUrn(urn string) *UploadDokumentParams {
	o.SetUrn(urn)
	return o
}

// SetUrn adds the urn to the upload dokument params
func (o *UploadDokumentParams) SetUrn(urn string) {
	o.Urn = urn
}

// WriteToRequest writes these params to a swagger request
func (o *UploadDokumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
