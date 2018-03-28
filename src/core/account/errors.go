package account

import (
	"fmt"
	"net/http"
	"qutils/basehandlers"
	"qutils/coder"
)

//Error handlers
//Use ONLY as enclosed handlers becuse this ones don't close Body

//Respond with JSON with error #1 'This account doesn't exists or password incorrect'
func loginIncorrect(resp http.ResponseWriter, request *http.Request) {
	err := basehandlers.StatusedResponse{
		Code:    3,
		Message: "This account doesn't exists or password incorrect",
	}
	fmt.Fprint(resp, coder.EncodeJson(err))
}

//Respond with JSON with error #2 'This email already registered'
func emailAlreadyExists(resp http.ResponseWriter, req *http.Request) {
	err := basehandlers.StatusedResponse{
		Code:    4,
		Message: "This email already registered",
	}
	fmt.Fprint(resp, coder.EncodeJson(err))
}

//Respond with JSON with error #3 'This nickname already exists'
func nearbyBuildingsExist(resp http.ResponseWriter, req *http.Request) {
	err := basehandlers.StatusedResponse{
		Code:    5,
		Message: "There is some structures near you",
	}
	fmt.Fprint(resp, coder.EncodeJson(err))
}
