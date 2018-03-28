package sessions

import (
	"fmt"
)

func CreateSession(id int64, keywords ...string) *Session {
	checkConnection()
	utc := getCurrentUTC()
	keys := map[string]interface{}{
		"user_id": id,
		"key":     getNewKey(keywords),
		"created": utc,
		"expires": utc + EXPIRETIME,
	}
	connection.BeginTransaction()
	insertionErr := connection.Insert(TABLENAME, keys)
	if insertionErr != nil {
		connection.RollbackTransaction()
		return nil
	}
	connection.CommitTransaction()
	return &Session{
		Id:      -1,
		AccountId:  id,
		Key:     keys["key"].(string),
		Created: utc,
		Expires: utc + EXPIRETIME,
	}
}

func GetSession(key string) *Session {
	checkConnection()
	var ses *Session = nil
	query := fmt.Sprintf("SELECT id, user_id, key, created, expires FROM sessions WHERE key='%s' AND expires>%d", key, getCurrentUTC())
	rows, _ := connection.ManualQuery(query)
	if rows.Next() {
		ses = new(Session)
		rows.Scan(&ses.Id, &ses.AccountId, &ses.Key, &ses.Created, &ses.Expires)
	}
	return ses
}

//Operation from session
func (ses *Session) ToResponse() SessionResponse {
	return SessionResponse{
		Key:     ses.Key,
		Expires: ses.Expires,
	}
}
