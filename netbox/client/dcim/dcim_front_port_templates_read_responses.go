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

// DcimFrontPortTemplatesReadReader is a Reader for the DcimFrontPortTemplatesRead structure.
type DcimFrontPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortTemplatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortTemplatesReadOK creates a DcimFrontPortTemplatesReadOK with default headers values
func NewDcimFrontPortTemplatesReadOK() *DcimFrontPortTemplatesReadOK {
	return &DcimFrontPortTemplatesReadOK{}
}

/* DcimFrontPortTemplatesReadOK describes a response with status code 200, with default header values.

DcimFrontPortTemplatesReadOK dcim front port templates read o k
*/
type DcimFrontPortTemplatesReadOK struct {
	Payload *models.FrontPortTemplate
}

func (o *DcimFrontPortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesReadOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortTemplatesReadOK) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortTemplatesReadDefault creates a DcimFrontPortTemplatesReadDefault with default headers values
func NewDcimFrontPortTemplatesReadDefault(code int) *DcimFrontPortTemplatesReadDefault {
	return &DcimFrontPortTemplatesReadDefault{
		_statusCode: code,
	}
}

/* DcimFrontPortTemplatesReadDefault describes a response with status code -1, with default header values.

DcimFrontPortTemplatesReadDefault dcim front port templates read default
*/
type DcimFrontPortTemplatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front port templates read default response
func (o *DcimFrontPortTemplatesReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortTemplatesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/front-port-templates/{id}/][%d] dcim_front-port-templates_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimFrontPortTemplatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortTemplatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
