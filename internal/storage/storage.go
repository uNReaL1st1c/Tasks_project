package storage

import (
	"encoding/json"
	"os"

	"github.com/uNReaL1st1c/Tasks_project/src/internal/config"
)

func SaveTasks[T any](filename string, tasks []T) error {

	data, err := json.Marshal(tasks)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, config.FilePerm)

	if err != nil {
		return err
	}

	return nil

}

func LoadTasks[T any](filename string) ([]T, error) {

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var tasks []T

	if len(data) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(data, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
