// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimInventoryItemsBulkDeleteReader is a Reader for the DcimInventoryItemsBulkDelete structure.
type DcimInventoryItemsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInventoryItemsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemsBulkDeleteNoContent creates a DcimInventoryItemsBulkDeleteNoContent with default headers values
func NewDcimInventoryItemsBulkDeleteNoContent() *DcimInventoryItemsBulkDeleteNoContent {
	return &DcimInventoryItemsBulkDeleteNoContent{}
}

/* DcimInventoryItemsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimInventoryItemsBulkDeleteNoContent dcim inventory items bulk delete no content
*/
type DcimInventoryItemsBulkDeleteNoContent struct {
}

func (o *DcimInventoryItemsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/][%d] dcimInventoryItemsBulkDeleteNoContent ", 204)
}

func (o *DcimInventoryItemsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimInventoryItemsBulkDeleteDefault creates a DcimInventoryItemsBulkDeleteDefault with default headers values
func NewDcimInventoryItemsBulkDeleteDefault(code int) *DcimInventoryItemsBulkDeleteDefault {
	return &DcimInventoryItemsBulkDeleteDefault{
		_statusCode: code,
	}
}

/* DcimInventoryItemsBulkDeleteDefault describes a response with status code -1, with default header values.

DcimInventoryItemsBulkDeleteDefault dcim inventory items bulk delete default
*/
type DcimInventoryItemsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory items bulk delete default response
func (o *DcimInventoryItemsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/][%d] dcim_inventory-items_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInventoryItemsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
