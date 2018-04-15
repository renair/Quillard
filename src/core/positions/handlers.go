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
	sessionKey := req.Header.Get("Q-Session")
	session, ok := sessions.GetSession(sessionKey)
	if !ok {
		basehandlers.UnauthorizedRequest(resp, req)
		return
	}
	position := GetAccountHomePosition(session.AccountId)
	fmt.Fprint(resp, coder.EncodeJson(position))
}
