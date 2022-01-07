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

// DcimPowerOutletTemplatesBulkPartialUpdateReader is a Reader for the DcimPowerOutletTemplatesBulkPartialUpdate structure.
type DcimPowerOutletTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletTemplatesBulkPartialUpdateOK creates a DcimPowerOutletTemplatesBulkPartialUpdateOK with default headers values
func NewDcimPowerOutletTemplatesBulkPartialUpdateOK() *DcimPowerOutletTemplatesBulkPartialUpdateOK {
	return &DcimPowerOutletTemplatesBulkPartialUpdateOK{}
}

/* DcimPowerOutletTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletTemplatesBulkPartialUpdateOK dcim power outlet templates bulk partial update o k
*/
type DcimPowerOutletTemplatesBulkPartialUpdateOK struct {
	Payload *models.PowerOutletTemplate
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletTemplatesBulkPartialUpdateDefault creates a DcimPowerOutletTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimPowerOutletTemplatesBulkPartialUpdateDefault(code int) *DcimPowerOutletTemplatesBulkPartialUpdateDefault {
	return &DcimPowerOutletTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimPowerOutletTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerOutletTemplatesBulkPartialUpdateDefault dcim power outlet templates bulk partial update default
*/
type DcimPowerOutletTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power outlet templates bulk partial update default response
func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlet-templates/][%d] dcim_power-outlet-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
