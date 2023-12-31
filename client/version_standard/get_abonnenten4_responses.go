// Code generated by go-swagger; DO NOT EDIT.

package version_standard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAbonnenten4Reader is a Reader for the GetAbonnenten4 structure.
type GetAbonnenten4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAbonnenten4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetAbonnenten4Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetAbonnenten4Default creates a GetAbonnenten4Default with default headers values
func NewGetAbonnenten4Default(code int) *GetAbonnenten4Default {
	return &GetAbonnenten4Default{
		_statusCode: code,
	}
}

/*
GetAbonnenten4Default describes a response with status code -1, with default header values.

successful operation
*/
type GetAbonnenten4Default struct {
	_statusCode int
}

// IsSuccess returns true when this get abonnenten 4 default response has a 2xx status code
func (o *GetAbonnenten4Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get abonnenten 4 default response has a 3xx status code
func (o *GetAbonnenten4Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get abonnenten 4 default response has a 4xx status code
func (o *GetAbonnenten4Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get abonnenten 4 default response has a 5xx status code
func (o *GetAbonnenten4Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get abonnenten 4 default response a status code equal to that given
func (o *GetAbonnenten4Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get abonnenten 4 default response
func (o *GetAbonnenten4Default) Code() int {
	return o._statusCode
}

func (o *GetAbonnenten4Default) Error() string {
	return fmt.Sprintf("[GET /version_standard/{urn}/abonnenten][%d] getAbonnenten_4 default ", o._statusCode)
}

func (o *GetAbonnenten4Default) String() string {
	return fmt.Sprintf("[GET /version_standard/{urn}/abonnenten][%d] getAbonnenten_4 default ", o._statusCode)
}

func (o *GetAbonnenten4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
