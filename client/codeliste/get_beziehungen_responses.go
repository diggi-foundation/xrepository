// Code generated by go-swagger; DO NOT EDIT.

package codeliste

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/diggi.foundation/xrepository/models"
)

// GetBeziehungenReader is a Reader for the GetBeziehungen structure.
type GetBeziehungenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBeziehungenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBeziehungenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /codeliste/{urn}/beziehungen] getBeziehungen", response, response.Code())
	}
}

// NewGetBeziehungenOK creates a GetBeziehungenOK with default headers values
func NewGetBeziehungenOK() *GetBeziehungenOK {
	return &GetBeziehungenOK{}
}

/*
GetBeziehungenOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBeziehungenOK struct {
	Payload *models.DebugInfo
}

// IsSuccess returns true when this get beziehungen o k response has a 2xx status code
func (o *GetBeziehungenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get beziehungen o k response has a 3xx status code
func (o *GetBeziehungenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get beziehungen o k response has a 4xx status code
func (o *GetBeziehungenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get beziehungen o k response has a 5xx status code
func (o *GetBeziehungenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get beziehungen o k response a status code equal to that given
func (o *GetBeziehungenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get beziehungen o k response
func (o *GetBeziehungenOK) Code() int {
	return 200
}

func (o *GetBeziehungenOK) Error() string {
	return fmt.Sprintf("[GET /codeliste/{urn}/beziehungen][%d] getBeziehungenOK  %+v", 200, o.Payload)
}

func (o *GetBeziehungenOK) String() string {
	return fmt.Sprintf("[GET /codeliste/{urn}/beziehungen][%d] getBeziehungenOK  %+v", 200, o.Payload)
}

func (o *GetBeziehungenOK) GetPayload() *models.DebugInfo {
	return o.Payload
}

func (o *GetBeziehungenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebugInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}