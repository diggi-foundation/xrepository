// Code generated by go-swagger; DO NOT EDIT.

package version_standard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DownloadGenutztenCodelistenReader is a Reader for the DownloadGenutztenCodelisten structure.
type DownloadGenutztenCodelistenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadGenutztenCodelistenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDownloadGenutztenCodelistenDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDownloadGenutztenCodelistenDefault creates a DownloadGenutztenCodelistenDefault with default headers values
func NewDownloadGenutztenCodelistenDefault(code int) *DownloadGenutztenCodelistenDefault {
	return &DownloadGenutztenCodelistenDefault{
		_statusCode: code,
	}
}

/*
DownloadGenutztenCodelistenDefault describes a response with status code -1, with default header values.

successful operation
*/
type DownloadGenutztenCodelistenDefault struct {
	_statusCode int
}

// IsSuccess returns true when this download genutzten codelisten default response has a 2xx status code
func (o *DownloadGenutztenCodelistenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this download genutzten codelisten default response has a 3xx status code
func (o *DownloadGenutztenCodelistenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this download genutzten codelisten default response has a 4xx status code
func (o *DownloadGenutztenCodelistenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this download genutzten codelisten default response has a 5xx status code
func (o *DownloadGenutztenCodelistenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this download genutzten codelisten default response a status code equal to that given
func (o *DownloadGenutztenCodelistenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the download genutzten codelisten default response
func (o *DownloadGenutztenCodelistenDefault) Code() int {
	return o._statusCode
}

func (o *DownloadGenutztenCodelistenDefault) Error() string {
	return fmt.Sprintf("[GET /version_standard/{urn}/genutzteAktuelleCodelisten][%d] downloadGenutztenCodelisten default ", o._statusCode)
}

func (o *DownloadGenutztenCodelistenDefault) String() string {
	return fmt.Sprintf("[GET /version_standard/{urn}/genutzteAktuelleCodelisten][%d] downloadGenutztenCodelisten default ", o._statusCode)
}

func (o *DownloadGenutztenCodelistenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
