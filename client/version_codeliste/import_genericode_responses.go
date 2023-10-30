// Code generated by go-swagger; DO NOT EDIT.

package version_codeliste

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ImportGenericodeReader is a Reader for the ImportGenericode structure.
type ImportGenericodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportGenericodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewImportGenericodeDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewImportGenericodeDefault creates a ImportGenericodeDefault with default headers values
func NewImportGenericodeDefault(code int) *ImportGenericodeDefault {
	return &ImportGenericodeDefault{
		_statusCode: code,
	}
}

/*
ImportGenericodeDefault describes a response with status code -1, with default header values.

successful operation
*/
type ImportGenericodeDefault struct {
	_statusCode int
}

// IsSuccess returns true when this import genericode default response has a 2xx status code
func (o *ImportGenericodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this import genericode default response has a 3xx status code
func (o *ImportGenericodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this import genericode default response has a 4xx status code
func (o *ImportGenericodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this import genericode default response has a 5xx status code
func (o *ImportGenericodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this import genericode default response a status code equal to that given
func (o *ImportGenericodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the import genericode default response
func (o *ImportGenericodeDefault) Code() int {
	return o._statusCode
}

func (o *ImportGenericodeDefault) Error() string {
	return fmt.Sprintf("[POST /version_codeliste/{kennung}/genericode/load][%d] importGenericode default ", o._statusCode)
}

func (o *ImportGenericodeDefault) String() string {
	return fmt.Sprintf("[POST /version_codeliste/{kennung}/genericode/load][%d] importGenericode default ", o._statusCode)
}

func (o *ImportGenericodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}