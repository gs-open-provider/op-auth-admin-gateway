// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// LoginCallbackHandlerFunc turns a function with the right signature into a login callback handler
type LoginCallbackHandlerFunc func(LoginCallbackParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LoginCallbackHandlerFunc) Handle(params LoginCallbackParams) middleware.Responder {
	return fn(params)
}

// LoginCallbackHandler interface for that can handle valid login callback params
type LoginCallbackHandler interface {
	Handle(LoginCallbackParams) middleware.Responder
}

// NewLoginCallback creates a new http.Handler for the login callback operation
func NewLoginCallback(ctx *middleware.Context, handler LoginCallbackHandler) *LoginCallback {
	return &LoginCallback{Context: ctx, Handler: handler}
}

/*LoginCallback swagger:route GET /v1/oauth2/oauth-signin-cb auth loginCallback

LoginCallback login callback API

*/
type LoginCallback struct {
	Context *middleware.Context
	Handler LoginCallbackHandler
}

func (o *LoginCallback) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLoginCallbackParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}