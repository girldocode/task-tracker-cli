
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"` // "todo", "in-progress", "done"
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
	// fmt.Println("Enter Your First Name: ")

	// var first string
	// fmt.Scanln(&first)

	// fmt.Println("Enter Your Last Name: ")

	// var second string
	// fmt.Scanln(&second)

	// fmt.Println("Hello " + first + " " + second)	

	//provide arguments
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	// command: add, delete, update, list, mark, mark-in-progress, task-by-status
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <description>")
			return
		}
		description := os.Args[2]
		addTask(description)
	case "list":
		listTasks()
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> <new-description>")
			return
		}
		id := parseID(os.Args[2])
		newDescription := os.Args[3]
		updateTask(id, newDescription)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id := parseID(os.Args[2])
		deleteTask(id)
	case "mark":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli mark <id> <status>")
			return
		}
		id := parseID(os.Args[2])
		status := os.Args[3]
		if status != "todo" && status != "in-progress" && status != "done" {
			fmt.Println("Invalid status. Valid statuses are: todo, in-progress, done")
			return
		}
		markTask(id, status)
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}
		id := parseID(os.Args[2])
		markTask(id, "in-progress")
	case "list-task-by-status":
		if len(os.Args) == 3 {
			status := os.Args[2]
			if status == "todo" || status == "done" || status == "in-progress" {
				listTasksByStatus(status)
			} else {
				fmt.Println("Invalid status. Use 'todo', 'in-progress', or 'done'.")
			}
		} else {
			listTasks()
		}
	default:
		fmt.Println("Unknown command:", command)
	}
}
func parseID(idStr string) int {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. ID must be a number.")
		os.Exit(1) 
	}
	return id
}


// read from json file
func readTasks() []Task {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("Error decoding JSON:", err)
	}
	return tasks
}

// save to json file
func saveTasks(tasks []Task) {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

// add task 
func addTask(description string) {
	tasks := readTasks()

	newTask := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	tasks = append(tasks, newTask)

	saveTasks(tasks)
	fmt.Println("Task added successfully!")
}

// update task 
func updateTask(id int, newDescription string) {
	tasks := readTasks() 

	for i, task := range tasks { 
		if task.ID == id { 
			tasks[i].Description = newDescription 
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339) 
			saveTasks(tasks) 
			fmt.Println("Task updated successfully!")
			return 
		}
	}

	fmt.Println("Task not found!")
}

// delete task
func deleteTask(id int) {
    tasks := readTasks() 
    newTasks := []Task{} 

    for _, task := range tasks { 
        if task.ID != id { 
            newTasks = append(newTasks, task) 
        }
    }

    saveTasks(newTasks) 
    fmt.Println("Task deleted successfully!")
}

// lists all tasks
func listTasks() {
	tasks := readTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, task := range tasks {
		fmt.Printf("ID: %d | Description: %s | Status: %s\n", task.ID, task.Description, task.Status)
	}
}

// mark Tasks
func markTask(id int, status string) {
    tasks := readTasks() 

    for i, task := range tasks { 
        if task.ID == id { 
            tasks[i].Status = status 
            tasks[i].UpdatedAt = time.Now().Format(time.RFC3339) 
            saveTasks(tasks) 
            fmt.Printf("Task marked as %s successfully!\n", status)
            return 
        }
    }

    fmt.Println("Task not found!") 
}

// lists staus
func listTasksByStatus(status string) {
    tasks := readTasks()

    filteredTasks := []Task{}
    for _, task := range tasks {
        if task.Status == status {
            filteredTasks = append(filteredTasks, task)
        }
    }

    if len(filteredTasks) == 0 {
        fmt.Printf("No tasks found with status: %s\n", status)
        return
    }

    for _, task := range filteredTasks {
        fmt.Printf("ID: %d | Description: %s | Status: %s\n", task.ID, task.Description, task.Status)
    }
}
