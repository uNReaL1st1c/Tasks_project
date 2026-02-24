package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"taskTracker/src/internal/service"
	"taskTracker/src/internal/storage"
)

var fileName = "tasks.json"

func main() {

	var (
		decision int
		isQuit   bool
	)

	for {
		currentMenu()
		fmt.Print("Ваш выбор: ")
		fmt.Scan(&decision)

		switch decision {
		case 1:
			addTask()
		case 2:
			viewAllTask()
		case 3:
			markTaskAsDone()
		case 4:
			deleteTask()
		case 5:
			isQuit = quitProgram()
		default:
			fmt.Println()
			fmt.Println("Неизвестный тип операции.")
			fmt.Println()
		}

		if isQuit {
			break
		}
	}
}

func currentMenu() {
	fmt.Println("📋 Менеджер задач v2.0")
	fmt.Println("======================")
	fmt.Println("1. ➕ Добавить задачу")
	fmt.Println("2. 📋 Показать все задачи")
	fmt.Println("3. ✅ Отметить задачу как выполненную")
	fmt.Println("4. ❌ Удалить задачу")
	fmt.Println("5. 🚪 Выйти")
	fmt.Println()
}

func addTask() {

	fmt.Print("Введите название задачи: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		tasks, err := storage.LoadTasks(fileName)

		if err != nil {
			log.Fatalf("Ошибка загрузки: %v.", err)
			return
		}

		service.AddTask(text, &tasks)
		storage.SaveTasks(fileName, tasks)
	}

	fmt.Println()
}

func viewAllTask() {

	tasks, err := storage.LoadTasks(fileName)

	if err != nil {
		log.Fatalf("Ошибка загрузки: %v.", err)
		return
	}

	service.ListTasks(tasks)

	fmt.Println()

}

func markTaskAsDone() {

	tasks, err := storage.LoadTasks(fileName)

	if err != nil {
		log.Fatalf("Ошибка загрузки: %v.", err)
		return
	}

	toDoTask := service.ToDoTasks(tasks)
	service.ListTasks(toDoTask)

	fmt.Print("Введите ID задачи для отметки: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		ID, err := strconv.Atoi(text)
		if err != nil {
			log.Fatalf("Выбор неопределен %v", err)
			return
		}
		task := service.GetTaskByID(tasks, ID)
		if task != nil {
			task.Done = true
		}
	}
	storage.SaveTasks(fileName, tasks)

	fmt.Println()
}

func deleteTask() {

	tasks, err := storage.LoadTasks(fileName)

	if err != nil {
		log.Fatalf("Ошибка загрузки: %v.", err)
		return
	}

	service.ListTasks(tasks)

	fmt.Print("Введите ID задачи для удаления: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		ID, err := strconv.Atoi(text)
		if err != nil {
			log.Fatalf("Выбор неопределен %v", err)
			return
		}
		service.DeleteTask(&tasks, ID)
	}
	storage.SaveTasks(fileName, tasks)

	fmt.Println()


}

func quitProgram() bool {

	tasks, err := storage.LoadTasks(fileName)

	if err != nil {
		log.Fatalf("Ошибка загрузки: %v.", err)
		return false
	}

	fmt.Println("💾 Сохраняем данные...")
	storage.SaveTasks(fileName, tasks)
	fmt.Println("👋 До свидания!")

	return true
}
