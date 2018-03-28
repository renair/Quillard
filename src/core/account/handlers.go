package account

import (
	"core/sessions"
	"net/http"
	"qutils/coder"
)

func LoginHandler(resp http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	userJson := LoginRequest{}
	unmarshallingError := coder.DecodeJson(request.Body, &userJson)
	if unmarshallingError != nil {
		jsonUnmarshallingError(resp, request)
		return
	}
	user := logInAccount(userJson)
	if user != nil {
		session := sessions.CreateSession(user.Id, user.Email, user.Password)
		if session == nil {
			internalError(resp, request)
			return
		}
		resp.Write([]byte(coder.EncodeJson(session.ToResponse())))
	} else {
		loginIncorrect(resp, request)
	}
}

func RegisterHandler(resp http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	registerJson := RegisterRequest{}
	unmarshallingError := coder.DecodeJson(request.Body, &registerJson)
	if unmarshallingError != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(unmarshallingError.Error()))
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
			internalError(resp, request)
		}
	} else {
		session := sessions.CreateSession(user.Id, user.Email, user.Password)
		if session == nil {
			internalError(resp, request)
			return
		}
		resp.Write([]byte(coder.EncodeJson(session.ToResponse())))
	}
}
