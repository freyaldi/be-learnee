package main

import (
	"log"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/db"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/server"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
	}
	server.Init()
}