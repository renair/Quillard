package personages

import (
	"core/sessions"
	"fmt"
	"net/http"
	"qutils/coder"
)

func CreatePersonageHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	sessionKey := req.Header.Get("Q-Session")
	session := sessions.GetSession(sessionKey)
	if session == nil {
		//handle error "Session expired"
	}
	data := PersonageRequest{}
	coder.DecodeJson(req.Body, &data)
	registrationError := registerPersonage(session, data)
	if registrationError != nil {
		//handle registration error
	}
	fmt.Fprint(resp, "OK") //TODO handle with default OK handler
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
