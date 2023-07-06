package main

import (
	server "github.com/KayoRonald/golang-api/app/cmd"
	"github.com/KayoRonald/golang-api/app/database"
)

func main() {
	// inialize database
	database.ConnectDB()
	err := server.Run()
	if err != nil {
		panic(err)
	}
}