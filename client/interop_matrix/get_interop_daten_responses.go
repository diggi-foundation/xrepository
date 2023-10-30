// Code generated by go-swagger; DO NOT EDIT.

package interop_matrix

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetInteropDatenReader is a Reader for the GetInteropDaten structure.
type GetInteropDatenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInteropDatenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetInteropDatenDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetInteropDatenDefault creates a GetInteropDatenDefault with default headers values
func NewGetInteropDatenDefault(code int) *GetInteropDatenDefault {
	return &GetInteropDatenDefault{
		_statusCode: code,
	}
}

/*
GetInteropDatenDefault describes a response with status code -1, with default header values.

successful operation
*/
type GetInteropDatenDefault struct {
	_statusCode int
}

// IsSuccess returns true when this get interop daten default response has a 2xx status code
func (o *GetInteropDatenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get interop daten default response has a 3xx status code
func (o *GetInteropDatenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get interop daten default response has a 4xx status code
func (o *GetInteropDatenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get interop daten default response has a 5xx status code
func (o *GetInteropDatenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get interop daten default response a status code equal to that given
func (o *GetInteropDatenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get interop daten default response
func (o *GetInteropDatenDefault) Code() int {
	return o._statusCode
}

func (o *GetInteropDatenDefault) Error() string {
	return fmt.Sprintf("[GET /interop/matrix][%d] getInteropDaten default ", o._statusCode)
}

func (o *GetInteropDatenDefault) String() string {
	return fmt.Sprintf("[GET /interop/matrix][%d] getInteropDaten default ", o._statusCode)
}

func (o *GetInteropDatenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
