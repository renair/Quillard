package account

import (
	"net/http"
	"qutils/dbwrapper"
)

const (
	TABLENAME = "accounts"
	APIPREFIX = "account"
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
