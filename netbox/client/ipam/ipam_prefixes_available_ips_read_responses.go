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

// IpamPrefixesAvailableIpsReadReader is a Reader for the IpamPrefixesAvailableIpsRead structure.
type IpamPrefixesAvailableIpsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailableIpsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesAvailableIpsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamPrefixesAvailableIpsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesAvailableIpsReadOK creates a IpamPrefixesAvailableIpsReadOK with default headers values
func NewIpamPrefixesAvailableIpsReadOK() *IpamPrefixesAvailableIpsReadOK {
	return &IpamPrefixesAvailableIpsReadOK{}
}

/* IpamPrefixesAvailableIpsReadOK describes a response with status code 200, with default header values.

IpamPrefixesAvailableIpsReadOK ipam prefixes available ips read o k
*/
type IpamPrefixesAvailableIpsReadOK struct {
	Payload []*models.AvailableIP
}

func (o *IpamPrefixesAvailableIpsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsReadOK  %+v", 200, o.Payload)
}
func (o *IpamPrefixesAvailableIpsReadOK) GetPayload() []*models.AvailableIP {
	return o.Payload
}

func (o *IpamPrefixesAvailableIpsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamPrefixesAvailableIpsReadDefault creates a IpamPrefixesAvailableIpsReadDefault with default headers values
func NewIpamPrefixesAvailableIpsReadDefault(code int) *IpamPrefixesAvailableIpsReadDefault {
	return &IpamPrefixesAvailableIpsReadDefault{
		_statusCode: code,
	}
}

/* IpamPrefixesAvailableIpsReadDefault describes a response with status code -1, with default header values.

IpamPrefixesAvailableIpsReadDefault ipam prefixes available ips read default
*/
type IpamPrefixesAvailableIpsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam prefixes available ips read default response
func (o *IpamPrefixesAvailableIpsReadDefault) Code() int {
	return o._statusCode
}

func (o *IpamPrefixesAvailableIpsReadDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-ips/][%d] ipam_prefixes_available-ips_read default  %+v", o._statusCode, o.Payload)
}
func (o *IpamPrefixesAvailableIpsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesAvailableIpsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
