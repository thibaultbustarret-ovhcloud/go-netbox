// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimDevicesBulkUpdateReader is a Reader for the DcimDevicesBulkUpdate structure.
type DcimDevicesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDevicesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDevicesBulkUpdateOK creates a DcimDevicesBulkUpdateOK with default headers values
func NewDcimDevicesBulkUpdateOK() *DcimDevicesBulkUpdateOK {
	return &DcimDevicesBulkUpdateOK{}
}

/* DcimDevicesBulkUpdateOK describes a response with status code 200, with default header values.

DcimDevicesBulkUpdateOK dcim devices bulk update o k
*/
type DcimDevicesBulkUpdateOK struct {
	Payload *models.DeviceWithConfigContext
}

func (o *DcimDevicesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/devices/][%d] dcimDevicesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDevicesBulkUpdateOK) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDevicesBulkUpdateDefault creates a DcimDevicesBulkUpdateDefault with default headers values
func NewDcimDevicesBulkUpdateDefault(code int) *DcimDevicesBulkUpdateDefault {
	return &DcimDevicesBulkUpdateDefault{
		_statusCode: code,
	}
}

/* DcimDevicesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimDevicesBulkUpdateDefault dcim devices bulk update default
*/
type DcimDevicesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim devices bulk update default response
func (o *DcimDevicesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDevicesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/devices/][%d] dcim_devices_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimDevicesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
