// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasConfigTemplatesListReader is a Reader for the ExtrasConfigTemplatesList structure.
type ExtrasConfigTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasConfigTemplatesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasConfigTemplatesListOK creates a ExtrasConfigTemplatesListOK with default headers values
func NewExtrasConfigTemplatesListOK() *ExtrasConfigTemplatesListOK {
	return &ExtrasConfigTemplatesListOK{}
}

/*
ExtrasConfigTemplatesListOK describes a response with status code 200, with default header values.

ExtrasConfigTemplatesListOK extras config templates list o k
*/
type ExtrasConfigTemplatesListOK struct {
	Payload *ExtrasConfigTemplatesListOKBody
}

// IsSuccess returns true when this extras config templates list o k response has a 2xx status code
func (o *ExtrasConfigTemplatesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config templates list o k response has a 3xx status code
func (o *ExtrasConfigTemplatesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config templates list o k response has a 4xx status code
func (o *ExtrasConfigTemplatesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config templates list o k response has a 5xx status code
func (o *ExtrasConfigTemplatesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config templates list o k response a status code equal to that given
func (o *ExtrasConfigTemplatesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras config templates list o k response
func (o *ExtrasConfigTemplatesListOK) Code() int {
	return 200
}

func (o *ExtrasConfigTemplatesListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/config-templates/][%d] extrasConfigTemplatesListOK %s", 200, payload)
}

func (o *ExtrasConfigTemplatesListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/config-templates/][%d] extrasConfigTemplatesListOK %s", 200, payload)
}

func (o *ExtrasConfigTemplatesListOK) GetPayload() *ExtrasConfigTemplatesListOKBody {
	return o.Payload
}

func (o *ExtrasConfigTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExtrasConfigTemplatesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasConfigTemplatesListDefault creates a ExtrasConfigTemplatesListDefault with default headers values
func NewExtrasConfigTemplatesListDefault(code int) *ExtrasConfigTemplatesListDefault {
	return &ExtrasConfigTemplatesListDefault{
		_statusCode: code,
	}
}

/*
ExtrasConfigTemplatesListDefault describes a response with status code -1, with default header values.

ExtrasConfigTemplatesListDefault extras config templates list default
*/
type ExtrasConfigTemplatesListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras config templates list default response has a 2xx status code
func (o *ExtrasConfigTemplatesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras config templates list default response has a 3xx status code
func (o *ExtrasConfigTemplatesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras config templates list default response has a 4xx status code
func (o *ExtrasConfigTemplatesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras config templates list default response has a 5xx status code
func (o *ExtrasConfigTemplatesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras config templates list default response a status code equal to that given
func (o *ExtrasConfigTemplatesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras config templates list default response
func (o *ExtrasConfigTemplatesListDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasConfigTemplatesListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/config-templates/][%d] extras_config-templates_list default %s", o._statusCode, payload)
}

func (o *ExtrasConfigTemplatesListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/config-templates/][%d] extras_config-templates_list default %s", o._statusCode, payload)
}

func (o *ExtrasConfigTemplatesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigTemplatesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ExtrasConfigTemplatesListOKBody extras config templates list o k body
swagger:model ExtrasConfigTemplatesListOKBody
*/
type ExtrasConfigTemplatesListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.ConfigTemplate `json:"results"`
}

// Validate validates this extras config templates list o k body
func (o *ExtrasConfigTemplatesListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExtrasConfigTemplatesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("extrasConfigTemplatesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasConfigTemplatesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasConfigTemplatesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasConfigTemplatesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasConfigTemplatesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasConfigTemplatesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("extrasConfigTemplatesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasConfigTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasConfigTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this extras config templates list o k body based on the context it is used
func (o *ExtrasConfigTemplatesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExtrasConfigTemplatesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasConfigTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasConfigTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExtrasConfigTemplatesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExtrasConfigTemplatesListOKBody) UnmarshalBinary(b []byte) error {
	var res ExtrasConfigTemplatesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}