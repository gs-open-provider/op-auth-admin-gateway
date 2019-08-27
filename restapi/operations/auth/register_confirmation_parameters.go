// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRegisterConfirmationParams creates a new RegisterConfirmationParams object
// no default values defined in spec.
func NewRegisterConfirmationParams() RegisterConfirmationParams {

	return RegisterConfirmationParams{}
}

// RegisterConfirmationParams contains all the bound params for the register confirmation operation
// typically these are obtained from a http.Request
//
// swagger:parameters RegisterConfirmation
type RegisterConfirmationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Token string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRegisterConfirmationParams() beforehand.
func (o *RegisterConfirmationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rToken, rhkToken, _ := route.Params.GetOK("token")
	if err := o.bindToken(rToken, rhkToken, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindToken binds and validates parameter Token from path.
func (o *RegisterConfirmationParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Token = raw

	return nil
}
