package users

import (
	"net/http"
	"qutils/dbwrapper"
)

const (
	TABLENAME = "users"
	APIPREFIX = "user"
)

var connection *dbwrapper.DBConnection = nil

//Initialize module to be able work with db.
//Required to work
func Init(conn *dbwrapper.DBConnection) {
	connection = conn
}

func ExportedHandlers() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"login":    LoginHandler,
		"register": RegisterHandler,
	}
}
