package main

import (
	"log"

	"github.com/RonitSamaddar/Golang-Projects/basicwebserver/internal/server"
)

func main() {
	if err := server.Setup(); err != nil {
		log.Fatalf("Failed to setup server: %+v", err)
		return
	}
	log.Println("Server setup complete!")

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %+v", err)
		return
	}
	log.Println("Server started!")

}
