// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FwDropsMetrics fw drops metrics
//
// swagger:model FwDropsMetrics
type FwDropsMetrics struct {

	// total fw drops
	TotalFwDrops float64 `json:"total_fw_drops,omitempty"`

	// total fw drops per rule
	TotalFwDropsPerRule []*FwDropsMetricsTotalFwDropsPerRuleItems0 `json:"total_fw_drops_per_rule"`
}

// Validate validates this fw drops metrics
func (m *FwDropsMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTotalFwDropsPerRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwDropsMetrics) validateTotalFwDropsPerRule(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalFwDropsPerRule) { // not required
		return nil
	}

	for i := 0; i < len(m.TotalFwDropsPerRule); i++ {
		if swag.IsZero(m.TotalFwDropsPerRule[i]) { // not required
			continue
		}

		if m.TotalFwDropsPerRule[i] != nil {
			if err := m.TotalFwDropsPerRule[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("total_fw_drops_per_rule" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("total_fw_drops_per_rule" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this fw drops metrics based on the context it is used
func (m *FwDropsMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTotalFwDropsPerRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwDropsMetrics) contextValidateTotalFwDropsPerRule(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TotalFwDropsPerRule); i++ {

		if m.TotalFwDropsPerRule[i] != nil {
			if err := m.TotalFwDropsPerRule[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("total_fw_drops_per_rule" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("total_fw_drops_per_rule" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FwDropsMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwDropsMetrics) UnmarshalBinary(b []byte) error {
	var res FwDropsMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FwDropsMetricsTotalFwDropsPerRuleItems0 fw drops metrics total fw drops per rule items0
//
// swagger:model FwDropsMetricsTotalFwDropsPerRuleItems0
type FwDropsMetricsTotalFwDropsPerRuleItems0 struct {

	// fw rule
	FwRule string `json:"fw_rule,omitempty"`

	// value
	Value float64 `json:"value,omitempty"`
}

// Validate validates this fw drops metrics total fw drops per rule items0
func (m *FwDropsMetricsTotalFwDropsPerRuleItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fw drops metrics total fw drops per rule items0 based on context it is used
func (m *FwDropsMetricsTotalFwDropsPerRuleItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FwDropsMetricsTotalFwDropsPerRuleItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwDropsMetricsTotalFwDropsPerRuleItems0) UnmarshalBinary(b []byte) error {
	var res FwDropsMetricsTotalFwDropsPerRuleItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
