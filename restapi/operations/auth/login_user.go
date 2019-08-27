// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// LoginUserHandlerFunc turns a function with the right signature into a login user handler
type LoginUserHandlerFunc func(LoginUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LoginUserHandlerFunc) Handle(params LoginUserParams) middleware.Responder {
	return fn(params)
}

// LoginUserHandler interface for that can handle valid login user params
type LoginUserHandler interface {
	Handle(LoginUserParams) middleware.Responder
}

// NewLoginUser creates a new http.Handler for the login user operation
func NewLoginUser(ctx *middleware.Context, handler LoginUserHandler) *LoginUser {
	return &LoginUser{Context: ctx, Handler: handler}
}

/*LoginUser swagger:route POST /v1/auth/signin auth loginUser

LoginUser login user API

*/
type LoginUser struct {
	Context *middleware.Context
	Handler LoginUserHandler
}

func (o *LoginUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLoginUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
