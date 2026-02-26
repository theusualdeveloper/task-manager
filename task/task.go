package task

import (
	"errors"
	"time"
)

type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

type TaskStore struct {
	tasks   []Task
	counter int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{}
}

func (ts *TaskStore) Add(title string) Task {
	task := Task{
		ID:        ts.counter + 1,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	ts.tasks = append(ts.tasks, task)
	ts.counter++
	return task
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
	return nil
}
