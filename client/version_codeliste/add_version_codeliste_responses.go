// Code generated by go-swagger; DO NOT EDIT.

package version_codeliste

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddVersionCodelisteReader is a Reader for the AddVersionCodeliste structure.
type AddVersionCodelisteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddVersionCodelisteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewAddVersionCodelisteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewAddVersionCodelisteDefault creates a AddVersionCodelisteDefault with default headers values
func NewAddVersionCodelisteDefault(code int) *AddVersionCodelisteDefault {
	return &AddVersionCodelisteDefault{
		_statusCode: code,
	}
}

/*
AddVersionCodelisteDefault describes a response with status code -1, with default header values.

successful operation
*/
type AddVersionCodelisteDefault struct {
	_statusCode int
}

// IsSuccess returns true when this add version codeliste default response has a 2xx status code
func (o *AddVersionCodelisteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add version codeliste default response has a 3xx status code
func (o *AddVersionCodelisteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add version codeliste default response has a 4xx status code
func (o *AddVersionCodelisteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add version codeliste default response has a 5xx status code
func (o *AddVersionCodelisteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add version codeliste default response a status code equal to that given
func (o *AddVersionCodelisteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add version codeliste default response
func (o *AddVersionCodelisteDefault) Code() int {
	return o._statusCode
}

func (o *AddVersionCodelisteDefault) Error() string {
	return fmt.Sprintf("[POST /version_codeliste/hochladen][%d] addVersionCodeliste default ", o._statusCode)
}

func (o *AddVersionCodelisteDefault) String() string {
	return fmt.Sprintf("[POST /version_codeliste/hochladen][%d] addVersionCodeliste default ", o._statusCode)
}

func (o *AddVersionCodelisteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
