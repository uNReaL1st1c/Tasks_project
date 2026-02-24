package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"taskTracker/src/internal/service"
	"taskTracker/src/internal/storage"
)

var fileName = "tasks.json"

func main() {

	var (
		isQuit bool
		input  string
	)

	for {
		currentMenu()
		fmt.Print("Ваш выбор: ")
		fmt.Scan(&input)
		decision, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("❌ Ошибка: введите число")
			continue
		}

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
		if text == "" {
			fmt.Println("❌ Название не может быть пустым")
			return
		}

		tasks, err := storage.LoadTasks(fileName)
		if err != nil {
			fmt.Printf("❌ Ошибка загрузки: %v\n", err)
			return
		}

		service.AddTask(text, &tasks)
		storage.SaveTasks(fileName, tasks)
		fmt.Printf("✅ Задача \"%s\" добавлена (ID: %d)\n",
			text, len(tasks))
	}
	fmt.Println()
}

func viewAllTask() {

	tasks, err := storage.LoadTasks(fileName)

	if err != nil {
		fmt.Printf("❌ Ошибка загрузки: %v\n", err)
		return
	}

	service.ListTasks(tasks)

	fmt.Println()

}

func markTaskAsDone() {

	tasks, err := storage.LoadTasks(fileName)

	if err != nil {
		fmt.Printf("❌ Ошибка загрузки: %v\n", err)
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
			fmt.Printf("Выбор неопределен %v", err)
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
		fmt.Printf("❌ Ошибка загрузки: %v\n", err)
		return
	}

	service.ListTasks(tasks)

	fmt.Print("Введите ID задачи для удаления: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		ID, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("Выбор неопределен %v", err)
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
		fmt.Printf("❌ Ошибка загрузки: %v\n", err)
		return false
	}

	fmt.Println("💾 Сохраняем данные...")
	storage.SaveTasks(fileName, tasks)
	fmt.Println("👋 До свидания!")

	return true
}
