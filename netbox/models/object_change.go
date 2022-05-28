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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ObjectChange object change
//
// swagger:model ObjectChange
type ObjectChange struct {

	// action
	Action *ObjectChangeAction `json:"action,omitempty"`

	// Changed object
	//
	//
	// Serialize a nested representation of the changed object.
	//
	// Read Only: true
	ChangedObject map[string]*string `json:"changed_object,omitempty"`

	// Changed object id
	// Required: true
	// Minimum: 0
	ChangedObjectID *int64 `json:"changed_object_id"`

	// Changed object type
	// Read Only: true
	ChangedObjectType string `json:"changed_object_type,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Postchange data
	// Read Only: true
	PostchangeData string `json:"postchange_data,omitempty"`

	// Prechange data
	// Read Only: true
	PrechangeData string `json:"prechange_data,omitempty"`

	// Request id
	// Read Only: true
	// Format: uuid
	RequestID strfmt.UUID `json:"request_id,omitempty"`

	// Time
	// Read Only: true
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// user
	User *NestedUser `json:"user,omitempty"`

	// User name
	// Read Only: true
	// Min Length: 1
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this object change
func (m *ObjectChange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangedObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectChange) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectChange) validateChangedObjectID(formats strfmt.Registry) error {

	if err := validate.Required("changed_object_id", "body", m.ChangedObjectID); err != nil {
		return err
	}

	if err := validate.MinimumInt("changed_object_id", "body", *m.ChangedObjectID, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) validateRequestID(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestID) { // not required
		return nil
	}

	if err := validate.FormatOf("request_id", "body", "uuid", m.RequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectChange) validateUserName(formats strfmt.Registry) error {
	if swag.IsZero(m.UserName) { // not required
		return nil
	}

	if err := validate.MinLength("user_name", "body", m.UserName, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this object change based on the context it is used
func (m *ObjectChange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChangedObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChangedObjectType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostchangeData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrechangeData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectChange) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if m.Action != nil {
		if err := m.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectChange) contextValidateChangedObject(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ObjectChange) contextValidateChangedObjectType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "changed_object_type", "body", string(m.ChangedObjectType)); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) contextValidatePostchangeData(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "postchange_data", "body", string(m.PostchangeData)); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) contextValidatePrechangeData(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "prechange_data", "body", string(m.PrechangeData)); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) contextValidateRequestID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "request_id", "body", strfmt.UUID(m.RequestID)); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "time", "body", strfmt.DateTime(m.Time)); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectChange) contextValidateUserName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "user_name", "body", string(m.UserName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectChange) UnmarshalBinary(b []byte) error {
	var res ObjectChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ObjectChangeAction Action
//
// swagger:model ObjectChangeAction
type ObjectChangeAction struct {

	// label
	// Required: true
	// Enum: [Created Updated Deleted]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [create update delete]
	Value *string `json:"value"`
}

// Validate validates this object change action
func (m *ObjectChangeAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var objectChangeActionTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Created","Updated","Deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectChangeActionTypeLabelPropEnum = append(objectChangeActionTypeLabelPropEnum, v)
	}
}

const (

	// ObjectChangeActionLabelCreated captures enum value "Created"
	ObjectChangeActionLabelCreated string = "Created"

	// ObjectChangeActionLabelUpdated captures enum value "Updated"
	ObjectChangeActionLabelUpdated string = "Updated"

	// ObjectChangeActionLabelDeleted captures enum value "Deleted"
	ObjectChangeActionLabelDeleted string = "Deleted"
)

// prop value enum
func (m *ObjectChangeAction) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectChangeActionTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectChangeAction) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("action"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("action"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var objectChangeActionTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["create","update","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectChangeActionTypeValuePropEnum = append(objectChangeActionTypeValuePropEnum, v)
	}
}

const (

	// ObjectChangeActionValueCreate captures enum value "create"
	ObjectChangeActionValueCreate string = "create"

	// ObjectChangeActionValueUpdate captures enum value "update"
	ObjectChangeActionValueUpdate string = "update"

	// ObjectChangeActionValueDelete captures enum value "delete"
	ObjectChangeActionValueDelete string = "delete"
)

// prop value enum
func (m *ObjectChangeAction) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectChangeActionTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectChangeAction) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("action"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("action"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this object change action based on the context it is used
func (m *ObjectChangeAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectChangeAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectChangeAction) UnmarshalBinary(b []byte) error {
	var res ObjectChangeAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
