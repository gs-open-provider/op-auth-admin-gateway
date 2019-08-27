package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gs-open-provider/op-auth-admin-gateway/internal/logger"
	"github.com/gs-open-provider/op-auth-admin-gateway/models"
	"github.com/gs-open-provider/op-auth-admin-gateway/restapi/operations/auth"
)

func getErrorResponse(resp *http.Response) middleware.Responder {
	logger.Log.Info("Generating error response for response with code: " + strconv.Itoa(resp.StatusCode))

	r := models.GeneralResponse{}
	json.NewDecoder(resp.Body).Decode(&r)

	switch resp.StatusCode {
	case 400:
		return auth.NewRegisterUserBadRequest().WithPayload(&r)
	case 401:
		return auth.NewRegisterUserUnauthorized().WithPayload(&r)
	case 403:
		return auth.NewRegisterUserForbidden().WithPayload(&r)
	case 404:
		return auth.NewRegisterUserNotFound().WithPayload(&r)
	default:
		return auth.NewRegisterUserInternalServerError().WithPayload(&r)
	}
}

func getGeneralResponseOK(resp *http.Response) middleware.Responder {
	logger.Log.Info("Generating general response for response with code: " + strconv.Itoa(resp.StatusCode))

	r := models.GeneralResponse{}
	json.NewDecoder(resp.Body).Decode(&r)

	return auth.NewRegisterUserOK().WithPayload(&r)
}
