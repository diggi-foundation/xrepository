// Code generated by go-swagger; DO NOT EDIT.

package nutzer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LoginReader is a Reader for the Login structure.
type LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewLoginDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewLoginDefault creates a LoginDefault with default headers values
func NewLoginDefault(code int) *LoginDefault {
	return &LoginDefault{
		_statusCode: code,
	}
}

/*
LoginDefault describes a response with status code -1, with default header values.

successful operation
*/
type LoginDefault struct {
	_statusCode int
}

// IsSuccess returns true when this login default response has a 2xx status code
func (o *LoginDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this login default response has a 3xx status code
func (o *LoginDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this login default response has a 4xx status code
func (o *LoginDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this login default response has a 5xx status code
func (o *LoginDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this login default response a status code equal to that given
func (o *LoginDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the login default response
func (o *LoginDefault) Code() int {
	return o._statusCode
}

func (o *LoginDefault) Error() string {
	return fmt.Sprintf("[POST /session/login][%d] login default ", o._statusCode)
}

func (o *LoginDefault) String() string {
	return fmt.Sprintf("[POST /session/login][%d] login default ", o._statusCode)
}

func (o *LoginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
