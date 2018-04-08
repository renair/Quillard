package resources

import (
	"net/http"
	"qutils/dbwrapper"
)

const (
	STORAGETABLE = "resource_storages"
	TYPETABLE    = "recource_types"
	APIPREFIX    = "resources"
)

var connection *dbwrapper.DBConnection = nil

func Init(conn *dbwrapper.DBConnection) {
	connection = conn
	getResourceTypes()
}

func ExportedHandlers() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"myresources": PersonageResourcesHandler,
	}
}
