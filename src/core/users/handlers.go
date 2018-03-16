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
		resp.Write([]byte(unmarshallingError.Error()))
	}
	user := logInUser(userJson)
	if user != nil {
		resp.Write([]byte("Logged in!"))
	} else {
		resp.Write([]byte("There is no user with this status!"))
		resp.WriteHeader(http.StatusForbidden)
	}
}

func RegisterHandler(resp http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	registerJson := RegisterRequest{}
	decoder := json.NewDecoder(request.Body)
	unmarshallingError := decoder.Decode(&registerJson)
	if unmarshallingError != nil {
		resp.Write([]byte(unmarshallingError.Error()))
	}
	registerError := registerUser(registerJson)
	if registerError != nil {
		resp.Write([]byte(registerError.Error()))
	} else {
		resp.Write([]byte("New user registered!!!"))
	}
}
