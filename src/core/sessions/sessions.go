package sessions

import (
	"dbwrapper"
)

const (
	TABLENAME  = "sessions"
	EXPIRETIME = 45 * 60 //45 minutes
)

var connection *dbwrapper.DBConnection = nil

//Initialize module to be able work with db.
//Required to work
func Init(conn *dbwrapper.DBConnection) {
	connection = conn
}
