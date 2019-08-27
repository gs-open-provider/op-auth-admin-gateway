// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RegisterUserDetailsHandlerFunc turns a function with the right signature into a register user details handler
type RegisterUserDetailsHandlerFunc func(RegisterUserDetailsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterUserDetailsHandlerFunc) Handle(params RegisterUserDetailsParams) middleware.Responder {
	return fn(params)
}

// RegisterUserDetailsHandler interface for that can handle valid register user details params
type RegisterUserDetailsHandler interface {
	Handle(RegisterUserDetailsParams) middleware.Responder
}

// NewRegisterUserDetails creates a new http.Handler for the register user details operation
func NewRegisterUserDetails(ctx *middleware.Context, handler RegisterUserDetailsHandler) *RegisterUserDetails {
	return &RegisterUserDetails{Context: ctx, Handler: handler}
}

/*RegisterUserDetails swagger:route POST /v1/auth/user-details auth registerUserDetails

RegisterUserDetails register user details API

*/
type RegisterUserDetails struct {
	Context *middleware.Context
	Handler RegisterUserDetailsHandler
}

func (o *RegisterUserDetails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRegisterUserDetailsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}