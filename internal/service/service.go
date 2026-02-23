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
		if task.Done {
			fmt.Printf("%d. [%s] %s\n", task.ID, task.Title, "успешно выполнена.")
		} else {
			fmt.Printf("%d. [%s] %s\n", task.ID, task.Title, "еще выполняется.")
		}
	}
}
