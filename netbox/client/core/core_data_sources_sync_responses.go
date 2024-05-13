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

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// CoreDataSourcesSyncReader is a Reader for the CoreDataSourcesSync structure.
type CoreDataSourcesSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CoreDataSourcesSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCoreDataSourcesSyncCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCoreDataSourcesSyncDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCoreDataSourcesSyncCreated creates a CoreDataSourcesSyncCreated with default headers values
func NewCoreDataSourcesSyncCreated() *CoreDataSourcesSyncCreated {
	return &CoreDataSourcesSyncCreated{}
}

/*
CoreDataSourcesSyncCreated describes a response with status code 201, with default header values.

CoreDataSourcesSyncCreated core data sources sync created
*/
type CoreDataSourcesSyncCreated struct {
	Payload *models.DataSource
}

// IsSuccess returns true when this core data sources sync created response has a 2xx status code
func (o *CoreDataSourcesSyncCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this core data sources sync created response has a 3xx status code
func (o *CoreDataSourcesSyncCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this core data sources sync created response has a 4xx status code
func (o *CoreDataSourcesSyncCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this core data sources sync created response has a 5xx status code
func (o *CoreDataSourcesSyncCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this core data sources sync created response a status code equal to that given
func (o *CoreDataSourcesSyncCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the core data sources sync created response
func (o *CoreDataSourcesSyncCreated) Code() int {
	return 201
}

func (o *CoreDataSourcesSyncCreated) Error() string {
	return fmt.Sprintf("[POST /core/data-sources/{id}/sync/][%d] coreDataSourcesSyncCreated  %+v", 201, o.Payload)
}

func (o *CoreDataSourcesSyncCreated) String() string {
	return fmt.Sprintf("[POST /core/data-sources/{id}/sync/][%d] coreDataSourcesSyncCreated  %+v", 201, o.Payload)
}

func (o *CoreDataSourcesSyncCreated) GetPayload() *models.DataSource {
	return o.Payload
}

func (o *CoreDataSourcesSyncCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCoreDataSourcesSyncDefault creates a CoreDataSourcesSyncDefault with default headers values
func NewCoreDataSourcesSyncDefault(code int) *CoreDataSourcesSyncDefault {
	return &CoreDataSourcesSyncDefault{
		_statusCode: code,
	}
}

/*
CoreDataSourcesSyncDefault describes a response with status code -1, with default header values.

CoreDataSourcesSyncDefault core data sources sync default
*/
type CoreDataSourcesSyncDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this core data sources sync default response has a 2xx status code
func (o *CoreDataSourcesSyncDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this core data sources sync default response has a 3xx status code
func (o *CoreDataSourcesSyncDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this core data sources sync default response has a 4xx status code
func (o *CoreDataSourcesSyncDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this core data sources sync default response has a 5xx status code
func (o *CoreDataSourcesSyncDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this core data sources sync default response a status code equal to that given
func (o *CoreDataSourcesSyncDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the core data sources sync default response
func (o *CoreDataSourcesSyncDefault) Code() int {
	return o._statusCode
}

func (o *CoreDataSourcesSyncDefault) Error() string {
	return fmt.Sprintf("[POST /core/data-sources/{id}/sync/][%d] core_data-sources_sync default  %+v", o._statusCode, o.Payload)
}

func (o *CoreDataSourcesSyncDefault) String() string {
	return fmt.Sprintf("[POST /core/data-sources/{id}/sync/][%d] core_data-sources_sync default  %+v", o._statusCode, o.Payload)
}

func (o *CoreDataSourcesSyncDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CoreDataSourcesSyncDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
