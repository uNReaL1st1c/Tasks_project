package main

import (
	"fmt"
	"taskTracker/src/internal/service"
	"taskTracker/src/internal/storage"
)

func main() {

	var fileName = "tasks.json"

	tasks, err := storage.LoadTasks(fileName)

	if err != nil {
		fmt.Println(err)
	}

	service.AddTask("Реализовать первый пункт", &tasks)
	service.AddTask("Проверить реализацию", &tasks)
	service.AddTask("Скормить DeepSeek для Code Review", &tasks)

	service.ListTasks(tasks)

	storage.SaveTasks(fileName, tasks)

}
