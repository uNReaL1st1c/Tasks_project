package service

import (
	"fmt"

	"github.com/uNReaL1st1c/Tasks_project/src/internal/models"
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

	if len(tasks) == 0 {
		fmt.Println("📭 Список задач пуст")
		return
	}

	doneCount := 0
	for _, task := range tasks {
		status := "⏳"
		if task.Done {
			status = "✅"
			doneCount++
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
	}
	fmt.Printf("\n📊 Итого: ✅ %d выполнено, ⏳ %d осталось\n",
		doneCount, len(tasks)-doneCount)
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

func generateID[T models.Identifiable](items *[]T) int {

	if items == nil || len(*items) == 0 {
		return 1
	}

	maxID := 0
	for _, item := range *items {
		ID := item.GetID()
		if ID > maxID {
			maxID = ID
		}
	}

	return maxID + 1
}

func AddActiveTask(title string, activeTasks *[]models.ActiveTask) {
	activeTask := models.ActiveTask{
		ID:    generateID(activeTasks),
		Title: title,
	}
	*activeTasks = append(*activeTasks, activeTask)
}
