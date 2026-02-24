package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/uNReaL1st1c/Tasks_project/src/internal/config"
	"github.com/uNReaL1st1c/Tasks_project/src/internal/service"
	"github.com/uNReaL1st1c/Tasks_project/src/internal/storage"
)

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
	fmt.Println(config.AppName, config.AppVersion)
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

		tasks, err := storage.LoadTasks(config.FileName)
		if err != nil {
			fmt.Printf("❌ Ошибка загрузки: %v\n", err)
			return
		}

		service.AddTask(text, &tasks)
		storage.SaveTasks(config.FileName, tasks)
		fmt.Printf("✅ Задача \"%s\" добавлена (ID: %d)\n",
			text, len(tasks))
	}
	fmt.Println()
}

func viewAllTask() {

	tasks, err := storage.LoadTasks(config.FileName)

	if err != nil {
		fmt.Printf("❌ Ошибка загрузки: %v\n", err)
		return
	}

	service.ListTasks(tasks)

	fmt.Println()

}

func markTaskAsDone() {

	tasks, err := storage.LoadTasks(config.FileName)

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
			fmt.Printf("✅ Задача \"%s\" отмечена как выполненная\n", task.Title)
		} else {
			fmt.Printf("❌ Задача с ID %d не найдена\n", ID)
		}
	}
	storage.SaveTasks(config.FileName, tasks)

	fmt.Println()
}

func deleteTask() {

	tasks, err := storage.LoadTasks(config.FileName)

	if err != nil {
		fmt.Printf("❌ Ошибка загрузки: %v\n", err)
		return
	}

	fmt.Print("Введите ID задачи для удаления: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		ID, _ := strconv.Atoi(text)

		task := service.GetTaskByID(tasks, ID)
		if task == nil {
			fmt.Printf("❌ Задача с ID %d не найдена\n", ID)
			return
		}

		fmt.Printf("Удалить задачу \"%s\"? (y/N): ", task.Title)
		scanner.Scan()
		confirm := scanner.Text()
		if confirm != "y" && confirm != "Y" {
			fmt.Println("❌ Удаление отменено")
			return
		}

		service.DeleteTask(&tasks, ID)
		storage.SaveTasks(config.FileName, tasks)
		fmt.Printf("✅ Задача \"%s\" удалена\n", task.Title)
	}

}

func quitProgram() bool {

	tasks, err := storage.LoadTasks(config.FileName)

	if err != nil {
		fmt.Printf("❌ Ошибка загрузки: %v\n", err)
		return false
	}

	fmt.Println("💾 Сохраняем данные...")
	storage.SaveTasks(config.FileName, tasks)
	fmt.Println("👋 До свидания!")

	return true
}
