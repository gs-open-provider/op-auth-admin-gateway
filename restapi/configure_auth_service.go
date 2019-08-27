// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"github.com/gs-open-provider/op-auth-admin-gateway/internal/config"
	"github.com/gs-open-provider/op-auth-admin-gateway/internal/handlers"
	"github.com/gs-open-provider/op-auth-admin-gateway/internal/logger"
	"github.com/gs-open-provider/op-auth-admin-gateway/restapi/operations"
	"github.com/gs-open-provider/op-auth-admin-gateway/restapi/operations/auth"
)

//go:generate swagger generate server --target ../../op-auth-admin-gateway --name AuthService --spec ../spec/root.yml

func configureFlags(api *operations.AuthServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AuthServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	// Manual Configurations and setup here..
	config.InitializeViper()
	logger.InitializeZapCustomLogger()
	// Manual Configurations and setup here..

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.GetHandler == nil {
		api.GetHandler = operations.GetHandlerFunc(func(params operations.GetParams) middleware.Responder {
			return middleware.NotImplemented("operation .Get has not yet been implemented")
		})
	}
	if api.AuthLoginCallbackHandler == nil {
		api.AuthLoginCallbackHandler = auth.LoginCallbackHandlerFunc(func(params auth.LoginCallbackParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.LoginCallback has not yet been implemented")
		})
	}
	if api.AuthLoginUserHandler == nil {
		api.AuthLoginUserHandler = auth.LoginUserHandlerFunc(func(params auth.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.LoginUser has not yet been implemented")
		})
	}
	if api.AuthRefreshTokenHandler == nil {
		api.AuthRefreshTokenHandler = auth.RefreshTokenHandlerFunc(func(params auth.RefreshTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.RefreshToken has not yet been implemented")
		})
	}
	if api.AuthRegisterCalbackHandler == nil {
		api.AuthRegisterCalbackHandler = auth.RegisterCalbackHandlerFunc(func(params auth.RegisterCalbackParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.RegisterCalback has not yet been implemented")
		})
	}
	if api.AuthRegisterConfirmationHandler == nil {
		api.AuthRegisterConfirmationHandler = auth.RegisterConfirmationHandlerFunc(func(params auth.RegisterConfirmationParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.RegisterConfirmation has not yet been implemented")
		})
	}
	if api.AuthRegisterUserHandler == nil {
		api.AuthRegisterUserHandler = auth.RegisterUserHandlerFunc(func(params auth.RegisterUserParams) middleware.Responder {
			return handlers.RegisterUser(&params)
		})
	}
	if api.AuthRegisterUserDetailsHandler == nil {
		api.AuthRegisterUserDetailsHandler = auth.RegisterUserDetailsHandlerFunc(func(params auth.RegisterUserDetailsParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.RegisterUserDetails has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.AllowAll().Handler
	return handleCORS(handler)
}
