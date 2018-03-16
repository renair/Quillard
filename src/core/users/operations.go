package users

import (
	"errors"
)

//Ruturn user object based on login and password
func logInUser(credential LoginRequest) *User {
	chekInitialization()
	keys := map[string]interface{}{
		"email":    credential.Email,
		"password": encodePassword(credential.Password),
	}
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

//Check if user already exist and if no create it
func registerUser(credential RegisterRequest) error {
	chekInitialization()
	if isFieldExist("nickname", credential.Nickname) || isFieldExist("email", credential.Email) {
		return errors.New("Email or Nickname Already exists.")
	}
	args := map[string]interface{}{
		"email":    credential.Email,
		"nickname": credential.Nickname,
		"password": encodePassword(credential.Password),
	}
	return connection.Insert(TABLENAME, args)
}
