// Code generated by go-swagger; DO NOT EDIT.

package redaktion

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new redaktion API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for redaktion API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Create(params *CreateParams, opts ...ClientOption) error

	Delete(params *DeleteParams, opts ...ClientOption) error

	Delete1(params *Delete1Params, opts ...ClientOption) error

	Fetch(params *FetchParams, opts ...ClientOption) error

	FetchAll(params *FetchAllParams, opts ...ClientOption) error

	FetchAll1(params *FetchAll1Params, opts ...ClientOption) error

	FetchAll2(params *FetchAll2Params, opts ...ClientOption) error

	FetchAll3(params *FetchAll3Params, opts ...ClientOption) error

	FetchByPath(params *FetchByPathParams, opts ...ClientOption) error

	Fetch1(params *Fetch1Params, opts ...ClientOption) error

	Fetch2(params *Fetch2Params, opts ...ClientOption) error

	Fetch3(params *Fetch3Params, opts ...ClientOption) error

	Update3(params *Update3Params, opts ...ClientOption) error

	Update4(params *Update4Params, opts ...ClientOption) error

	Update5(params *Update5Params, opts ...ClientOption) error

	Update6(params *Update6Params, opts ...ClientOption) error

	Upload(params *UploadParams, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
Create anlegens von bildern
*/
func (a *Client) Create(params *CreateParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create",
		Method:             "POST",
		PathPattern:        "/bild",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Delete delete API
*/
func (a *Client) Delete(params *DeleteParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete",
		Method:             "DELETE",
		PathPattern:        "/bild/{urn}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Delete1 delete 1 API
*/
func (a *Client) Delete1(params *Delete1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDelete1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_1",
		Method:             "DELETE",
		PathPattern:        "/seite/{urn}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Delete1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Fetch abrufs des bildes oder der meta daten je nach angefragtem mime type via accept header default ist die bereitstellung des eigentlichen bildes
*/
func (a *Client) Fetch(params *FetchParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetch",
		Method:             "GET",
		PathPattern:        "/bild/{urn}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FetchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
FetchAll lieferts eine liste alle inhalte dieses typs
*/
func (a *Client) FetchAll(params *FetchAllParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchAllParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchAll",
		Method:             "GET",
		PathPattern:        "/bild",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FetchAllReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
FetchAll1 lieferts eine liste alle inhalte dieses typs
*/
func (a *Client) FetchAll1(params *FetchAll1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchAll1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchAll_1",
		Method:             "GET",
		PathPattern:        "/redaktionelle_liste",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FetchAll1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
FetchAll2 lieferts eine liste alle inhalte dieses typs
*/
func (a *Client) FetchAll2(params *FetchAll2Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchAll2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchAll_2",
		Method:             "GET",
		PathPattern:        "/template",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FetchAll2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
FetchAll3 lieferts eine liste alle inhalte dieses typs
*/
func (a *Client) FetchAll3(params *FetchAll3Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchAll3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchAll_3",
		Method:             "GET",
		PathPattern:        "/seite",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FetchAll3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
FetchByPath fetch by path API
*/
func (a *Client) FetchByPath(params *FetchByPathParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchByPathParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchByPath",
		Method:             "GET",
		PathPattern:        "/seite/pfad/{pfad}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FetchByPathReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Fetch1 lädts das modell als JSON oder XML
*/
func (a *Client) Fetch1(params *Fetch1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetch1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetch_1",
		Method:             "GET",
		PathPattern:        "/redaktionelle_liste/{urn}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Fetch1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Fetch2 lädts das modell als JSON oder XML
*/
func (a *Client) Fetch2(params *Fetch2Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetch2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetch_2",
		Method:             "GET",
		PathPattern:        "/template/{urn}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Fetch2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Fetch3 lädts das modell als JSON oder XML
*/
func (a *Client) Fetch3(params *Fetch3Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetch3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetch_3",
		Method:             "GET",
		PathPattern:        "/seite/{urn}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Fetch3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Update3 aktualisierungs oder anlegen von bildern
*/
func (a *Client) Update3(params *Update3Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdate3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_3",
		Method:             "POST",
		PathPattern:        "/bild/{urn}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Update3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Update4 aktualisierungs des objekts muss xml sein
*/
func (a *Client) Update4(params *Update4Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdate4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_4",
		Method:             "POST",
		PathPattern:        "/redaktionelle_liste/{urn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Update4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Update5 aktualisierungs des objekts muss xml sein
*/
func (a *Client) Update5(params *Update5Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdate5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_5",
		Method:             "POST",
		PathPattern:        "/template/{urn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Update5Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Update6 aktualisierungs des objekts muss xml sein
*/
func (a *Client) Update6(params *Update6Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdate6Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_6",
		Method:             "POST",
		PathPattern:        "/seite/{urn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Update6Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Upload upload API
*/
func (a *Client) Upload(params *UploadParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "upload",
		Method:             "POST",
		PathPattern:        "/seite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UploadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
