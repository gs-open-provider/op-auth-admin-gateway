// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RegisterDetailsRequest register details request
// swagger:model registerDetailsRequest
type RegisterDetailsRequest struct {

	// details
	Details string `json:"details,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`
}

// Validate validates this register details request
func (m *RegisterDetailsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegisterDetailsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterDetailsRequest) UnmarshalBinary(b []byte) error {
	var res RegisterDetailsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}