package storage

import (
	"encoding/json"
	"os"
	"taskTracker/src/internal/models"
)

func SaveTasks(filename string, tasks []models.Task) error {

	data, err := json.Marshal(tasks)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		return err
	}

	return nil

}

func LoadTasks(filename string) ([]models.Task, error) {

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var tasks []models.Task

	if len(data) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(data, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
