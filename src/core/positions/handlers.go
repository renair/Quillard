package positions

import (
	"core/sessions"
	"fmt"
	"net/http"
	"qutils/basehandlers"
	"qutils/coder"
)

func AccountHomeHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	session, ok := sessions.GetSessionByRequest(req)
	if !ok {
		basehandlers.UnauthorizedRequest(resp, req)
		return
	}
	position := GetAccountHomePosition(session.AccountId)
	fmt.Fprint(resp, coder.EncodeJson(position))
}

func PersonageGetNearestHomes(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	session, ok := sessions.GetSessionByRequest(req)
	if !ok {
		basehandlers.UnauthorizedRequest(resp, req)
		return
	}
	persId := PersonageId{}
	coder.DecodeJson(req.Body, &persId)
	positions, err := getNearestHomes(persId.Id, session.AccountId, VIEWDISTANCE)
	if err != nil {
		fmt.Println(err.Error())
		basehandlers.InternalError(resp, req)
		return
	}
	fmt.Fprint(resp, coder.EncodeJson(positions))
}
