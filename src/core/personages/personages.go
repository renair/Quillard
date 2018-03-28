package personages

import (
	"net/http"
	"qutils/dbwrapper"
)

const (
	TABLENAME = "personages"
	APIPREFIX = "personage"
)

var connection *dbwrapper.DBConnection = nil

func Init(conn *dbwrapper.DBConnection) {
	connection = conn
}

func checkConnection() {
	if connection == nil {
		panic("Personages module not initialized.")
	}
}

func ExportedHandlers() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"create": CreatePersonageHandler,
		"list":   GetOwnPersonagesHandler,
	}
}
