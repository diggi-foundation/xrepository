// Code generated by go-swagger; DO NOT EDIT.

package codeliste

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UploadDokumentReader is a Reader for the UploadDokument structure.
type UploadDokumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadDokumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewUploadDokumentDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewUploadDokumentDefault creates a UploadDokumentDefault with default headers values
func NewUploadDokumentDefault(code int) *UploadDokumentDefault {
	return &UploadDokumentDefault{
		_statusCode: code,
	}
}

/*
UploadDokumentDefault describes a response with status code -1, with default header values.

successful operation
*/
type UploadDokumentDefault struct {
	_statusCode int
}

// IsSuccess returns true when this upload dokument default response has a 2xx status code
func (o *UploadDokumentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this upload dokument default response has a 3xx status code
func (o *UploadDokumentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this upload dokument default response has a 4xx status code
func (o *UploadDokumentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this upload dokument default response has a 5xx status code
func (o *UploadDokumentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this upload dokument default response a status code equal to that given
func (o *UploadDokumentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the upload dokument default response
func (o *UploadDokumentDefault) Code() int {
	return o._statusCode
}

func (o *UploadDokumentDefault) Error() string {
	return fmt.Sprintf("[POST /codeliste/{urn}/add-document][%d] uploadDokument default ", o._statusCode)
}

func (o *UploadDokumentDefault) String() string {
	return fmt.Sprintf("[POST /codeliste/{urn}/add-document][%d] uploadDokument default ", o._statusCode)
}

func (o *UploadDokumentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
