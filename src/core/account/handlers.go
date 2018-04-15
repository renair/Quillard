package account

import (
	"core/sessions"
	"fmt"
	"net/http"
	"qutils/basehandlers"
	"qutils/coder"
)

func LoginHandler(resp http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	userJson := LoginRequest{}
	unmarshallingError := coder.DecodeJson(request.Body, &userJson)
	if unmarshallingError != nil {
		basehandlers.JsonUnmarshallingError(resp, request)
		return
	}
	user := logInAccount(userJson)
	if user != nil {
		session, ok := sessions.CreateSession(user.Id, user.Email, user.Password)
		if !ok {
			basehandlers.InternalError(resp, request)
			return
		}
		fmt.Fprint(resp, coder.EncodeJson(session.ToResponse()))
	} else {
		loginIncorrect(resp, request)
	}
}

func RegisterHandler(resp http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	registerJson := RegisterRequest{}
	unmarshallingError := coder.DecodeJson(request.Body, &registerJson)
	if unmarshallingError != nil {
		basehandlers.JsonUnmarshallingError(resp, request)
		return
	}
	user, registerError := registerAccount(registerJson)
	if registerError != nil {
		switch registerError.Error() {
		case "email":
			emailAlreadyExists(resp, request)
		case "position":
			nearbyBuildingsExist(resp, request)
		default:
			basehandlers.InternalError(resp, request)
		}
	} else {
		session, ok := sessions.CreateSession(user.Id, user.Email, user.Password)
		if !ok {
			basehandlers.InternalError(resp, request)
			return
		}
		fmt.Fprint(resp, coder.EncodeJson(session.ToResponse()))
	}
}
