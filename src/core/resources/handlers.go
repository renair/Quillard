package resources

import (
	"core/sessions"
	"fmt"
	"net/http"
	"qutils/basehandlers"
	"qutils/coder"
)

func PersonageResourcesHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	session, ok := sessions.GetSessionByRequest(req)
	if !ok {
		basehandlers.UnauthorizedRequest(resp, req)
		return
	}
	reqData := ResourceRequest{}
	decodingError := coder.DecodeJson(req.Body, &reqData)
	if decodingError != nil {
		basehandlers.JsonUnmarshallingError(resp, req)
	}
	personageResources := GetPersonageResources(reqData.PersonageId, session.AccountId)
	if len(personageResources) == 0 {
		InitResourcesForPersonage(reqData.PersonageId)
		personageResources = GetPersonageResources(reqData.PersonageId, session.AccountId)
	}
	encodedAnswer := coder.EncodeJson(personageResources)
	fmt.Fprint(resp, encodedAnswer)
}
