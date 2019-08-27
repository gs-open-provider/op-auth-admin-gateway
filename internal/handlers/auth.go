package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gs-open-provider/op-auth-admin-gateway/internal/logger"
	"github.com/gs-open-provider/op-auth-admin-gateway/restapi/operations/auth"
)

// RegisterUser Function
func RegisterUser(params *auth.RegisterUserParams) middleware.Responder {
	logger.Log.Info("RegisterUser called..")

	resp, err := handlePostBasicRequest("http://localhost:9090/v1/auth/signup", params.RegisterRequest)
	if err != nil {
		logger.Log.Error(err.Error())
	}
	logger.Log.Info(resp.Status)

	if resp.StatusCode == 200 {
		return getGeneralResponseOK(resp)
	}
	return getErrorResponse(resp)
}
