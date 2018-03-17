package sessions

import (
	"core/users"
	"crypto/sha1"
	"fmt"
	"time"
)

func checkConnection() {
	if connection == nil {
		panic("Sessions module critical! Connection to database is not set!")
	}
}

func getCurrentUTC() int64 {
	return time.Now().UTC().Unix()
}

func getNewKey(user *users.User) string {
	in := fmt.Sprintf("%d:%s:%d", user.Id, user.Nickname, getCurrentUTC())
	return fmt.Sprintf("%x", sha1.Sum([]byte(in)))
}
