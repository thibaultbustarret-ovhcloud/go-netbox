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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UsersConfigList(params *UsersConfigListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersConfigListOK, error)

	UsersGroupsCreate(params *UsersGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsCreateCreated, error)

	UsersGroupsDelete(params *UsersGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsDeleteNoContent, error)

	UsersGroupsList(params *UsersGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsListOK, error)

	UsersGroupsPartialUpdate(params *UsersGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsPartialUpdateOK, error)

	UsersGroupsRead(params *UsersGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsReadOK, error)

	UsersGroupsUpdate(params *UsersGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsUpdateOK, error)

	UsersPermissionsCreate(params *UsersPermissionsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsCreateCreated, error)

	UsersPermissionsDelete(params *UsersPermissionsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsDeleteNoContent, error)

	UsersPermissionsList(params *UsersPermissionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsListOK, error)

	UsersPermissionsPartialUpdate(params *UsersPermissionsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsPartialUpdateOK, error)

	UsersPermissionsRead(params *UsersPermissionsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsReadOK, error)

	UsersPermissionsUpdate(params *UsersPermissionsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsUpdateOK, error)

	UsersTokensCreate(params *UsersTokensCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensCreateCreated, error)

	UsersTokensDelete(params *UsersTokensDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensDeleteNoContent, error)

	UsersTokensList(params *UsersTokensListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensListOK, error)

	UsersTokensPartialUpdate(params *UsersTokensPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensPartialUpdateOK, error)

	UsersTokensProvisionCreate(params *UsersTokensProvisionCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensProvisionCreateCreated, error)

	UsersTokensRead(params *UsersTokensReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensReadOK, error)

	UsersTokensUpdate(params *UsersTokensUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensUpdateOK, error)

	UsersUsersCreate(params *UsersUsersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersCreateCreated, error)

	UsersUsersDelete(params *UsersUsersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersDeleteNoContent, error)

	UsersUsersList(params *UsersUsersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersListOK, error)

	UsersUsersPartialUpdate(params *UsersUsersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersPartialUpdateOK, error)

	UsersUsersRead(params *UsersUsersReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersReadOK, error)

	UsersUsersUpdate(params *UsersUsersUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
UsersConfigList Return the UserConfig for the currently authenticated User.
*/
func (a *Client) UsersConfigList(params *UsersConfigListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersConfigListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersConfigListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_config_list",
		Method:             "GET",
		PathPattern:        "/users/config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersConfigListReader{formats: a.formats},
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
	success, ok := result.(*UsersConfigListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersConfigListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersGroupsCreate users groups create API
*/
func (a *Client) UsersGroupsCreate(params *UsersGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_create",
		Method:             "POST",
		PathPattern:        "/users/groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersGroupsCreateReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersGroupsCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersGroupsDelete users groups delete API
*/
func (a *Client) UsersGroupsDelete(params *UsersGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_delete",
		Method:             "DELETE",
		PathPattern:        "/users/groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersGroupsDeleteReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersGroupsDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersGroupsList users groups list API
*/
func (a *Client) UsersGroupsList(params *UsersGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_list",
		Method:             "GET",
		PathPattern:        "/users/groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersGroupsListReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersGroupsListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersGroupsPartialUpdate users groups partial update API
*/
func (a *Client) UsersGroupsPartialUpdate(params *UsersGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_partial_update",
		Method:             "PATCH",
		PathPattern:        "/users/groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersGroupsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersGroupsPartialUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersGroupsRead users groups read API
*/
func (a *Client) UsersGroupsRead(params *UsersGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_read",
		Method:             "GET",
		PathPattern:        "/users/groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersGroupsReadReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersGroupsReadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersGroupsUpdate users groups update API
*/
func (a *Client) UsersGroupsUpdate(params *UsersGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_update",
		Method:             "PUT",
		PathPattern:        "/users/groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersGroupsUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersGroupsUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersPermissionsCreate users permissions create API
*/
func (a *Client) UsersPermissionsCreate(params *UsersPermissionsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_create",
		Method:             "POST",
		PathPattern:        "/users/permissions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersPermissionsCreateReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersPermissionsCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersPermissionsDelete users permissions delete API
*/
func (a *Client) UsersPermissionsDelete(params *UsersPermissionsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_delete",
		Method:             "DELETE",
		PathPattern:        "/users/permissions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersPermissionsDeleteReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersPermissionsDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersPermissionsList users permissions list API
*/
func (a *Client) UsersPermissionsList(params *UsersPermissionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_list",
		Method:             "GET",
		PathPattern:        "/users/permissions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersPermissionsListReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersPermissionsListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersPermissionsPartialUpdate users permissions partial update API
*/
func (a *Client) UsersPermissionsPartialUpdate(params *UsersPermissionsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_partial_update",
		Method:             "PATCH",
		PathPattern:        "/users/permissions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersPermissionsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersPermissionsPartialUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersPermissionsRead users permissions read API
*/
func (a *Client) UsersPermissionsRead(params *UsersPermissionsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_read",
		Method:             "GET",
		PathPattern:        "/users/permissions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersPermissionsReadReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersPermissionsReadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersPermissionsUpdate users permissions update API
*/
func (a *Client) UsersPermissionsUpdate(params *UsersPermissionsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_update",
		Method:             "PUT",
		PathPattern:        "/users/permissions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersPermissionsUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersPermissionsUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersTokensCreate users tokens create API
*/
func (a *Client) UsersTokensCreate(params *UsersTokensCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersTokensCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_tokens_create",
		Method:             "POST",
		PathPattern:        "/users/tokens/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersTokensCreateReader{formats: a.formats},
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
	success, ok := result.(*UsersTokensCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersTokensCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersTokensDelete users tokens delete API
*/
func (a *Client) UsersTokensDelete(params *UsersTokensDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersTokensDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_tokens_delete",
		Method:             "DELETE",
		PathPattern:        "/users/tokens/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersTokensDeleteReader{formats: a.formats},
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
	success, ok := result.(*UsersTokensDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersTokensDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersTokensList users tokens list API
*/
func (a *Client) UsersTokensList(params *UsersTokensListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersTokensListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_tokens_list",
		Method:             "GET",
		PathPattern:        "/users/tokens/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersTokensListReader{formats: a.formats},
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
	success, ok := result.(*UsersTokensListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersTokensListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersTokensPartialUpdate users tokens partial update API
*/
func (a *Client) UsersTokensPartialUpdate(params *UsersTokensPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersTokensPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_tokens_partial_update",
		Method:             "PATCH",
		PathPattern:        "/users/tokens/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersTokensPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersTokensPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersTokensPartialUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersTokensProvisionCreate Non-authenticated REST API endpoint via which a user may create a Token.
*/
func (a *Client) UsersTokensProvisionCreate(params *UsersTokensProvisionCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensProvisionCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersTokensProvisionCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_tokens_provision_create",
		Method:             "POST",
		PathPattern:        "/users/tokens/provision/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersTokensProvisionCreateReader{formats: a.formats},
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
	success, ok := result.(*UsersTokensProvisionCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersTokensProvisionCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersTokensRead users tokens read API
*/
func (a *Client) UsersTokensRead(params *UsersTokensReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersTokensReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_tokens_read",
		Method:             "GET",
		PathPattern:        "/users/tokens/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersTokensReadReader{formats: a.formats},
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
	success, ok := result.(*UsersTokensReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersTokensReadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersTokensUpdate users tokens update API
*/
func (a *Client) UsersTokensUpdate(params *UsersTokensUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersTokensUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersTokensUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_tokens_update",
		Method:             "PUT",
		PathPattern:        "/users/tokens/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersTokensUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersTokensUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersTokensUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersUsersCreate users users create API
*/
func (a *Client) UsersUsersCreate(params *UsersUsersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_create",
		Method:             "POST",
		PathPattern:        "/users/users/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersUsersCreateReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersUsersCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersUsersDelete users users delete API
*/
func (a *Client) UsersUsersDelete(params *UsersUsersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_delete",
		Method:             "DELETE",
		PathPattern:        "/users/users/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersUsersDeleteReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersUsersDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersUsersList users users list API
*/
func (a *Client) UsersUsersList(params *UsersUsersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_list",
		Method:             "GET",
		PathPattern:        "/users/users/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersUsersListReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersUsersListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersUsersPartialUpdate users users partial update API
*/
func (a *Client) UsersUsersPartialUpdate(params *UsersUsersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_partial_update",
		Method:             "PATCH",
		PathPattern:        "/users/users/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersUsersPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersUsersPartialUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersUsersRead users users read API
*/
func (a *Client) UsersUsersRead(params *UsersUsersReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_read",
		Method:             "GET",
		PathPattern:        "/users/users/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersUsersReadReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersUsersReadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsersUsersUpdate users users update API
*/
func (a *Client) UsersUsersUpdate(params *UsersUsersUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_update",
		Method:             "PUT",
		PathPattern:        "/users/users/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UsersUsersUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsersUsersUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
