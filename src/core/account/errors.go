package account

import (
	"fmt"
	"net/http"
)

//Error structure
type UserError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//Error handlers
//Use ONLY as enclosed handlers becuse this ones don't close Body

func internalError(resp http.ResponseWriter, request *http.Request) {
	err := UserError{
		Code:    0,
		Message: "Some internal server error occured",
	}
	resp.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(resp, encodeJson(err))
}

func jsonUnmarshallingError(resp http.ResponseWriter, request *http.Request) {
	err := UserError{
		Code:    0,
		Message: "Bad formed JSON",
	}
	resp.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(resp, encodeJson(err))
}

//Respond with JSON with error #1 'This account doesn't exists or password incorrect'
func loginIncorrect(resp http.ResponseWriter, request *http.Request) {
	err := UserError{
		Code:    1,
		Message: "This account doesn't exists or password incorrect",
	}
	fmt.Fprint(resp, encodeJson(err))
}

//Respond with JSON with error #2 'This email already registered'
func emailAlreadyExists(resp http.ResponseWriter, req *http.Request) {
	err := UserError{
		Code:    2,
		Message: "This email already registered",
	}
	fmt.Fprint(resp, encodeJson(err))
}

//Respond with JSON with error #3 'This nickname already exists'
func nearbyBuildingsExist(resp http.ResponseWriter, req *http.Request) {
	err := UserError{
		Code:    3,
		Message: "There is some structures near you",
	}
	fmt.Fprint(resp, encodeJson(err))
}
