package users

import (
	"dbwrapper"
	"net/http"
)

const (
	TABLENAME = "users"
	APIPREFIX = "user"
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

func ExportedHandlers() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"login":    LoginHandler,
		"register": RegisterHandler,
	}
}
