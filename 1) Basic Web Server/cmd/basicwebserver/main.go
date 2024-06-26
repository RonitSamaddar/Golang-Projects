package main

import (
	"log"

	"github.com/RonitSamaddar/Golang-Projects/basicwebserver/internal/server"
)

func main() {
	if err := server.Setup(); err != nil {
		log.Fatalf("Failed to setup server: %+v", err)
	} else {
		log.Println("Server setup complete!")
	}

}
