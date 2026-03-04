package task

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

type TaskStore struct {
	tasks []Task
	*JSONStorage
}

func NewTaskStore(filePath string) (*TaskStore, error) {
	js := NewJSONStorage(filePath)
	tasks, err := js.Load()
	if err != nil {
		return nil, fmt.Errorf("Error in load tasks: %s", err)
	}
	ts := &TaskStore{
		tasks:       tasks,
		JSONStorage: &js,
	}
	return ts, nil
}

func (ts *TaskStore) Add(title string) (Task, error) {
	lid := 0
	if len(ts.tasks) > 0 {
		lid = ts.tasks[len(ts.tasks)-1].ID
	}
	task := Task{
		ID:        lid + 1,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	ts.tasks = append(ts.tasks, task)
	err := ts.JSONStorage.Save(ts.tasks)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (ts *TaskStore) List() []Task {
	return ts.tasks
}

func (ts *TaskStore) Complete(id int) error {
	ti := -1
	for i, t := range ts.tasks {
		if t.ID == id {
			ti = i
			break
		}
	}
	if ti < 0 {
		return errors.New("task not found")
	}
	ts.tasks[ti].Done = true
	err := ts.JSONStorage.Save(ts.tasks)
	if err != nil {
		return fmt.Errorf("error in saving tasks: %s", err)
	}
	return nil
}

func (ts *TaskStore) Delete(id int) error {
	filtered := make([]Task, 0)
	found := -1
	for _, t := range ts.tasks {
		if t.ID == id {
			found = t.ID
			continue
		}
		filtered = append(filtered, t)
	}
	if found < 0 {
		return errors.New("task not found")
	}
	ts.tasks = filtered
	err := ts.JSONStorage.Save(ts.tasks)
	if err != nil {
		return fmt.Errorf("error in saving tasks: %s", err)
	}
	return nil
}
