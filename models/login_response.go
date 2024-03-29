// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LoginResponse login response
// swagger:model loginResponse
type LoginResponse struct {

	// data
	Data *LoginResponseData `json:"data,omitempty"`

	// error
	Error *LoginResponseError `json:"error,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this login response
func (m *LoginResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *LoginResponse) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoginResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginResponse) UnmarshalBinary(b []byte) error {
	var res LoginResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoginResponseData login response data
// swagger:model LoginResponseData
type LoginResponseData struct {

	// access token
	AccessToken string `json:"accessToken,omitempty"`

	// customizations
	Customizations *LoginResponseDataCustomizations `json:"customizations,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// expires in
	ExpiresIn string `json:"expiresIn,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// user ref
	UserRef string `json:"userRef,omitempty"`

	// user type
	UserType string `json:"userType,omitempty"`
}

// Validate validates this login response data
func (m *LoginResponseData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomizations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginResponseData) validateCustomizations(formats strfmt.Registry) error {

	if swag.IsZero(m.Customizations) { // not required
		return nil
	}

	if m.Customizations != nil {
		if err := m.Customizations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "customizations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoginResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginResponseData) UnmarshalBinary(b []byte) error {
	var res LoginResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoginResponseDataCustomizations login response data customizations
// swagger:model LoginResponseDataCustomizations
type LoginResponseDataCustomizations struct {

	// background color
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// background image URL
	BackgroundImageURL string `json:"backgroundImageURL,omitempty"`

	// header1 font
	Header1Font string `json:"header1Font,omitempty"`

	// header1 font color
	Header1FontColor string `json:"header1FontColor,omitempty"`

	// header1 font size
	Header1FontSize string `json:"header1FontSize,omitempty"`

	// header2 font
	Header2Font string `json:"header2Font,omitempty"`

	// header2 font color
	Header2FontColor string `json:"header2FontColor,omitempty"`

	// header2 font size
	Header2FontSize string `json:"header2FontSize,omitempty"`

	// logo URL
	LogoURL string `json:"logoURL,omitempty"`

	// navbar color
	NavbarColor string `json:"navbarColor,omitempty"`

	// navbar font
	NavbarFont string `json:"navbarFont,omitempty"`

	// navbar font color
	NavbarFontColor string `json:"navbarFontColor,omitempty"`

	// navbar font size
	NavbarFontSize string `json:"navbarFontSize,omitempty"`

	// navbar select color
	NavbarSelectColor string `json:"navbarSelectColor,omitempty"`

	// text font
	TextFont string `json:"textFont,omitempty"`

	// text font color
	TextFontColor string `json:"textFontColor,omitempty"`

	// text font size
	TextFontSize string `json:"textFontSize,omitempty"`
}

// Validate validates this login response data customizations
func (m *LoginResponseDataCustomizations) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoginResponseDataCustomizations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginResponseDataCustomizations) UnmarshalBinary(b []byte) error {
	var res LoginResponseDataCustomizations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoginResponseError login response error
// swagger:model LoginResponseError
type LoginResponseError struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this login response error
func (m *LoginResponseError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoginResponseError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginResponseError) UnmarshalBinary(b []byte) error {
	var res LoginResponseError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
