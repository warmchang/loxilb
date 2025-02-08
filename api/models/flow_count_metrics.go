// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FlowCountMetrics flow count metrics
//
// swagger:model FlowCountMetrics
type FlowCountMetrics struct {

	// active conntrack count
	ActiveConntrackCount float64 `json:"active_conntrack_count,omitempty"`

	// active flow count sctp
	ActiveFlowCountSctp float64 `json:"active_flow_count_sctp,omitempty"`

	// active flow count tcp
	ActiveFlowCountTCP float64 `json:"active_flow_count_tcp,omitempty"`

	// active flow count udp
	ActiveFlowCountUDP float64 `json:"active_flow_count_udp,omitempty"`

	// inactive flow count
	InactiveFlowCount float64 `json:"inactive_flow_count,omitempty"`
}

// Validate validates this flow count metrics
func (m *FlowCountMetrics) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this flow count metrics based on context it is used
func (m *FlowCountMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlowCountMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowCountMetrics) UnmarshalBinary(b []byte) error {
	var res FlowCountMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
