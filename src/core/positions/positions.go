package positions

import (
	"qutils/dbwrapper"
)

const (
	TABLENAME     = "positions"
	BUILDDISTANCE = 0.0003
)

var connection *dbwrapper.DBConnection = nil

//Using to connect module to the existing connection.
//IMPORTANT! Module should be initialized to work!
func Init(conn *dbwrapper.DBConnection) {
	connection = conn
}

func checkConnection() {
	if connection == nil {
		panic("There is no DBConnection in position module.")
	}
}
