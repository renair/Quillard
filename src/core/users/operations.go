package users

import (
	//	"dbwrapper"
	"crypto/sha1"
	"fmt"
)

//Ruturn user object based on login and password
func logInUser(email string, password string) *User {
	chekInitialization()
	keys := make(map[string]interface{})
	keys["email"] = email
	keys["password"] = fmt.Sprintf("%x", sha1.Sum([]byte(password)))
	res, err := connection.SelectBy(TABLENAME, keys, "id", "nickname", "password")
	defer res.Close()
	var user *User = nil
	for err == nil && res.Next() {
		user = &User{
			Email:       email,
			rawPassword: password,
		}
		res.Scan(&user.Id, &user.Nickname, &user.Password)
	}
	return user
}
