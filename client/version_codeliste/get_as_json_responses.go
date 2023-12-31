// Code generated by go-swagger; DO NOT EDIT.

package version_codeliste

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAsJSONReader is a Reader for the GetAsJSON structure.
type GetAsJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAsJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetAsJSONDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetAsJSONDefault creates a GetAsJSONDefault with default headers values
func NewGetAsJSONDefault(code int) *GetAsJSONDefault {
	return &GetAsJSONDefault{
		_statusCode: code,
	}
}

/*
GetAsJSONDefault describes a response with status code -1, with default header values.

successful operation
*/
type GetAsJSONDefault struct {
	_statusCode int
}

// IsSuccess returns true when this get as Json default response has a 2xx status code
func (o *GetAsJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get as Json default response has a 3xx status code
func (o *GetAsJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get as Json default response has a 4xx status code
func (o *GetAsJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get as Json default response has a 5xx status code
func (o *GetAsJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get as Json default response a status code equal to that given
func (o *GetAsJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get as Json default response
func (o *GetAsJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetAsJSONDefault) Error() string {
	return fmt.Sprintf("[GET /version_codeliste/{kennung}/download/{name}.json][%d] getAsJson default ", o._statusCode)
}

func (o *GetAsJSONDefault) String() string {
	return fmt.Sprintf("[GET /version_codeliste/{kennung}/download/{name}.json][%d] getAsJson default ", o._statusCode)
}

func (o *GetAsJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
