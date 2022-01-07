// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// UsersTokensCreateReader is a Reader for the UsersTokensCreate structure.
type UsersTokensCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersTokensCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersTokensCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersTokensCreateCreated creates a UsersTokensCreateCreated with default headers values
func NewUsersTokensCreateCreated() *UsersTokensCreateCreated {
	return &UsersTokensCreateCreated{}
}

/* UsersTokensCreateCreated describes a response with status code 201, with default header values.

UsersTokensCreateCreated users tokens create created
*/
type UsersTokensCreateCreated struct {
	Payload *models.Token
}

func (o *UsersTokensCreateCreated) Error() string {
	return fmt.Sprintf("[POST /users/tokens/][%d] usersTokensCreateCreated  %+v", 201, o.Payload)
}
func (o *UsersTokensCreateCreated) GetPayload() *models.Token {
	return o.Payload
}

func (o *UsersTokensCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensCreateDefault creates a UsersTokensCreateDefault with default headers values
func NewUsersTokensCreateDefault(code int) *UsersTokensCreateDefault {
	return &UsersTokensCreateDefault{
		_statusCode: code,
	}
}

/* UsersTokensCreateDefault describes a response with status code -1, with default header values.

UsersTokensCreateDefault users tokens create default
*/
type UsersTokensCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users tokens create default response
func (o *UsersTokensCreateDefault) Code() int {
	return o._statusCode
}

func (o *UsersTokensCreateDefault) Error() string {
	return fmt.Sprintf("[POST /users/tokens/][%d] users_tokens_create default  %+v", o._statusCode, o.Payload)
}
func (o *UsersTokensCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
