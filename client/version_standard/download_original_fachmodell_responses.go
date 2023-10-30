// Code generated by go-swagger; DO NOT EDIT.

package version_standard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DownloadOriginalFachmodellReader is a Reader for the DownloadOriginalFachmodell structure.
type DownloadOriginalFachmodellReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadOriginalFachmodellReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDownloadOriginalFachmodellDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDownloadOriginalFachmodellDefault creates a DownloadOriginalFachmodellDefault with default headers values
func NewDownloadOriginalFachmodellDefault(code int) *DownloadOriginalFachmodellDefault {
	return &DownloadOriginalFachmodellDefault{
		_statusCode: code,
	}
}

/*
DownloadOriginalFachmodellDefault describes a response with status code -1, with default header values.

successful operation
*/
type DownloadOriginalFachmodellDefault struct {
	_statusCode int
}

// IsSuccess returns true when this download original fachmodell default response has a 2xx status code
func (o *DownloadOriginalFachmodellDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this download original fachmodell default response has a 3xx status code
func (o *DownloadOriginalFachmodellDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this download original fachmodell default response has a 4xx status code
func (o *DownloadOriginalFachmodellDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this download original fachmodell default response has a 5xx status code
func (o *DownloadOriginalFachmodellDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this download original fachmodell default response a status code equal to that given
func (o *DownloadOriginalFachmodellDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the download original fachmodell default response
func (o *DownloadOriginalFachmodellDefault) Code() int {
	return o._statusCode
}

func (o *DownloadOriginalFachmodellDefault) Error() string {
	return fmt.Sprintf("[GET /version_standard/{urn}/originalFachmodellXMI][%d] downloadOriginalFachmodell default ", o._statusCode)
}

func (o *DownloadOriginalFachmodellDefault) String() string {
	return fmt.Sprintf("[GET /version_standard/{urn}/originalFachmodellXMI][%d] downloadOriginalFachmodell default ", o._statusCode)
}

func (o *DownloadOriginalFachmodellDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}