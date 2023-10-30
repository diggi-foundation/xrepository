// Code generated by go-swagger; DO NOT EDIT.

package version_standard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveDokumentFromVersionStandardReader is a Reader for the RemoveDokumentFromVersionStandard structure.
type RemoveDokumentFromVersionStandardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveDokumentFromVersionStandardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewRemoveDokumentFromVersionStandardDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewRemoveDokumentFromVersionStandardDefault creates a RemoveDokumentFromVersionStandardDefault with default headers values
func NewRemoveDokumentFromVersionStandardDefault(code int) *RemoveDokumentFromVersionStandardDefault {
	return &RemoveDokumentFromVersionStandardDefault{
		_statusCode: code,
	}
}

/*
RemoveDokumentFromVersionStandardDefault describes a response with status code -1, with default header values.

successful operation
*/
type RemoveDokumentFromVersionStandardDefault struct {
	_statusCode int
}

// IsSuccess returns true when this remove dokument from version standard default response has a 2xx status code
func (o *RemoveDokumentFromVersionStandardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this remove dokument from version standard default response has a 3xx status code
func (o *RemoveDokumentFromVersionStandardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this remove dokument from version standard default response has a 4xx status code
func (o *RemoveDokumentFromVersionStandardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this remove dokument from version standard default response has a 5xx status code
func (o *RemoveDokumentFromVersionStandardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this remove dokument from version standard default response a status code equal to that given
func (o *RemoveDokumentFromVersionStandardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the remove dokument from version standard default response
func (o *RemoveDokumentFromVersionStandardDefault) Code() int {
	return o._statusCode
}

func (o *RemoveDokumentFromVersionStandardDefault) Error() string {
	return fmt.Sprintf("[DELETE /version_standard/{kennung}/dokument/{name}][%d] removeDokumentFromVersionStandard default ", o._statusCode)
}

func (o *RemoveDokumentFromVersionStandardDefault) String() string {
	return fmt.Sprintf("[DELETE /version_standard/{kennung}/dokument/{name}][%d] removeDokumentFromVersionStandard default ", o._statusCode)
}

func (o *RemoveDokumentFromVersionStandardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
