package main

import (
	"dbwrapper"
	"fmt"
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
	} else {
		fmt.Println(err.Error())
	}
}
