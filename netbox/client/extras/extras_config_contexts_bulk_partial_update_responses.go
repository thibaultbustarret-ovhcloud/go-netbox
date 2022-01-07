// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasConfigContextsBulkPartialUpdateReader is a Reader for the ExtrasConfigContextsBulkPartialUpdate structure.
type ExtrasConfigContextsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasConfigContextsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasConfigContextsBulkPartialUpdateOK creates a ExtrasConfigContextsBulkPartialUpdateOK with default headers values
func NewExtrasConfigContextsBulkPartialUpdateOK() *ExtrasConfigContextsBulkPartialUpdateOK {
	return &ExtrasConfigContextsBulkPartialUpdateOK{}
}

/* ExtrasConfigContextsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextsBulkPartialUpdateOK extras config contexts bulk partial update o k
*/
type ExtrasConfigContextsBulkPartialUpdateOK struct {
	Payload *models.ConfigContext
}

func (o *ExtrasConfigContextsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-contexts/][%d] extrasConfigContextsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasConfigContextsBulkPartialUpdateOK) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasConfigContextsBulkPartialUpdateDefault creates a ExtrasConfigContextsBulkPartialUpdateDefault with default headers values
func NewExtrasConfigContextsBulkPartialUpdateDefault(code int) *ExtrasConfigContextsBulkPartialUpdateDefault {
	return &ExtrasConfigContextsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasConfigContextsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasConfigContextsBulkPartialUpdateDefault extras config contexts bulk partial update default
*/
type ExtrasConfigContextsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras config contexts bulk partial update default response
func (o *ExtrasConfigContextsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasConfigContextsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-contexts/][%d] extras_config-contexts_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasConfigContextsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigContextsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
