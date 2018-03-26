package positions

import (
	"qutils/dbwrapper"
)

const (
	TABLENAME = "positions"
)

var connection *dbwrapper.DBConnection = nil

//Using to connect module to the existing connection.
//IMPORTANT! Module should be initialized to work!
func Initialize(conn *dbwrapper.DBConnection) {
	connection = conn
}

func checkConnection() {
	if connection == nil {
		panic("There is no DBConnection in position module.")
	}
}
