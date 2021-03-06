package main

import (
	//core packages
	"core/account"
	"core/personages"
	"core/positions"
	"core/resources"
	"core/sessions"

	//other packages
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
	account.Init(&connection)
	positions.Init(&connection)
	personages.Init(&connection)
	resources.Init(&connection)
	// Setup web handlers
	// core/account handlers
	for url, handler := range account.ExportedHandlers() {
		absoluteUrl := fmt.Sprintf("/%s/%s", account.APIPREFIX, url)
		handlerMux.HandleFunc(absoluteUrl, handler)
	}
	//core/personages handler
	for url, handler := range personages.ExportedHandlers() {
		absoluteUrl := fmt.Sprintf("/%s/%s", personages.APIPREFIX, url)
		handlerMux.HandleFunc(absoluteUrl, handler)
	}
	//core/positions
	for url, handler := range positions.ExportedHandlers() {
		absoluteUrl := fmt.Sprintf("/%s/%s", positions.APIPREFIX, url)
		handlerMux.HandleFunc(absoluteUrl, handler)
	}
	//core/resources
	for url, handler := range resources.ExportedHandlers() {
		absoluteUrl := fmt.Sprintf("/%s/%s", resources.APIPREFIX, url)
		handlerMux.HandleFunc(absoluteUrl, handler)
	}
	//Setup static files handling
	handlerMux.Handle("/", getFileHandler("/", "web"))
	// Starting server
	log.Println("Server ready. Starting...")
	webServer.Handler = handlerMux
	webServer.ListenAndServe()
}
