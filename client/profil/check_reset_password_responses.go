// Code generated by go-swagger; DO NOT EDIT.

package profil

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CheckResetPasswordReader is a Reader for the CheckResetPassword structure.
type CheckResetPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckResetPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewCheckResetPasswordDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewCheckResetPasswordDefault creates a CheckResetPasswordDefault with default headers values
func NewCheckResetPasswordDefault(code int) *CheckResetPasswordDefault {
	return &CheckResetPasswordDefault{
		_statusCode: code,
	}
}

/*
CheckResetPasswordDefault describes a response with status code -1, with default header values.

successful operation
*/
type CheckResetPasswordDefault struct {
	_statusCode int
}

// IsSuccess returns true when this check reset password default response has a 2xx status code
func (o *CheckResetPasswordDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this check reset password default response has a 3xx status code
func (o *CheckResetPasswordDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this check reset password default response has a 4xx status code
func (o *CheckResetPasswordDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this check reset password default response has a 5xx status code
func (o *CheckResetPasswordDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this check reset password default response a status code equal to that given
func (o *CheckResetPasswordDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the check reset password default response
func (o *CheckResetPasswordDefault) Code() int {
	return o._statusCode
}

func (o *CheckResetPasswordDefault) Error() string {
	return fmt.Sprintf("[GET /profil/passwort_zuruecksetzen/{resetKey}][%d] checkResetPassword default ", o._statusCode)
}

func (o *CheckResetPasswordDefault) String() string {
	return fmt.Sprintf("[GET /profil/passwort_zuruecksetzen/{resetKey}][%d] checkResetPassword default ", o._statusCode)
}

func (o *CheckResetPasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
