package account

import (
	"core/positions"
	"errors"
)

//Ruturn user object based on login and password
func logInAccount(credential LoginRequest) *Account {
	chekInitialization()
	keys := map[string]interface{}{
		"email":    credential.Email,
		"password": encodePassword(credential.Password),
	}
	res, err := connection.SelectBy(TABLENAME, keys, "id", "password", "home_position")
	defer res.Close()
	var user *Account = nil
	for err == nil && res.Next() {
		user = &Account{
			Email:       credential.Email,
			rawPassword: credential.Password,
		}
		res.Scan(&user.Id, &user.Password, &user.HomePositionId)
	}
	return user
}

//Check if account already exist and if no - create it
func registerAccount(credential RegisterRequest) (*Account, error) {
	chekInitialization()
	if !positions.CanBuild(credential.Position) {
		return nil, errors.New("position")
	}
	if isFieldExist("email", credential.Email) {
		return nil, errors.New("email")
	}
	connection.BeginTransaction()
	userPos := positions.SavePosition(credential.Position)
	args := map[string]interface{}{
		"email":         credential.Email,
		"password":      encodePassword(credential.Password),
		"home_position": userPos.Id,
	}
	insertError := connection.Insert(TABLENAME, args)
	if insertError != nil {
		connection.RollbackTransaction()
		return nil, insertError
	}
	connection.CommitTransaction()
	return logInAccount(LoginRequest{credential.Email, credential.Password}), nil
}
