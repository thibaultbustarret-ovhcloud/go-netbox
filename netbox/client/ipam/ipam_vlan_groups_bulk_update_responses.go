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

// IpamVlanGroupsBulkUpdateReader is a Reader for the IpamVlanGroupsBulkUpdate structure.
type IpamVlanGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlanGroupsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsBulkUpdateOK creates a IpamVlanGroupsBulkUpdateOK with default headers values
func NewIpamVlanGroupsBulkUpdateOK() *IpamVlanGroupsBulkUpdateOK {
	return &IpamVlanGroupsBulkUpdateOK{}
}

/* IpamVlanGroupsBulkUpdateOK describes a response with status code 200, with default header values.

IpamVlanGroupsBulkUpdateOK ipam vlan groups bulk update o k
*/
type IpamVlanGroupsBulkUpdateOK struct {
	Payload *models.VLANGroup
}

func (o *IpamVlanGroupsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/][%d] ipamVlanGroupsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVlanGroupsBulkUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlanGroupsBulkUpdateDefault creates a IpamVlanGroupsBulkUpdateDefault with default headers values
func NewIpamVlanGroupsBulkUpdateDefault(code int) *IpamVlanGroupsBulkUpdateDefault {
	return &IpamVlanGroupsBulkUpdateDefault{
		_statusCode: code,
	}
}

/* IpamVlanGroupsBulkUpdateDefault describes a response with status code -1, with default header values.

IpamVlanGroupsBulkUpdateDefault ipam vlan groups bulk update default
*/
type IpamVlanGroupsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlan groups bulk update default response
func (o *IpamVlanGroupsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlanGroupsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/][%d] ipam_vlan-groups_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *IpamVlanGroupsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
