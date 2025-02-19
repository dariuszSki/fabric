// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TerminatorCreate terminator create
//
// swagger:model terminatorCreate
type TerminatorCreate struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// binding
	// Required: true
	Binding *string `json:"binding"`

	// cost
	Cost *TerminatorCost `json:"cost,omitempty"`

	// instance Id
	InstanceID string `json:"instanceId,omitempty"`

	// instance secret
	// Format: byte
	InstanceSecret strfmt.Base64 `json:"instanceSecret,omitempty"`

	// precedence
	Precedence TerminatorPrecedence `json:"precedence,omitempty"`

	// router
	// Required: true
	Router *string `json:"router"`

	// service
	// Required: true
	Service *string `json:"service"`

	// tags
	Tags *Tags `json:"tags,omitempty"`
}

// Validate validates this terminator create
func (m *TerminatorCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBinding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrecedence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerminatorCreate) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *TerminatorCreate) validateBinding(formats strfmt.Registry) error {

	if err := validate.Required("binding", "body", m.Binding); err != nil {
		return err
	}

	return nil
}

func (m *TerminatorCreate) validateCost(formats strfmt.Registry) error {
	if swag.IsZero(m.Cost) { // not required
		return nil
	}

	if m.Cost != nil {
		if err := m.Cost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cost")
			}
			return err
		}
	}

	return nil
}

func (m *TerminatorCreate) validatePrecedence(formats strfmt.Registry) error {
	if swag.IsZero(m.Precedence) { // not required
		return nil
	}

	if err := m.Precedence.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("precedence")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("precedence")
		}
		return err
	}

	return nil
}

func (m *TerminatorCreate) validateRouter(formats strfmt.Registry) error {

	if err := validate.Required("router", "body", m.Router); err != nil {
		return err
	}

	return nil
}

func (m *TerminatorCreate) validateService(formats strfmt.Registry) error {

	if err := validate.Required("service", "body", m.Service); err != nil {
		return err
	}

	return nil
}

func (m *TerminatorCreate) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if m.Tags != nil {
		if err := m.Tags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this terminator create based on the context it is used
func (m *TerminatorCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrecedence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerminatorCreate) contextValidateCost(ctx context.Context, formats strfmt.Registry) error {

	if m.Cost != nil {
		if err := m.Cost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cost")
			}
			return err
		}
	}

	return nil
}

func (m *TerminatorCreate) contextValidatePrecedence(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Precedence.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("precedence")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("precedence")
		}
		return err
	}

	return nil
}

func (m *TerminatorCreate) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	if m.Tags != nil {
		if err := m.Tags.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerminatorCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerminatorCreate) UnmarshalBinary(b []byte) error {
	var res TerminatorCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
