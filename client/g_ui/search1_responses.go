// Code generated by go-swagger; DO NOT EDIT.

package g_ui

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Search1Reader is a Reader for the Search1 structure.
type Search1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Search1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewSearch1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewSearch1Default creates a Search1Default with default headers values
func NewSearch1Default(code int) *Search1Default {
	return &Search1Default{
		_statusCode: code,
	}
}

/*
Search1Default describes a response with status code -1, with default header values.

successful operation
*/
type Search1Default struct {
	_statusCode int
}

// IsSuccess returns true when this search 1 default response has a 2xx status code
func (o *Search1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search 1 default response has a 3xx status code
func (o *Search1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search 1 default response has a 4xx status code
func (o *Search1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search 1 default response has a 5xx status code
func (o *Search1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search 1 default response a status code equal to that given
func (o *Search1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search 1 default response
func (o *Search1Default) Code() int {
	return o._statusCode
}

func (o *Search1Default) Error() string {
	return fmt.Sprintf("[POST /xrepository/suche][%d] search_1 default ", o._statusCode)
}

func (o *Search1Default) String() string {
	return fmt.Sprintf("[POST /xrepository/suche][%d] search_1 default ", o._statusCode)
}

func (o *Search1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}