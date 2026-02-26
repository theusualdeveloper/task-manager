# Task Manager CLI

A command-line task management tool built with Go.

## Features

- Add new tasks
- List all tasks
- Mark tasks as done
- Delete tasks
- Persistent storage using JSON

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.21 or higher

### Installation

```bash
git clone https://github.com/YourGitHubUsername/task-manager.git
cd task-manager
go build -o task-manager .
```

### Usage

```bash
# Add a task
./task-manager add "Buy groceries"

# List all tasks
./task-manager list

# Mark a task as done
./task-manager done 1

# Delete a task
./task-manager delete 1
```

## Project Structure

```
task-manager/
├── main.go         # Entry point
├── task/
│   └── task.go     # Task struct and logic
└── go.mod
```

## Tech Stack

- **Language:** Go
- **Storage:** JSON file

## License

MIT
