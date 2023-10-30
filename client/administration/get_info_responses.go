// Code generated by go-swagger; DO NOT EDIT.

package administration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetInfoReader is a Reader for the GetInfo structure.
type GetInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetInfoDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetInfoDefault creates a GetInfoDefault with default headers values
func NewGetInfoDefault(code int) *GetInfoDefault {
	return &GetInfoDefault{
		_statusCode: code,
	}
}

/*
GetInfoDefault describes a response with status code -1, with default header values.

successful operation
*/
type GetInfoDefault struct {
	_statusCode int
}

// IsSuccess returns true when this get info default response has a 2xx status code
func (o *GetInfoDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get info default response has a 3xx status code
func (o *GetInfoDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get info default response has a 4xx status code
func (o *GetInfoDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get info default response has a 5xx status code
func (o *GetInfoDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get info default response a status code equal to that given
func (o *GetInfoDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get info default response
func (o *GetInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetInfoDefault) Error() string {
	return fmt.Sprintf("[GET /administration/extraktion][%d] getInfo default ", o._statusCode)
}

func (o *GetInfoDefault) String() string {
	return fmt.Sprintf("[GET /administration/extraktion][%d] getInfo default ", o._statusCode)
}

func (o *GetInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
