package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RonitSamaddar/Golang-Projects/todolistserver/internal/task"
)

func RegisterRoutes() {
	http.HandleFunc("/create", newTaskHandler)
}

func newTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %+v", r.URL.String())

	taskName := r.URL.Query().Get("task")
	date := r.URL.Query().Get("date")
	time := r.URL.Query().Get("time")
	log.Printf("Task %s to be created for date %s, time %s\n", taskName, date, time)
	task.RegisterTask(taskName, date, time)
	fmt.Fprintln(w, "Task created!")
}
