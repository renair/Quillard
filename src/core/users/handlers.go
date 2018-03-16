package users

import (
	"encoding/json"
	"net/http"
)

func LoginHandler(resp http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	userJson := LoginRequest{}
	decoder := json.NewDecoder(request.Body)
	unmarshallingError := decoder.Decode(&userJson)
	if unmarshallingError != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(unmarshallingError.Error()))
		return
	}
	user := logInUser(userJson)
	if user != nil {
		resp.Write([]byte("Logged in!"))
	} else {
		resp.WriteHeader(http.StatusForbidden)
		resp.Write([]byte("There is no user with this status!"))
	}
}

func RegisterHandler(resp http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	registerJson := RegisterRequest{}
	decoder := json.NewDecoder(request.Body)
	unmarshallingError := decoder.Decode(&registerJson)
	if unmarshallingError != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(unmarshallingError.Error()))
		return
	}
	registerError := registerUser(registerJson)
	if registerError != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(registerError.Error()))
	} else {
		resp.Write([]byte("New user registered!!!"))
	}
}
