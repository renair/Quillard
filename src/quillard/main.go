package main

import (
	"core/users"
	"dbwrapper"
	"fmt"
	"net/http"
)

const (
	USERNAME = "postgres"
	PASSWORD = "abcd1234"
	DBNAME   = "quillard"
)

func main() {
	connection := dbwrapper.DBConnection{
		User:     USERNAME,
		Password: PASSWORD,
		Dbname:   DBNAME,
	}
	if err := connection.Connect(); err == nil {
		fmt.Println("Connected!")
		users.Init(&connection)
		http.HandleFunc("/", users.ExportHandlers()[0])
		http.ListenAndServe(":8080", nil)
	} else {
		fmt.Println(err.Error())
	}
}
