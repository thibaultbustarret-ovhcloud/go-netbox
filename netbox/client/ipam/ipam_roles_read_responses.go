// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamRolesReadReader is a Reader for the IpamRolesRead structure.
type IpamRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRolesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRolesReadOK creates a IpamRolesReadOK with default headers values
func NewIpamRolesReadOK() *IpamRolesReadOK {
	return &IpamRolesReadOK{}
}

/* IpamRolesReadOK describes a response with status code 200, with default header values.

IpamRolesReadOK ipam roles read o k
*/
type IpamRolesReadOK struct {
	Payload *models.Role
}

func (o *IpamRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipamRolesReadOK  %+v", 200, o.Payload)
}
func (o *IpamRolesReadOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRolesReadDefault creates a IpamRolesReadDefault with default headers values
func NewIpamRolesReadDefault(code int) *IpamRolesReadDefault {
	return &IpamRolesReadDefault{
		_statusCode: code,
	}
}

/* IpamRolesReadDefault describes a response with status code -1, with default header values.

IpamRolesReadDefault ipam roles read default
*/
type IpamRolesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam roles read default response
func (o *IpamRolesReadDefault) Code() int {
	return o._statusCode
}

func (o *IpamRolesReadDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipam_roles_read default  %+v", o._statusCode, o.Payload)
}
func (o *IpamRolesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRolesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
