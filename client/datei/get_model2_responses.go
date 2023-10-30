// Code generated by go-swagger; DO NOT EDIT.

package datei

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetModel2Reader is a Reader for the GetModel2 structure.
type GetModel2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModel2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetModel2Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetModel2Default creates a GetModel2Default with default headers values
func NewGetModel2Default(code int) *GetModel2Default {
	return &GetModel2Default{
		_statusCode: code,
	}
}

/*
GetModel2Default describes a response with status code -1, with default header values.

successful operation
*/
type GetModel2Default struct {
	_statusCode int
}

// IsSuccess returns true when this get model 2 default response has a 2xx status code
func (o *GetModel2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get model 2 default response has a 3xx status code
func (o *GetModel2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get model 2 default response has a 4xx status code
func (o *GetModel2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get model 2 default response has a 5xx status code
func (o *GetModel2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get model 2 default response a status code equal to that given
func (o *GetModel2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get model 2 default response
func (o *GetModel2Default) Code() int {
	return o._statusCode
}

func (o *GetModel2Default) Error() string {
	return fmt.Sprintf("[GET /datei/{kennung}/model][%d] getModel_2 default ", o._statusCode)
}

func (o *GetModel2Default) String() string {
	return fmt.Sprintf("[GET /datei/{kennung}/model][%d] getModel_2 default ", o._statusCode)
}

func (o *GetModel2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
