package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func findTaskIndexByID(tasks []Task, id int) int {
	for index, task := range tasks {
		if task.ID == id {
			return index
		}
	}

	return -1
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
			fmt.Println("Error loading tasks:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		filter := ""
		if len(args) >= 3 {
			filter = args[2]

			if filter != "todo" && filter != "in-progress" && filter != "done" {
				fmt.Println("Error: invalid status filter")
				fmt.Println("Allowed filters: todo, in-progress, done")
				return
			}
		}

		found := false

		for _, task := range tasks {
			if filter == "" || task.Status == filter {
				fmt.Printf(
					"ID: %d | Status: %s | Description: %s\n",
					task.ID,
					task.Status,
					task.Description,
				)
				found = true
			}
		}

		if !found {
			fmt.Println("No tasks found for this status.")
		}
	case "update":
		if len(args) < 4 {
			fmt.Println(`Usage: task-cli update <id> "New description"`)
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: task ID must be a number")
			return
		}

		newDescription := args[3]

		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		index := findTaskIndexByID(tasks, id)
		if index == -1 {
			fmt.Println("Error: task not found")
			return
		}

		tasks[index].Description = newDescription
		tasks[index].UpdatedAt = time.Now()

		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}

		fmt.Println("Task updated successfully")
	case "delete":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: task ID must be a number")
			return
		}

		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		index := findTaskIndexByID(tasks, id)
		if index == -1 {
			fmt.Println("Error: task not found")
			return
		}

		tasks = append(tasks[:index], tasks[index+1:]...)

		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}

		fmt.Println("Task deleted successfully")
	case "mark-in-progress":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: task ID must be a number")
			return
		}

		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		index := findTaskIndexByID(tasks, id)
		if index == -1 {
			fmt.Println("Error: task not found")
			return
		}

		tasks[index].Status = "in-progress"
		tasks[index].UpdatedAt = time.Now()

		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}

		fmt.Println("Task marked as in-progress")
	case "mark-done":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli mark-done <id>")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: task ID must be a number")
			return
		}

		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		index := findTaskIndexByID(tasks, id)
		if index == -1 {
			fmt.Println("Error: task not found")
			return
		}

		tasks[index].Status = "done"
		tasks[index].UpdatedAt = time.Now()

		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}

		fmt.Println("Task marked as done")
	default:
		fmt.Println("Unknown command:", command)
	}
}
