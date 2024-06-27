package task

import (
	"fmt"
	"log"
	"time"
)

func RegisterTask(name string, date string, time string) error {
	task, err := createTask(name, date, time)
	if err != nil {
		log.Fatalf("Error in creating task for name %s, date %s, time %s : %+v", name, date, time, err)
		return err
	}
	log.Printf("Created task : %+v", task)
	return nil
}

func createTask(name string, date string, timeString string) (*Task, error) {
	timeString = fmt.Sprintf("%s %s", date, timeString)
	timestamp, err := time.Parse("2000-01-01 00:00", timeString)
	if err != nil {
		return nil, err
	}

	task := Task{
		name:      name,
		createdAt: time.Now().In(time.Local),
		timestamp: timestamp,
	}
	return &task, nil
}
