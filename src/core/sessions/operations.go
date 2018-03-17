package sessions

import (
	"core/users"
	"fmt"
)

func CreateSession(user *users.User) *Session {
	checkConnection()
	utc := getCurrentUTC()
	keys := map[string]interface{}{
		"user_id": user.Id,
		"key":     getNewKey(user),
		"created": utc,
		"expires": utc + EXPIRETIME,
	}
	insertionErr := connection.Insert(TABLENAME, keys)
	if insertionErr != nil {
		return nil
	}
	return &Session{
		Id:      -1,
		UserId:  keys["user_id"].(int64),
		Key:     keys["key"].(string),
		Created: keys["created"].(int64),
		Expires: keys["expires"].(int64),
	}
}

func GetSession(key string) *Session {
	checkConnection()
	var ses *Session = nil
	query := fmt.Sprintf("SELECT id, user_id, key, created, expires FROM sessions WHERE key='%s' AND expires>%d", key, getCurrentUTC())
	rows, _ := connection.ManualQuery(query)
	if rows.Next() {
		ses = new(Session)
		rows.Scan(&ses.Id, &ses.UserId, &ses.Key, &ses.Created, &ses.Expires)
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
