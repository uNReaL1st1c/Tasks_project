package service

import (
	"fmt"
	"taskTracker/src/internal/models"
)

func AddTask(title string, tasks *[]models.Task) {
	task := models.Task{
		ID:    len(*tasks) + 1,
		Title: title,
		Done:  false,
	}
	*tasks = append(*tasks, task)
}

func ListTasks(tasks []models.Task) {
	for _, task := range tasks {
		status := "⏳"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
	}
}
