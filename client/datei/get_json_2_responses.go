// Code generated by go-swagger; DO NOT EDIT.

package datei

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetJSON2Reader is a Reader for the GetJSON2 structure.
type GetJSON2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJSON2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetJSON2Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetJSON2Default creates a GetJSON2Default with default headers values
func NewGetJSON2Default(code int) *GetJSON2Default {
	return &GetJSON2Default{
		_statusCode: code,
	}
}

/*
GetJSON2Default describes a response with status code -1, with default header values.

successful operation
*/
type GetJSON2Default struct {
	_statusCode int
}

// IsSuccess returns true when this get Json 2 default response has a 2xx status code
func (o *GetJSON2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get Json 2 default response has a 3xx status code
func (o *GetJSON2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get Json 2 default response has a 4xx status code
func (o *GetJSON2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get Json 2 default response has a 5xx status code
func (o *GetJSON2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get Json 2 default response a status code equal to that given
func (o *GetJSON2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get Json 2 default response
func (o *GetJSON2Default) Code() int {
	return o._statusCode
}

func (o *GetJSON2Default) Error() string {
	return fmt.Sprintf("[GET /datei/{urn}/model][%d] getJson_2 default ", o._statusCode)
}

func (o *GetJSON2Default) String() string {
	return fmt.Sprintf("[GET /datei/{urn}/model][%d] getJson_2 default ", o._statusCode)
}

func (o *GetJSON2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
