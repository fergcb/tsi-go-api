package main

import (
	"tsi.co/go-api/database"
	"tsi.co/go-api/resources/actors"
	"tsi.co/go-api/server"
)

func main() {
	database.Init()
	database.DB.AutoMigrate(&actors.Actor{})

	server.Init()
}
