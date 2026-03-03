package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/uNReaL1st1c/Tasks_project/src/internal/config"
	"github.com/uNReaL1st1c/Tasks_project/src/internal/models"
	"github.com/uNReaL1st1c/Tasks_project/src/internal/service"
	"github.com/uNReaL1st1c/Tasks_project/src/internal/storage"
)

var (
    activeTasks   = make(map[int]context.CancelFunc)
    doneChannel   = make(chan int)                    
)

func main() {

	var (
		isQuit bool
		input  string
	)

	go handleCompletedTasks()

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
			startWorkWithTask()
		case 6:
			stopWorkWithTask()
		case 7:
			activeTaskInProgress()
		case 8:
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

func handleCompletedTasks() {
    for completedID := range doneChannel {

        tasks, err := storage.LoadTasks[models.Task](config.FileName)
        if err != nil {
            fmt.Printf("❌ Ошибка загрузки: %v\n", err)
            continue
        }
        
        for i := range tasks {
            if tasks[i].ID == completedID {
                tasks[i].Done = true
                fmt.Printf("\n✅ Задача \"%s\" выполнена!\n", tasks[i].Title)
                break
            }
        }
        
        storage.SaveTasks(config.FileName, tasks)
        
        delete(activeTasks, completedID)
    }
}

func currentMenu() {
	fmt.Println(config.AppName, config.AppVersion)
	fmt.Println("======================")
	fmt.Println("1. ➕ Добавить задачу")
	fmt.Println("2. 📋 Показать все задачи")
	fmt.Println("3. ✅ Отметить задачу как выполненную")
	fmt.Println("4. ❌ Удалить задачу")
	fmt.Println("5. 🍅 Начать работу над задачей (таймер)")
	fmt.Println("6. ⏹ Остановить выполнение задачи")
	fmt.Println("7. 📊 Статус выполнения")
	fmt.Println("8. 🚪 Выйти")
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

		tasks, err := storage.LoadTasks[models.Task](config.FileName)
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

	tasks, err := storage.LoadTasks[models.Task](config.FileName)

	if err != nil {
		fmt.Printf("❌ Ошибка загрузки: %v\n", err)
		return
	}

	service.ListTasks(tasks)

	fmt.Println()

}

func markTaskAsDone() {

	tasks, err := storage.LoadTasks[models.Task](config.FileName)

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

	tasks, err := storage.LoadTasks[models.Task](config.FileName)

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

	tasks, err := storage.LoadTasks[models.Task](config.FileName)

	if err != nil {
		fmt.Printf("❌ Ошибка загрузки: %v\n", err)
		return false
	}

	fmt.Println("💾 Сохраняем данные...")
	storage.SaveTasks(config.FileName, tasks)
	fmt.Println("👋 До свидания!")

	return true
}

func startWorkWithTask() {
    tasks, err := storage.LoadTasks[models.Task](config.FileName)
    if err != nil {
        fmt.Printf("❌ Ошибка загрузки: %v\n", err)
        return
    }

    toDoTask := service.ToDoTasks(tasks)
    if len(toDoTask) == 0 {
        fmt.Println("📭 Нет задач для выполнения")
        return
    }
    
    service.ListTasks(toDoTask)

    fmt.Print("Введите ID задачи: ")
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        return
    }
    
    text := scanner.Text()
    ID, err := strconv.Atoi(text)
    if err != nil {
        fmt.Printf("❌ Ошибка: %v\n", err)
        return
    }

    if _, exists := activeTasks[ID]; exists {
        fmt.Printf("❌ Задача %d уже выполняется\n", ID)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    activeTasks[ID] = cancel

    activeTasksList, _ := storage.LoadTasks[models.ActiveTask](config.FileNameForActiveTask)
    
    task := service.GetTaskByID(tasks, ID)
    if task == nil {
        fmt.Println("❌ Задача не найдена")
        return
    }

    service.AddActiveTask(task.ID, task.Title, &activeTasksList)
    storage.SaveTasks(config.FileNameForActiveTask, activeTasksList)

    fmt.Printf("▶️ Запущен таймер для задачи \"%s\" (10 секунд)\n", task.Title)

    go func(taskID int, cancelFunc context.CancelFunc) {
        defer cancelFunc()          
        defer delete(activeTasks, taskID)
        
        select {
        case <-time.After(10 * time.Second):
            doneChannel <- taskID
        case <-ctx.Done():
            fmt.Printf("\n⏹ Задача %d остановлена\n", taskID)
            return
        }
    }(ID, cancel)
}

func activeTaskInProgress() {
    activeTasksList, err := storage.LoadTasks[models.ActiveTask](config.FileNameForActiveTask)
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Println("📭 Нет активных задач")
            return
        }
        fmt.Printf("❌ Ошибка загрузки: %v\n", err)
        return
    }

    if len(activeTasksList) == 0 {
        fmt.Println("📭 Нет активных задач")
        return
    }

    fmt.Println("\n📊 Активные задачи:")
    for _, task := range activeTasksList {
        fmt.Printf("  %d. 🍅 %s\n", task.ID, task.Title)
    }
    fmt.Printf("🔢 Всего: %d\n\n", len(activeTasksList))
}

func stopWorkWithTask() {
    activeTasksList, err := storage.LoadTasks[models.ActiveTask](config.FileNameForActiveTask)
    if err != nil || len(activeTasksList) == 0 {
        fmt.Println("📭 Нет активных задач для остановки")
        return
    }

    service.ListActiveTasks(activeTasksList)

    fmt.Print("Введите ID задачи для остановки: ")
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        return
    }
    
    text := scanner.Text()
    ID, err := strconv.Atoi(text)
    if err != nil {
        fmt.Printf("❌ Ошибка: %v\n", err)
        return
    }

    if cancel, exists := activeTasks[ID]; exists {
        cancel()
        
        service.DeleteTask(&activeTasksList, ID)
        storage.SaveTasks(config.FileNameForActiveTask, activeTasksList)
        
        fmt.Printf("⏹ Задача %d остановлена\n", ID)
    } else {
        fmt.Printf("❌ Задача %d не выполняется\n", ID)
    }
}
