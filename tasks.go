package main

import "time"

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(description string) Task {
	return Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) MarkDone(){
	t.Status = "done"
	t.UpdatedAt = time.Now()
	
}

func (t *Task) MarkInProgress(){
	t.Status = "In Progress"
	t.UpdatedAt = time.Now()

}

