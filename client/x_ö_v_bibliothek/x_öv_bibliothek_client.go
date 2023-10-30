// Code generated by go-swagger; DO NOT EDIT.

package x_ö_v_bibliothek

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new x ö v bibliothek API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for x ö v bibliothek API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddVersionXoevBibliothek(params *AddVersionXoevBibliothekParams, opts ...ClientOption) error

	GetAll(params *GetAllParams, opts ...ClientOption) error

	XoevFreigeben(params *XoevFreigebenParams, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
AddVersionXoevBibliothek add version xoev bibliothek API
*/
func (a *Client) AddVersionXoevBibliothek(params *AddVersionXoevBibliothekParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddVersionXoevBibliothekParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addVersionXoevBibliothek",
		Method:             "POST",
		PathPattern:        "/xoev_bibliothek/hochladen",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddVersionXoevBibliothekReader{formats: a.formats},
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
GetAll get all API
*/
func (a *Client) GetAll(params *GetAllParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAll",
		Method:             "GET",
		PathPattern:        "/xoev_bibliothek",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAllReader{formats: a.formats},
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
XoevFreigeben xoev freigeben API
*/
func (a *Client) XoevFreigeben(params *XoevFreigebenParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewXoevFreigebenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "xoevFreigeben",
		Method:             "POST",
		PathPattern:        "/xoev_bibliothek/freigeben",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &XoevFreigebenReader{formats: a.formats},
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
