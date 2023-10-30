// Code generated by go-swagger; DO NOT EDIT.

package interop_matrix

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetSingleInteropDatenReader is a Reader for the GetSingleInteropDaten structure.
type GetSingleInteropDatenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSingleInteropDatenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetSingleInteropDatenDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetSingleInteropDatenDefault creates a GetSingleInteropDatenDefault with default headers values
func NewGetSingleInteropDatenDefault(code int) *GetSingleInteropDatenDefault {
	return &GetSingleInteropDatenDefault{
		_statusCode: code,
	}
}

/*
GetSingleInteropDatenDefault describes a response with status code -1, with default header values.

successful operation
*/
type GetSingleInteropDatenDefault struct {
	_statusCode int
}

// IsSuccess returns true when this get single interop daten default response has a 2xx status code
func (o *GetSingleInteropDatenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get single interop daten default response has a 3xx status code
func (o *GetSingleInteropDatenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get single interop daten default response has a 4xx status code
func (o *GetSingleInteropDatenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get single interop daten default response has a 5xx status code
func (o *GetSingleInteropDatenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get single interop daten default response a status code equal to that given
func (o *GetSingleInteropDatenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get single interop daten default response
func (o *GetSingleInteropDatenDefault) Code() int {
	return o._statusCode
}

func (o *GetSingleInteropDatenDefault) Error() string {
	return fmt.Sprintf("[GET /interop/daten-{kennung}][%d] getSingleInteropDaten default ", o._statusCode)
}

func (o *GetSingleInteropDatenDefault) String() string {
	return fmt.Sprintf("[GET /interop/daten-{kennung}][%d] getSingleInteropDaten default ", o._statusCode)
}

func (o *GetSingleInteropDatenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}