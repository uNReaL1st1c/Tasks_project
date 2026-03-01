package models

import "context"

type Identifiable interface {
	GetID() int
}

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type ActiveTask struct {
	ID     int
	Title  string
	Cancel context.CancelFunc
}

func (task Task) GetID() int { return task.ID }

func (task ActiveTask) GetID() int { return task.ID }
