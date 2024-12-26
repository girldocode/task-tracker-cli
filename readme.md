# Task Tracker

A simple command-line interface (CLI) tool for managing tasks. This project allows you to add, update, delete, and list tasks, as well as mark them as "in-progress" or "done."

### Features

- Add new tasks with a description.
- Update task descriptions.
- Delete tasks by their ID.
- Mark tasks as in-progress or done.
- List all tasks or filter by status (todo, in-progress, done).

## Installation

```Golang
git clone git@github.com:girldocode/task-tracker-cli.git
```

run command:

```Golang
task-cli go run main.go
```

### Requirements

Requirements
Go 1.18 or higher

## Usage

- Add a Task

```Golang
task-cli go run main.go add "Buy groceries"
```

- Update a Task

```Golang
task-cli go run main.go update 1 "Buy groceries and cook dinner"

```

- Delete a Task

```Golang
task-cli go run main.go delete 1


```

- Mark a Task as In-Progress

```Golang
task-cli go run main.go mark-in-progress 1

```

- Mark a Task as Done

```Golang
task-cli go run main.go mark-done 1

```

- List All Tasks

```Golang
task-cli go run main.go list

```

- List Tasks by Status

```Golang
# task-cli go run main.go list todo
# task-cli go run main.go list in-progress
# task-cli go run main.go list in-progress

```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
