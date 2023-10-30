// Code generated by go-swagger; DO NOT EDIT.

package version_codeliste

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDokumentReader is a Reader for the GetDokument structure.
type GetDokumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDokumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetDokumentDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetDokumentDefault creates a GetDokumentDefault with default headers values
func NewGetDokumentDefault(code int) *GetDokumentDefault {
	return &GetDokumentDefault{
		_statusCode: code,
	}
}

/*
GetDokumentDefault describes a response with status code -1, with default header values.

successful operation
*/
type GetDokumentDefault struct {
	_statusCode int
}

// IsSuccess returns true when this get dokument default response has a 2xx status code
func (o *GetDokumentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get dokument default response has a 3xx status code
func (o *GetDokumentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get dokument default response has a 4xx status code
func (o *GetDokumentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get dokument default response has a 5xx status code
func (o *GetDokumentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get dokument default response a status code equal to that given
func (o *GetDokumentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get dokument default response
func (o *GetDokumentDefault) Code() int {
	return o._statusCode
}

func (o *GetDokumentDefault) Error() string {
	return fmt.Sprintf("[GET /version_codeliste/{kennung}/dokument/{name}][%d] getDokument default ", o._statusCode)
}

func (o *GetDokumentDefault) String() string {
	return fmt.Sprintf("[GET /version_codeliste/{kennung}/dokument/{name}][%d] getDokument default ", o._statusCode)
}

func (o *GetDokumentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}