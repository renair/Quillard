package main

import (
	"core/sessions"
	"core/users"
	"fmt"
	"log"
	"net/http"
	"qutils/dbwrapper"
	"time"
)

const (
	HOST     = "platinium.ddns.net"
	USERNAME = "postgres"
	PASSWORD = "abcd1234"
	DBNAME   = "quillard"
)

func main() {
	// Setup working environment
	// Web environment
	handlerMux := http.NewServeMux()
	webServer := http.Server{
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 20,
		Addr:         ":8080",
	}
	// DB connection
	connection := dbwrapper.DBConnection{
		Host:     HOST,
		User:     USERNAME,
		Password: PASSWORD,
		Dbname:   DBNAME,
	}
	if err := connection.Connect(); err != nil {
		log.Fatalln("Critical db connection error!")
		log.Fatalln(err.Error())
		return
	}
	log.Println("DB connection established")
	// Init modules with DB connection
	sessions.Init(&connection)
	users.Init(&connection)
	// Setup web handlers
	// core/users handlers
	for url, handler := range users.ExportedHandlers() {
		absoluteUrl := fmt.Sprintf("/%s/%s", users.APIPREFIX, url)
		handlerMux.HandleFunc(absoluteUrl, handler)
	}
	//Setup static files handling
	handlerMux.Handle("/", getFileHandler("/", "web"))
	// Starting server
	log.Println("Server ready. Starting...")
	webServer.Handler = handlerMux
	webServer.ListenAndServe()
}
