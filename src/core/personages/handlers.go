package personages

import (
	"core/sessions"
	"fmt"
	"net/http"
	"qutils/basehandlers"
	"qutils/coder"
)

func CreatePersonageHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	session, ok := sessions.GetSessionByRequest(req)
	if !ok {
		basehandlers.UnauthorizedRequest(resp, req)
		return
	}
	data := PersonageRequest{}
	decodingErr := coder.DecodeJson(req.Body, &data)
	if decodingErr != nil {
		basehandlers.JsonUnmarshallingError(resp, req)
	}
	registrationError := registerPersonage(session, data)
	if registrationError != nil {
		//it's mean that some error occured during inserting
		basehandlers.InternalError(resp, req)
		return
	}
	basehandlers.SuccessResponse(resp, req)
}

func GetOwnPersonagesHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	session, ok := sessions.GetSessionByRequest(req)
	if !ok {
		basehandlers.UnauthorizedRequest(resp, req)
		return
	}
	personagesList := getAccountPersonages(session)
	fmt.Fprint(resp, coder.EncodeJson(personagesList))
}
