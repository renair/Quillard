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
