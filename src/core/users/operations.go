package users

import (
	"crypto/sha1"
	"errors"
	"fmt"
)

//Ruturn user object based on login and password
func logInUser(credential LoginRequest) *User {
	chekInitialization()
	keys := make(map[string]interface{})
	keys["email"] = credential.Email
	keys["password"] = fmt.Sprintf("%x", sha1.Sum([]byte(credential.Password)))
	res, err := connection.SelectBy(TABLENAME, keys, "id", "nickname", "password")
	defer res.Close()
	var user *User = nil
	for err == nil && res.Next() {
		user = &User{
			Email:       credential.Email,
			rawPassword: credential.Password,
		}
		res.Scan(&user.Id, &user.Nickname, &user.Password)
	}
	return user
}

func registerUser(credential RegisterRequest) error {
	chekInitialization()
	if isNicknameExist(credential.Nickname) || isEmailRegistered(credential.Email) {
		return errors.New("Email or Nickname Already exists.")
	}
	args := map[string]interface{}{
		"email":    credential.Email,
		"nickname": credential.Nickname,
		"password": fmt.Sprintf("%x", sha1.Sum([]byte(credential.Password))),
	}
	return connection.Insert(TABLENAME, args)
}
