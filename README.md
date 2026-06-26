# Task Tracker CLI

A simple command-line Task Tracker built in **Go** that allows you to manage your tasks from the terminal. Tasks are stored locally in a JSON file, making the application lightweight and easy to use without requiring a database or external dependencies.

This project was built as a portfolio project to practice:

* Go fundamentals
* Command-line application development
* File handling
* JSON serialization/deserialization
* Error handling
* CRUD operations

---

## Features

* Add a new task
* Update an existing task
* Delete a task
* Mark a task as **in progress**
* Mark a task as **done**
* List all tasks
* Filter tasks by status:

  * `todo`
  * `in-progress`
  * `done`
* Automatic JSON file creation
* Automatic timestamps for task creation and updates
* Graceful error handling

---

## Project Structure

```text
task-cli/
│
├── main.go
├── go.mod
└── tasks.json      # Created automatically after adding the first task
```

---

## Task Structure

Each task contains the following properties:

```json
{
  "id": 1,
  "description": "Buy groceries",
  "status": "todo",
  "createdAt": "2026-06-26T15:42:18Z",
  "updatedAt": "2026-06-26T15:42:18Z"
}
```

| Property    | Description                              |
| ----------- | ---------------------------------------- |
| id          | Unique task identifier                   |
| description | Task description                         |
| status      | `todo`, `in-progress`, or `done`         |
| createdAt   | Date and time the task was created       |
| updatedAt   | Date and time the task was last modified |

---

## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/task-cli.git
```

Navigate into the project:

```bash
cd task-cli
```

Run the application:

```bash
go run main.go
```

---

## Usage

### Add a Task

```bash
go run main.go add "Buy groceries"
```

Output

```text
Task added successfully (ID: 1)
```

---

### List All Tasks

```bash
go run main.go list
```

Example

```text
ID: 1 | Status: todo | Description: Buy groceries
```

---

### List Tasks by Status

Todo tasks

```bash
go run main.go list todo
```

In Progress tasks

```bash
go run main.go list in-progress
```

Done tasks

```bash
go run main.go list done
```

---

### Update a Task

```bash
go run main.go update 1 "Buy groceries and cook dinner"
```

Output

```text
Task updated successfully
```

---

### Mark a Task as In Progress

```bash
go run main.go mark-in-progress 1
```

Output

```text
Task marked as in-progress
```

---

### Mark a Task as Done

```bash
go run main.go mark-done 1
```

Output

```text
Task marked as done
```

---

### Delete a Task

```bash
go run main.go delete 1
```

Output

```text
Task deleted successfully
```

---

## Example Workflow

```bash
go run main.go add "Learn Go"

go run main.go add "Build CLI"

go run main.go list

go run main.go mark-in-progress 1

go run main.go mark-done 1

go run main.go update 2 "Build Go CLI Portfolio Project"

go run main.go delete 1

go run main.go list
```

---

## Error Handling

The application handles common errors, including:

* Missing command
* Missing task description
* Invalid task ID
* Task not found
* Invalid status filter
* File read/write errors
* Invalid JSON format

---

## Technologies Used

* Go
* Go Standard Library

  * `os`
  * `encoding/json`
  * `time`
  * `strconv`
  * `fmt`

No external libraries or frameworks were used.

---

## Future Improvements

Possible enhancements include:

* Search tasks by keyword
* Sort tasks by creation date
* Task priorities
* Due dates
* Colored terminal output
* Interactive CLI
* Unit tests
* Export tasks to CSV
* Configuration file support

---

## Learning Outcomes

Through this project, I gained experience with:

* Building command-line applications in Go
* Working with JSON data
* Reading and writing files
* Parsing command-line arguments
* Structs and slices
* Error handling
* Organizing application logic
* Implementing CRUD operations

---

## License

This project is open source and available under the MIT License.
