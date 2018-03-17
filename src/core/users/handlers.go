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
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(unmarshallingError.Error()))
		return
	}
	user := logInUser(userJson)
	if user != nil {
		session := sessions.CreateSession(user.Id, user.Email, user.Nickname)
		if session == nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte("Can't create session"))
			return
		}
		resp.Write([]byte(encodeJson(session.ToResponse())))
	} else {
		resp.WriteHeader(http.StatusForbidden)
		resp.Write([]byte("There is no user with this status!"))
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
	if registerError == nil && user != nil {
		session := sessions.CreateSession(user.Id, user.Email, user.Nickname)
		if session == nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte("Can't create session"))
			return
		}
		resp.Write([]byte(encodeJson(session.ToResponse())))
	} else {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(registerError.Error()))
	}
}
