package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil

}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Comandos disponíveis: add, list, update, delete")
		return
	}

	switch args[1] {
	case "add":
		if len(args) < 3 {
			fmt.Println("Erro: Forneça uma descrição para a tarefa.")
			return
		}
		description := args[2]

		task, err := loadTasks()
		if err != nil {
			fmt.Println("Erro ao carregar tarefas:", err)
			return
		}

		novaTask := Task{
			ID:          len(task) + 1,
			Description: description,
			Status:      StatusTodo,
			CreatedAt:   time.Now(),
			UpdateAt:    time.Now(),
		}

		task = append(task, novaTask)

		err = saveTasks(task)
		if err != nil {
			fmt.Println("Erro ao salvar tarefa:", err)
			return
		}

		fmt.Printf("Tarefa adicionada com sucesso! (ID: %d)\n", novaTask.ID)

	case "list":

		tasks, err := loadTasks()

		if err != nil {
			fmt.Println("Erro ao carregar tarefas:", err)
			return
		}

		for _, task := range tasks {
			fmt.Println(task)
		}

	case "update":
		if len(args) < 4 {
			fmt.Println("Erro: Forneça o ID e a descrição para a atualização da terefa")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Erro: ID inválido. Use um numero")
			return
		}

		newDescription := args[3]

		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Erro ao carregar tarefas:", err)
			return
		}

		findTask := false
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Description = newDescription
				tasks[i].UpdateAt = time.Now()
				findTask = true
				break
			}
		}

		if !findTask {
			fmt.Printf("Erro: Tarefa com ID %d não encontrada. \n", id)
			return
		}

		err = saveTasks(tasks)
		if err != nil {
			fmt.Println("Erro ao salvar tarefa:", err)
			return
		}

		fmt.Printf("Tarefa atualizada com sucesso! (ID: %d)\n", id)

	}

}
