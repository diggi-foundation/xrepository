// Code generated by go-swagger; DO NOT EDIT.

package version_standard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveProprietaeresFachmodellReader is a Reader for the RemoveProprietaeresFachmodell structure.
type RemoveProprietaeresFachmodellReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveProprietaeresFachmodellReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewRemoveProprietaeresFachmodellDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewRemoveProprietaeresFachmodellDefault creates a RemoveProprietaeresFachmodellDefault with default headers values
func NewRemoveProprietaeresFachmodellDefault(code int) *RemoveProprietaeresFachmodellDefault {
	return &RemoveProprietaeresFachmodellDefault{
		_statusCode: code,
	}
}

/*
RemoveProprietaeresFachmodellDefault describes a response with status code -1, with default header values.

successful operation
*/
type RemoveProprietaeresFachmodellDefault struct {
	_statusCode int
}

// IsSuccess returns true when this remove proprietaeres fachmodell default response has a 2xx status code
func (o *RemoveProprietaeresFachmodellDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this remove proprietaeres fachmodell default response has a 3xx status code
func (o *RemoveProprietaeresFachmodellDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this remove proprietaeres fachmodell default response has a 4xx status code
func (o *RemoveProprietaeresFachmodellDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this remove proprietaeres fachmodell default response has a 5xx status code
func (o *RemoveProprietaeresFachmodellDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this remove proprietaeres fachmodell default response a status code equal to that given
func (o *RemoveProprietaeresFachmodellDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the remove proprietaeres fachmodell default response
func (o *RemoveProprietaeresFachmodellDefault) Code() int {
	return o._statusCode
}

func (o *RemoveProprietaeresFachmodellDefault) Error() string {
	return fmt.Sprintf("[DELETE /version_standard/{urn}/originalFachmodellProprietaer][%d] removeProprietaeresFachmodell default ", o._statusCode)
}

func (o *RemoveProprietaeresFachmodellDefault) String() string {
	return fmt.Sprintf("[DELETE /version_standard/{urn}/originalFachmodellProprietaer][%d] removeProprietaeresFachmodell default ", o._statusCode)
}

func (o *RemoveProprietaeresFachmodellDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
