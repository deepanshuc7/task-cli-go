package main

import (
	"fmt"
	"os"
	"time"
)

const fileName = "tasks.json"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreateAt    time.Time `json:"createAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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
		fmt.Println("Add command selected")
	case "list":
		fmt.Println("List command selected")
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
