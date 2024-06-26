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
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new vpn API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new vpn API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new vpn API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for vpn API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	VpnTunnelGroupsCreate(params *VpnTunnelGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelGroupsCreateCreated, error)

	VpnTunnelGroupsDelete(params *VpnTunnelGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelGroupsDeleteNoContent, error)

	VpnTunnelGroupsList(params *VpnTunnelGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelGroupsListOK, error)

	VpnTunnelGroupsRead(params *VpnTunnelGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelGroupsReadOK, error)

	VpnTunnelGroupsUpdate(params *VpnTunnelGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelGroupsUpdateOK, error)

	VpnTunnelTerminationsCreate(params *VpnTunnelTerminationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelTerminationsCreateCreated, error)

	VpnTunnelTerminationsDelete(params *VpnTunnelTerminationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelTerminationsDeleteNoContent, error)

	VpnTunnelTerminationsList(params *VpnTunnelTerminationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelTerminationsListOK, error)

	VpnTunnelTerminationsRead(params *VpnTunnelTerminationsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelTerminationsReadOK, error)

	VpnTunnelTerminationsUpdate(params *VpnTunnelTerminationsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelTerminationsUpdateOK, error)

	VpnTunnelsCreate(params *VpnTunnelsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelsCreateCreated, error)

	VpnTunnelsDelete(params *VpnTunnelsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelsDeleteNoContent, error)

	VpnTunnelsList(params *VpnTunnelsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelsListOK, error)

	VpnTunnelsRead(params *VpnTunnelsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelsReadOK, error)

	VpnTunnelsUpdate(params *VpnTunnelsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
VpnTunnelGroupsCreate vpn tunnel groups create API
*/
func (a *Client) VpnTunnelGroupsCreate(params *VpnTunnelGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelGroupsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelGroupsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnel-groups_create",
		Method:             "POST",
		PathPattern:        "/vpn/tunnel-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelGroupsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelGroupsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelGroupsCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelGroupsDelete vpn tunnel groups delete API
*/
func (a *Client) VpnTunnelGroupsDelete(params *VpnTunnelGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelGroupsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelGroupsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnel-groups_delete",
		Method:             "DELETE",
		PathPattern:        "/vpn/tunnel-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelGroupsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelGroupsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelGroupsDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelGroupsList vpn tunnel groups list API
*/
func (a *Client) VpnTunnelGroupsList(params *VpnTunnelGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelGroupsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnel-groups_list",
		Method:             "GET",
		PathPattern:        "/vpn/tunnel-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelGroupsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelGroupsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelGroupsListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelGroupsRead vpn tunnel groups read API
*/
func (a *Client) VpnTunnelGroupsRead(params *VpnTunnelGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelGroupsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelGroupsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnel-groups_read",
		Method:             "GET",
		PathPattern:        "/vpn/tunnel-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelGroupsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelGroupsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelGroupsReadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelGroupsUpdate vpn tunnel groups update API
*/
func (a *Client) VpnTunnelGroupsUpdate(params *VpnTunnelGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelGroupsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelGroupsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnel-groups_update",
		Method:             "PUT",
		PathPattern:        "/vpn/tunnel-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelGroupsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelGroupsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelGroupsUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelTerminationsCreate Post a tunnel termination object.
*/
func (a *Client) VpnTunnelTerminationsCreate(params *VpnTunnelTerminationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelTerminationsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelTerminationsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnel-terminations_create",
		Method:             "POST",
		PathPattern:        "/vpn/tunnel-terminations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelTerminationsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelTerminationsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelTerminationsCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelTerminationsDelete vpn tunnel terminations delete API
*/
func (a *Client) VpnTunnelTerminationsDelete(params *VpnTunnelTerminationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelTerminationsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelTerminationsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnel-terminations_delete",
		Method:             "DELETE",
		PathPattern:        "/vpn/tunnel-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelTerminationsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelTerminationsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelTerminationsDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelTerminationsList vpn tunnel terminations list API
*/
func (a *Client) VpnTunnelTerminationsList(params *VpnTunnelTerminationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelTerminationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelTerminationsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnel-terminations_list",
		Method:             "GET",
		PathPattern:        "/vpn/tunnel-terminations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelTerminationsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelTerminationsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelTerminationsListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelTerminationsRead vpn tunnel terminations read API
*/
func (a *Client) VpnTunnelTerminationsRead(params *VpnTunnelTerminationsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelTerminationsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelTerminationsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnel-terminations_read",
		Method:             "GET",
		PathPattern:        "/vpn/tunnel-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelTerminationsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelTerminationsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelTerminationsReadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelTerminationsUpdate vpn tunnel terminations update API
*/
func (a *Client) VpnTunnelTerminationsUpdate(params *VpnTunnelTerminationsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelTerminationsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelTerminationsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnel-terminations_update",
		Method:             "PUT",
		PathPattern:        "/vpn/tunnel-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelTerminationsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelTerminationsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelTerminationsUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelsCreate Post a tunnel object.
*/
func (a *Client) VpnTunnelsCreate(params *VpnTunnelsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnels_create",
		Method:             "POST",
		PathPattern:        "/vpn/tunnels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelsCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelsDelete Delete a tunnel object.
*/
func (a *Client) VpnTunnelsDelete(params *VpnTunnelsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnels_delete",
		Method:             "DELETE",
		PathPattern:        "/vpn/tunnels/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelsDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelsList Get a list of tunnel objects.
*/
func (a *Client) VpnTunnelsList(params *VpnTunnelsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnels_list",
		Method:             "GET",
		PathPattern:        "/vpn/tunnels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelsListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelsRead Get a tunnel object.
*/
func (a *Client) VpnTunnelsRead(params *VpnTunnelsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnels_read",
		Method:             "GET",
		PathPattern:        "/vpn/tunnels/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelsReadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VpnTunnelsUpdate Put a tunnel object.
*/
func (a *Client) VpnTunnelsUpdate(params *VpnTunnelsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VpnTunnelsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVpnTunnelsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "vpn_tunnels_update",
		Method:             "PUT",
		PathPattern:        "/vpn/tunnels/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VpnTunnelsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VpnTunnelsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VpnTunnelsUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
