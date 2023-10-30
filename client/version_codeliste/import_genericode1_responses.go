// Code generated by go-swagger; DO NOT EDIT.

package version_codeliste

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ImportGenericode1Reader is a Reader for the ImportGenericode1 structure.
type ImportGenericode1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportGenericode1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewImportGenericode1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewImportGenericode1Default creates a ImportGenericode1Default with default headers values
func NewImportGenericode1Default(code int) *ImportGenericode1Default {
	return &ImportGenericode1Default{
		_statusCode: code,
	}
}

/*
ImportGenericode1Default describes a response with status code -1, with default header values.

successful operation
*/
type ImportGenericode1Default struct {
	_statusCode int
}

// IsSuccess returns true when this import genericode 1 default response has a 2xx status code
func (o *ImportGenericode1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this import genericode 1 default response has a 3xx status code
func (o *ImportGenericode1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this import genericode 1 default response has a 4xx status code
func (o *ImportGenericode1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this import genericode 1 default response has a 5xx status code
func (o *ImportGenericode1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this import genericode 1 default response a status code equal to that given
func (o *ImportGenericode1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the import genericode 1 default response
func (o *ImportGenericode1Default) Code() int {
	return o._statusCode
}

func (o *ImportGenericode1Default) Error() string {
	return fmt.Sprintf("[POST /version_codeliste/genericode/load][%d] importGenericode_1 default ", o._statusCode)
}

func (o *ImportGenericode1Default) String() string {
	return fmt.Sprintf("[POST /version_codeliste/genericode/load][%d] importGenericode_1 default ", o._statusCode)
}

func (o *ImportGenericode1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
