// Code generated by go-swagger; DO NOT EDIT.

package codeliste

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetTechnischerBestandteilReader is a Reader for the GetTechnischerBestandteil structure.
type GetTechnischerBestandteilReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTechnischerBestandteilReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetTechnischerBestandteilDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetTechnischerBestandteilDefault creates a GetTechnischerBestandteilDefault with default headers values
func NewGetTechnischerBestandteilDefault(code int) *GetTechnischerBestandteilDefault {
	return &GetTechnischerBestandteilDefault{
		_statusCode: code,
	}
}

/*
GetTechnischerBestandteilDefault describes a response with status code -1, with default header values.

successful operation
*/
type GetTechnischerBestandteilDefault struct {
	_statusCode int
}

// IsSuccess returns true when this get technischer bestandteil default response has a 2xx status code
func (o *GetTechnischerBestandteilDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get technischer bestandteil default response has a 3xx status code
func (o *GetTechnischerBestandteilDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get technischer bestandteil default response has a 4xx status code
func (o *GetTechnischerBestandteilDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get technischer bestandteil default response has a 5xx status code
func (o *GetTechnischerBestandteilDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get technischer bestandteil default response a status code equal to that given
func (o *GetTechnischerBestandteilDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get technischer bestandteil default response
func (o *GetTechnischerBestandteilDefault) Code() int {
	return o._statusCode
}

func (o *GetTechnischerBestandteilDefault) Error() string {
	return fmt.Sprintf("[GET /codeliste/{kennung}/technischerBestandteil/{name}][%d] getTechnischerBestandteil default ", o._statusCode)
}

func (o *GetTechnischerBestandteilDefault) String() string {
	return fmt.Sprintf("[GET /codeliste/{kennung}/technischerBestandteil/{name}][%d] getTechnischerBestandteil default ", o._statusCode)
}

func (o *GetTechnischerBestandteilDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
