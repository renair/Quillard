package main

import (
	"core/users"
	"dbwrapper"
	"fmt"
	"net/http"
)

const (
	HOST     = "platinium.ddns.net"
	USERNAME = "postgres"
	PASSWORD = "abcd1234"
	DBNAME   = "quillard"
)

func main() {
	connection := dbwrapper.DBConnection{
		Host:     HOST,
		User:     USERNAME,
		Password: PASSWORD,
		Dbname:   DBNAME,
	}
	if err := connection.Connect(); err == nil {
		fmt.Println("Connected!")
		users.Init(&connection)
		for url, handler := range users.ExportedHandlers() {
			absoluteUrl := fmt.Sprintf("/%s/%s", users.APIPREFIX, url)
			http.HandleFunc(absoluteUrl, handler)
		}
		http.ListenAndServe(":8080", nil)
	} else {
		fmt.Println(err.Error())
	}
}
