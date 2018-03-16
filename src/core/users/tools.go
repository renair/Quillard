package users

import (
	"crypto/sha1"
	"fmt"
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
