package v1

import "time"

// Status defines possible states of a To Do item
type Status string

const (
	NotStarted Status = "NotStarted"
	InProgress Status = "InProgress"
	Completed  Status = "Completed"
)

// ToDoItem defines the data model for each task
type ToDoItem struct {
	ID        int
	Title     string
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
	Due       *time.Time  // pointer because due date is optional
}
