// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceDistTrafficMetrics service dist traffic metrics
//
// swagger:model ServiceDistTrafficMetrics
type ServiceDistTrafficMetrics map[string]ServiceDistTrafficMetricsAnon

// Validate validates this service dist traffic metrics
func (m ServiceDistTrafficMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if swag.IsZero(m[k]) { // not required
			continue
		}
		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(k)
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this service dist traffic metrics based on the context it is used
func (m ServiceDistTrafficMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if val, ok := m[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ServiceDistTrafficMetricsAnon service dist traffic metrics anon
//
// swagger:model ServiceDistTrafficMetricsAnon
type ServiceDistTrafficMetricsAnon struct {

	// ratio
	Ratio float64 `json:"ratio,omitempty"`

	// value
	Value float64 `json:"value,omitempty"`
}

// Validate validates this service dist traffic metrics anon
func (m *ServiceDistTrafficMetricsAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service dist traffic metrics anon based on context it is used
func (m *ServiceDistTrafficMetricsAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDistTrafficMetricsAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDistTrafficMetricsAnon) UnmarshalBinary(b []byte) error {
	var res ServiceDistTrafficMetricsAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
