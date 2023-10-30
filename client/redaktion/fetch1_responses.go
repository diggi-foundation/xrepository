// Code generated by go-swagger; DO NOT EDIT.

package redaktion

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Fetch1Reader is a Reader for the Fetch1 structure.
type Fetch1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Fetch1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewFetch1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewFetch1Default creates a Fetch1Default with default headers values
func NewFetch1Default(code int) *Fetch1Default {
	return &Fetch1Default{
		_statusCode: code,
	}
}

/*
Fetch1Default describes a response with status code -1, with default header values.

successful operation
*/
type Fetch1Default struct {
	_statusCode int
}

// IsSuccess returns true when this fetch 1 default response has a 2xx status code
func (o *Fetch1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fetch 1 default response has a 3xx status code
func (o *Fetch1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fetch 1 default response has a 4xx status code
func (o *Fetch1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fetch 1 default response has a 5xx status code
func (o *Fetch1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fetch 1 default response a status code equal to that given
func (o *Fetch1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fetch 1 default response
func (o *Fetch1Default) Code() int {
	return o._statusCode
}

func (o *Fetch1Default) Error() string {
	return fmt.Sprintf("[GET /redaktionelle_liste/{urn}][%d] fetch_1 default ", o._statusCode)
}

func (o *Fetch1Default) String() string {
	return fmt.Sprintf("[GET /redaktionelle_liste/{urn}][%d] fetch_1 default ", o._statusCode)
}

func (o *Fetch1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}