// Code generated by go-swagger; DO NOT EDIT.

package nutzer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/diggi-foundation/xrepository/models"
)

// GetLoginInfoReader is a Reader for the GetLoginInfo structure.
type GetLoginInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoginInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLoginInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /session/login] getLoginInfo", response, response.Code())
	}
}

// NewGetLoginInfoOK creates a GetLoginInfoOK with default headers values
func NewGetLoginInfoOK() *GetLoginInfoOK {
	return &GetLoginInfoOK{}
}

/*
GetLoginInfoOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLoginInfoOK struct {
	Payload *models.LoginData
}

// IsSuccess returns true when this get login info o k response has a 2xx status code
func (o *GetLoginInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get login info o k response has a 3xx status code
func (o *GetLoginInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get login info o k response has a 4xx status code
func (o *GetLoginInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get login info o k response has a 5xx status code
func (o *GetLoginInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get login info o k response a status code equal to that given
func (o *GetLoginInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get login info o k response
func (o *GetLoginInfoOK) Code() int {
	return 200
}

func (o *GetLoginInfoOK) Error() string {
	return fmt.Sprintf("[GET /session/login][%d] getLoginInfoOK  %+v", 200, o.Payload)
}

func (o *GetLoginInfoOK) String() string {
	return fmt.Sprintf("[GET /session/login][%d] getLoginInfoOK  %+v", 200, o.Payload)
}

func (o *GetLoginInfoOK) GetPayload() *models.LoginData {
	return o.Payload
}

func (o *GetLoginInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoginData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
