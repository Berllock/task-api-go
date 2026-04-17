package main

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}

const (
	StatusTodo       = "todo"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)

func saveTasks(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile("tasks.json", data, 0644)
}

//func loadTasks([]Task, error)  {
//	data, err := os.ReadFile("tasks.json")
//	if err != nil {
//		if os.IsNotExist(err) {
//			return []Task{}, nil
//		}
//		return nil, err
//	}
//
//	var tasks []Task
//
//	err = json.Unmarshal(data, &tasks)
//	if err != nil {
//		return nil, err
//	}
//	return tasks, nil
//
//}

func main() {

}
