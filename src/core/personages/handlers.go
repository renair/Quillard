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
	sessionKey := req.Header.Get("Q-Session")
	session := sessions.GetSession(sessionKey)
	if session == nil {
		//handle error "Session expired"
		basehandlers.InternalError(resp, req)
		return
	}
	data := PersonageRequest{}
	decodingErr := coder.DecodeJson(req.Body, &data)
	if decodingErr != nil {
		basehandlers.JsonUnmarshallingError(resp, req)
	}
	registrationError := registerPersonage(session, data)
	if registrationError != nil {
		//handle registration error
		basehandlers.InternalError(resp, req)
		return
	}
	basehandlers.SuccessResponse(resp, req)
}

func GetOwnPersonagesHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	sessionKey := req.Header.Get("Q-Session")
	session := sessions.GetSession(sessionKey)
	if session == nil {
		//handle error "Session expired"
	}
	personagesList := getAccountPersonages(session)
	fmt.Fprint(resp, coder.EncodeJson(personagesList))
}
