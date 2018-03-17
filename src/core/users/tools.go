package users

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

//Create panic if connection not initialized
func chekInitialization() {
	if connection == nil {
		panic("DBConnection is not initialized!!!")
	}
}

func isFieldExist(field string, nickname string) bool {
	chekInitialization()
	args := map[string]interface{}{
		field: nickname,
	}
	res, _ := connection.SelectBy(TABLENAME, args, "id")
	defer res.Close()
	return res.Next()
}

func encodePassword(pass string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(pass)))
}

func decodeJson(io io.Reader, v interface{}) error {
	decoder := json.NewDecoder(io)
	return decoder.Decode(v)
}

func encodeJson(v interface{}) string {
	var builder strings.Builder
	encoder := json.NewEncoder(&builder)
	encoder.Encode(v)
	return builder.String()
}
