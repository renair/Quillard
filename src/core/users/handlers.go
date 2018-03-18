package users

import (
	"core/sessions"
	"net/http"
)

func LoginHandler(resp http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	userJson := LoginRequest{}
	unmarshallingError := decodeJson(request.Body, &userJson)
	if unmarshallingError != nil {
		jsonUnmurshallingError(resp, request)
		return
	}
	user := logInUser(userJson)
	if user != nil {
		session := sessions.CreateSession(user.Id, user.Email, user.Nickname)
		if session == nil {
			internalError(resp, request)
			return
		}
		resp.Write([]byte(encodeJson(session.ToResponse())))
	} else {
		loginIncorrect(resp, request)
	}
}

func RegisterHandler(resp http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	registerJson := RegisterRequest{}
	unmarshallingError := decodeJson(request.Body, &registerJson)
	if unmarshallingError != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(unmarshallingError.Error()))
		return
	}
	user, registerError := registerUser(registerJson)
	if registerError != nil {
		switch registerError.Error() {
		case "email":
			emailAlreadyExists(resp, request)
		case "nickname":
			nicknameAlreadyExists(resp, request)
		default:
			internalError(resp, request)
		}
	} else {
		session := sessions.CreateSession(user.Id, user.Email, user.Nickname)
		if session == nil {
			internalError(resp, request)
			return
		}
		resp.Write([]byte(encodeJson(session.ToResponse())))
	}
}
