package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/greeting", greetingHandler)
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %+v", r.URL.String())

	user := r.URL.Query().Get("user")
	fmt.Fprintf(w, "Hello %s", user)
}
