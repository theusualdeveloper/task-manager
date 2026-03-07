package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type JSONStorage struct {
	filePath string
}

func NewJSONStorage(path string) JSONStorage {
	return JSONStorage{
		filePath: path,
	}
}

func (js JSONStorage) Save(tasks []Task) error {
	b, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("error in marshaling: %w", err)
	}
	err = os.WriteFile(js.filePath, b, 0644)
	if err != nil {
		return fmt.Errorf("error in writing to the file: %w", err)
	}
	return nil
}

func (js JSONStorage) Load() ([]Task, error) {
	b, err := os.ReadFile(js.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("error in reading file: %w", err)
	}
	var tasks []Task
	err = json.Unmarshal(b, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error in unmarshaling: %w", err)
	}
	return tasks, nil
}
