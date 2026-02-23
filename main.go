package main

import (
	"log"
	"taskTracker/src/internal/service"
	"taskTracker/src/internal/storage"
)

func main() {

	var fileName = "tasks.json"

	tasks, err := storage.LoadTasks(fileName)

	if err != nil {
		log.Fatalf("Ошибка загрузки: %v.", err)
		return
	}

	service.AddTask("Реализовать первый пункт", &tasks)
	service.AddTask("Проверить реализацию", &tasks)
	service.AddTask("Скормить DeepSeek для Code Review", &tasks)

	service.ListTasks(tasks)

	storage.SaveTasks(fileName, tasks)

}
