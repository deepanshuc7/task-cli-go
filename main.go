package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const fileName = "tasks.json"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func loadTasks() ([]Task, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []Task{}, nil
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return nil
	}

	return os.WriteFile(fileName, data, 0644)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: task cli <command> [arguments]")
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Error: task description is required")
			fmt.Println(`Usage: task-cli add "Buy groceries"`)
			return
		}

		description := args[2]

		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		newID := 1
		if len(tasks) > 0 {
			newID = tasks[len(tasks)-1].ID + 1
		}

		now := time.Now()

		newTask := Task{
			ID:          newID,
			Description: description,
			Status:      "todo",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		tasks = append(tasks, newTask)

		err = saveTasks(tasks)
		if err != nil {
			fmt.Println("Error saving task:", err)
			return
		}

		fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)

	case "list":
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Erro loading tasks: ", err)
			return
		}

		fmt.Println(tasks)
	case "update":
		fmt.Println("Update command selected")
	case "delete":
		fmt.Println("Delete command selected")
	case "mark-in-progress":
		fmt.Println("Mark in progress command selected")
	case "mark-done":
		fmt.Println("Mark done command selected")
	default:
		fmt.Println("Unknown command:", command)
	}
}
