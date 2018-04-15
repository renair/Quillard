package sessions

import (
	"net/http"
)

func CreateSession(id int64, keywords ...string) (Session, bool) {
	checkConnection()
	utc := getCurrentUTC()
	keys := map[string]interface{}{
		"user_id": id,
		"key":     getNewKey(keywords),
		"created": utc,
		"expires": utc + EXPIRETIME,
	}
	insertionErr := connection.Insert(TABLENAME, keys)
	if insertionErr != nil {
		return Session{}, false
	}
	session := Session{
		Id:        -1,
		AccountId: id,
		Key:       keys["key"].(string),
		Created:   utc,
		Expires:   utc + EXPIRETIME,
	}
	activeSessions.Store(session.Key, &session)
	return session, true
}

func GetSessionByKey(key string) (Session, bool) {
	//checkConnection()
	//var ses *Session = nil
	//query := fmt.Sprintf("SELECT id, user_id, key, created, expires FROM sessions WHERE key='%s' AND expires>%d", key, getCurrentUTC())
	//rows, _ := connection.ManualQuery(query)
	//	if rows.Next() {
	//		ses = new(Session)
	//		rows.Scan(&ses.Id, &ses.AccountId, &ses.Key, &ses.Created, &ses.Expires)
	//	}
	//return ses
	sesInterface, ok := activeSessions.Load(key)
	if !ok {
		return Session{}, false
	}
	session := sesInterface.(*Session)
	//renewing session
	session.Expires = getCurrentUTC() + EXPIRETIME
	return *session, true
}

func GetSessionByRequest(req *http.Request) (Session, bool) {
	sessionKey := req.Header.Get("Q-Session")
	if sessionKey == "" {
		return Session{}, false
	}
	return GetSessionByKey(sessionKey)
}

//Operation from session
func (ses *Session) ToResponse() SessionResponse {
	return SessionResponse{
		Key:     ses.Key,
		Expires: ses.Expires,
	}
}
