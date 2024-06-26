package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RonitSamaddar/Golang-Projects/basicwebserver/internal/handlers"
)

func Setup() error {
	log.Printf("Setting up server")
	handlers.RegisterRoutes()
	return nil
}

func Start() error {
	log.Printf("Starting server")

	port := "8080"
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		return err
	}

	return nil
}
