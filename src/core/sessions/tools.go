package sessions

import (
	"crypto/sha1"
	"fmt"
	"strings"
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

func getNewKey(keywords []string) string {
	var in strings.Builder
	for _, word := range keywords {
		in.WriteString(":" + word)
	}
	inStr := fmt.Sprintf("%s:%v", in.String(), getCurrentUTC())
	return fmt.Sprintf("%x", sha1.Sum([]byte(inStr)))
}

func sessionClearer(durr time.Duration) {
	for {
		currTime := getCurrentUTC()
		//use anonymous function here to use calculate current UTC time once
		activeSessions.Range(func(k, v interface{}) bool {
			session := v.(Session)
			if session.Expires <= currTime {
				activeSessions.Delete(k)
			}
			return true
		})
		time.Sleep(durr)
	}
}
