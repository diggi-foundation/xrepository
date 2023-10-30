// Code generated by go-swagger; DO NOT EDIT.

package redaktion

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Delete1Reader is a Reader for the Delete1 structure.
type Delete1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Delete1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDelete1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDelete1Default creates a Delete1Default with default headers values
func NewDelete1Default(code int) *Delete1Default {
	return &Delete1Default{
		_statusCode: code,
	}
}

/*
Delete1Default describes a response with status code -1, with default header values.

successful operation
*/
type Delete1Default struct {
	_statusCode int
}

// IsSuccess returns true when this delete 1 default response has a 2xx status code
func (o *Delete1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete 1 default response has a 3xx status code
func (o *Delete1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete 1 default response has a 4xx status code
func (o *Delete1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete 1 default response has a 5xx status code
func (o *Delete1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete 1 default response a status code equal to that given
func (o *Delete1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete 1 default response
func (o *Delete1Default) Code() int {
	return o._statusCode
}

func (o *Delete1Default) Error() string {
	return fmt.Sprintf("[DELETE /seite/{urn}][%d] delete_1 default ", o._statusCode)
}

func (o *Delete1Default) String() string {
	return fmt.Sprintf("[DELETE /seite/{urn}][%d] delete_1 default ", o._statusCode)
}

func (o *Delete1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
