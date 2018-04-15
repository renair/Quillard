package sessions

import (
	"qutils/dbwrapper"
	"sync"
	"time"
)

const (
	TABLENAME    = "sessions"
	EXPIRETIME   = 45 * 60 //45 minutes
	CLEARINGTIME = time.Minute
)

var connection *dbwrapper.DBConnection = nil
var activeSessions sync.Map

//Initialize module to be able work with db.
//Required to work
func Init(conn *dbwrapper.DBConnection) {
	connection = conn
	go sessionClearer(CLEARINGTIME)
}
