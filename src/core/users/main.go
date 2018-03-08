package users

import (
	"dbwrapper"
)

const (
	TABLENAME = "users"
)

var connection *dbwrapper.DBConnection = nil

//Initialize module to work with db.
//Required to work
func Init(conn *dbwrapper.DBConnection) {
	connection = conn
}

//Create panic if connection not initialized
func chekInitialization() {
	if connection == nil {
		panic("DBConnection is not initialized!!!")
	}
}
