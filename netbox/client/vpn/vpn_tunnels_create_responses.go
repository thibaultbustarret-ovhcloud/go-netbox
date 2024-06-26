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

package vpn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// VpnTunnelsCreateReader is a Reader for the VpnTunnelsCreate structure.
type VpnTunnelsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VpnTunnelsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVpnTunnelsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVpnTunnelsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVpnTunnelsCreateCreated creates a VpnTunnelsCreateCreated with default headers values
func NewVpnTunnelsCreateCreated() *VpnTunnelsCreateCreated {
	return &VpnTunnelsCreateCreated{}
}

/*
VpnTunnelsCreateCreated describes a response with status code 201, with default header values.

VpnTunnelsCreateCreated vpn tunnels create created
*/
type VpnTunnelsCreateCreated struct {
	Payload *models.Tunnel
}

// IsSuccess returns true when this vpn tunnels create created response has a 2xx status code
func (o *VpnTunnelsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vpn tunnels create created response has a 3xx status code
func (o *VpnTunnelsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vpn tunnels create created response has a 4xx status code
func (o *VpnTunnelsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this vpn tunnels create created response has a 5xx status code
func (o *VpnTunnelsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this vpn tunnels create created response a status code equal to that given
func (o *VpnTunnelsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the vpn tunnels create created response
func (o *VpnTunnelsCreateCreated) Code() int {
	return 201
}

func (o *VpnTunnelsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vpn/tunnels/][%d] vpnTunnelsCreateCreated %s", 201, payload)
}

func (o *VpnTunnelsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vpn/tunnels/][%d] vpnTunnelsCreateCreated %s", 201, payload)
}

func (o *VpnTunnelsCreateCreated) GetPayload() *models.Tunnel {
	return o.Payload
}

func (o *VpnTunnelsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tunnel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVpnTunnelsCreateDefault creates a VpnTunnelsCreateDefault with default headers values
func NewVpnTunnelsCreateDefault(code int) *VpnTunnelsCreateDefault {
	return &VpnTunnelsCreateDefault{
		_statusCode: code,
	}
}

/*
VpnTunnelsCreateDefault describes a response with status code -1, with default header values.

VpnTunnelsCreateDefault vpn tunnels create default
*/
type VpnTunnelsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this vpn tunnels create default response has a 2xx status code
func (o *VpnTunnelsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vpn tunnels create default response has a 3xx status code
func (o *VpnTunnelsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vpn tunnels create default response has a 4xx status code
func (o *VpnTunnelsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vpn tunnels create default response has a 5xx status code
func (o *VpnTunnelsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vpn tunnels create default response a status code equal to that given
func (o *VpnTunnelsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vpn tunnels create default response
func (o *VpnTunnelsCreateDefault) Code() int {
	return o._statusCode
}

func (o *VpnTunnelsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vpn/tunnels/][%d] vpn_tunnels_create default %s", o._statusCode, payload)
}

func (o *VpnTunnelsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vpn/tunnels/][%d] vpn_tunnels_create default %s", o._statusCode, payload)
}

func (o *VpnTunnelsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VpnTunnelsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
