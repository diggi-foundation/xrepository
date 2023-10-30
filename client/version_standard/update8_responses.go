// Code generated by go-swagger; DO NOT EDIT.

package version_standard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Update8Reader is a Reader for the Update8 structure.
type Update8Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Update8Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewUpdate8Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewUpdate8Default creates a Update8Default with default headers values
func NewUpdate8Default(code int) *Update8Default {
	return &Update8Default{
		_statusCode: code,
	}
}

/*
Update8Default describes a response with status code -1, with default header values.

successful operation
*/
type Update8Default struct {
	_statusCode int
}

// IsSuccess returns true when this update 8 default response has a 2xx status code
func (o *Update8Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update 8 default response has a 3xx status code
func (o *Update8Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update 8 default response has a 4xx status code
func (o *Update8Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update 8 default response has a 5xx status code
func (o *Update8Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update 8 default response a status code equal to that given
func (o *Update8Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update 8 default response
func (o *Update8Default) Code() int {
	return o._statusCode
}

func (o *Update8Default) Error() string {
	return fmt.Sprintf("[POST /version_standard/{urn}][%d] update_8 default ", o._statusCode)
}

func (o *Update8Default) String() string {
	return fmt.Sprintf("[POST /version_standard/{urn}][%d] update_8 default ", o._statusCode)
}

func (o *Update8Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
