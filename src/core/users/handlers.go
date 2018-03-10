package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LoginHandler(resp http.ResponseWriter, request *http.Request) {
	fmt.Println("Here is handler!!!")
	userJson := LoginRequest{}
	defer request.Body.Close()
	bytes, bodyReadErr := ioutil.ReadAll(request.Body)
	if bodyReadErr != nil {
		resp.Write([]byte(bodyReadErr.Error()))
		return
	}
	unmarshallingError := json.Unmarshal(bytes, &userJson)
	if unmarshallingError != nil {
		resp.Write([]byte(unmarshallingError.Error()))
	}
	user := logInUser(userJson.Email, userJson.Password)
	if user != nil {
		resp.Write([]byte("Logged in!"))
	} else {
		resp.Write([]byte("There is no user with this status!"))
		resp.WriteHeader(http.StatusForbidden)
	}
}
