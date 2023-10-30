// Code generated by go-swagger; DO NOT EDIT.

package redaktion

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Update3Reader is a Reader for the Update3 structure.
type Update3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Update3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewUpdate3Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewUpdate3Default creates a Update3Default with default headers values
func NewUpdate3Default(code int) *Update3Default {
	return &Update3Default{
		_statusCode: code,
	}
}

/*
Update3Default describes a response with status code -1, with default header values.

successful operation
*/
type Update3Default struct {
	_statusCode int
}

// IsSuccess returns true when this update 3 default response has a 2xx status code
func (o *Update3Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update 3 default response has a 3xx status code
func (o *Update3Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update 3 default response has a 4xx status code
func (o *Update3Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update 3 default response has a 5xx status code
func (o *Update3Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update 3 default response a status code equal to that given
func (o *Update3Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update 3 default response
func (o *Update3Default) Code() int {
	return o._statusCode
}

func (o *Update3Default) Error() string {
	return fmt.Sprintf("[POST /bild/{urn}][%d] update_3 default ", o._statusCode)
}

func (o *Update3Default) String() string {
	return fmt.Sprintf("[POST /bild/{urn}][%d] update_3 default ", o._statusCode)
}

func (o *Update3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
