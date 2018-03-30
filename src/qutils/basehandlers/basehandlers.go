package basehandlers

import (
	"fmt"
	"net/http"
	"qutils/coder"
)

func SuccessResponse(resp http.ResponseWriter, request *http.Request) {
	response := StatusedResponse{
		Code:    0,
		Message: "Operation successed",
	}
	fmt.Fprint(resp, coder.EncodeJson(response))
}

func InternalError(resp http.ResponseWriter, request *http.Request) {
	err := StatusedResponse{
		Code:    1,
		Message: "Some internal server error occured",
	}
	resp.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(resp, coder.EncodeJson(err))
}

func JsonUnmarshallingError(resp http.ResponseWriter, request *http.Request) {
	err := StatusedResponse{
		Code:    2,
		Message: "Bad formed JSON",
	}
	resp.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(resp, coder.EncodeJson(err))
}
