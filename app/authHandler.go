package app

import (
	"capi/dto"
	"capi/logger"
	"capi/service"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var LoginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&LoginRequest); err != nil {
		logger.Error("error while decode login request: " + err.Error())
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}

	response, appErr := h.service.Login(LoginRequest)

	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, response)
}
