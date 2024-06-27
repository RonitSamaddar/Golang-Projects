package task

import (
	"time"
)

type Task struct {
	name      string
	createdAt time.Time
	timestamp time.Time
}
