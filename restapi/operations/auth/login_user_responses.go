// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/gs-open-provider/op-auth-admin-gateway/models"
)

// LoginUserOKCode is the HTTP code returned for type LoginUserOK
const LoginUserOKCode int = 200

/*LoginUserOK OK

swagger:response loginUserOK
*/
type LoginUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.LoginResponse `json:"body,omitempty"`
}

// NewLoginUserOK creates LoginUserOK with default headers values
func NewLoginUserOK() *LoginUserOK {

	return &LoginUserOK{}
}

// WithPayload adds the payload to the login user o k response
func (o *LoginUserOK) WithPayload(payload *models.LoginResponse) *LoginUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user o k response
func (o *LoginUserOK) SetPayload(payload *models.LoginResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUserBadRequestCode is the HTTP code returned for type LoginUserBadRequest
const LoginUserBadRequestCode int = 400

/*LoginUserBadRequest BAD REQUEST

swagger:response loginUserBadRequest
*/
type LoginUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralResponse `json:"body,omitempty"`
}

// NewLoginUserBadRequest creates LoginUserBadRequest with default headers values
func NewLoginUserBadRequest() *LoginUserBadRequest {

	return &LoginUserBadRequest{}
}

// WithPayload adds the payload to the login user bad request response
func (o *LoginUserBadRequest) WithPayload(payload *models.GeneralResponse) *LoginUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user bad request response
func (o *LoginUserBadRequest) SetPayload(payload *models.GeneralResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUserUnauthorizedCode is the HTTP code returned for type LoginUserUnauthorized
const LoginUserUnauthorizedCode int = 401

/*LoginUserUnauthorized UNAUTHORIZED

swagger:response loginUserUnauthorized
*/
type LoginUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralResponse `json:"body,omitempty"`
}

// NewLoginUserUnauthorized creates LoginUserUnauthorized with default headers values
func NewLoginUserUnauthorized() *LoginUserUnauthorized {

	return &LoginUserUnauthorized{}
}

// WithPayload adds the payload to the login user unauthorized response
func (o *LoginUserUnauthorized) WithPayload(payload *models.GeneralResponse) *LoginUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user unauthorized response
func (o *LoginUserUnauthorized) SetPayload(payload *models.GeneralResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUserForbiddenCode is the HTTP code returned for type LoginUserForbidden
const LoginUserForbiddenCode int = 403

/*LoginUserForbidden FORBIDDEN

swagger:response loginUserForbidden
*/
type LoginUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralResponse `json:"body,omitempty"`
}

// NewLoginUserForbidden creates LoginUserForbidden with default headers values
func NewLoginUserForbidden() *LoginUserForbidden {

	return &LoginUserForbidden{}
}

// WithPayload adds the payload to the login user forbidden response
func (o *LoginUserForbidden) WithPayload(payload *models.GeneralResponse) *LoginUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user forbidden response
func (o *LoginUserForbidden) SetPayload(payload *models.GeneralResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUserNotFoundCode is the HTTP code returned for type LoginUserNotFound
const LoginUserNotFoundCode int = 404

/*LoginUserNotFound NOT FOUND

swagger:response loginUserNotFound
*/
type LoginUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralResponse `json:"body,omitempty"`
}

// NewLoginUserNotFound creates LoginUserNotFound with default headers values
func NewLoginUserNotFound() *LoginUserNotFound {

	return &LoginUserNotFound{}
}

// WithPayload adds the payload to the login user not found response
func (o *LoginUserNotFound) WithPayload(payload *models.GeneralResponse) *LoginUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user not found response
func (o *LoginUserNotFound) SetPayload(payload *models.GeneralResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUserInternalServerErrorCode is the HTTP code returned for type LoginUserInternalServerError
const LoginUserInternalServerErrorCode int = 500

/*LoginUserInternalServerError INTERNAL SERVER ERROR

swagger:response loginUserInternalServerError
*/
type LoginUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralResponse `json:"body,omitempty"`
}

// NewLoginUserInternalServerError creates LoginUserInternalServerError with default headers values
func NewLoginUserInternalServerError() *LoginUserInternalServerError {

	return &LoginUserInternalServerError{}
}

// WithPayload adds the payload to the login user internal server error response
func (o *LoginUserInternalServerError) WithPayload(payload *models.GeneralResponse) *LoginUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user internal server error response
func (o *LoginUserInternalServerError) SetPayload(payload *models.GeneralResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
