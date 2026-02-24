package service

import (
	"fmt"
	"taskTracker/src/internal/models"
)

func AddTask(title string, tasks *[]models.Task) {

	task := models.Task{
		ID:    generateID(tasks),
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

func ToDoTasks(tasks []models.Task) []models.Task {
	var toDoTasks []models.Task

	for _, task := range tasks {
		if !task.Done {
			toDoTasks = append(toDoTasks, task)
		}
	}

	return toDoTasks
}

func GetTaskByID(tasks []models.Task, ID int) *models.Task {
	for i := range tasks {
		if tasks[i].ID == ID {
			return &tasks[i]
		}
	}
	return nil
}

func DeleteTask(tasks *[]models.Task, ID int) error {
	if tasks == nil {
		return fmt.Errorf("tasks slice is nil")
	}

	if len(*tasks) == 0 {
		return fmt.Errorf("tasks slice is empty")
	}

	for index, task := range *tasks {
		if task.ID == ID {
			*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", ID)
}

func generateID(tasks *[]models.Task) int {

	if tasks == nil || len(*tasks) == 0 {
		return 1
	}

	maxID := 0
	for _, task := range *tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}
