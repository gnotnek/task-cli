package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"taskQ/internal/model"
	"time"
)

const taskFile = "tasks.json"

// loadTasks reads tasks from the JSON file
func loadTasks() ([]model.Task, error) {
	file, err := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []model.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return tasks, nil
}

// saveTasks writes tasks to the JSON file
func saveTasks(tasks []model.Task) error {
	file, err := os.Create(taskFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}

// AddTask adds a new task to the task list
func AddTask(description, status string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	task := model.Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now().Format(time.DateTime),
		UpdatedAt:   time.Now().Format(time.DateTime),
	}

	tasks = append(tasks, task)
	return saveTasks(tasks)
}

// GetTask retrieves a task by its ID
func GetTask(id string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println(err)
		return
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, task := range tasks {
		if task.ID == ID {
			fmt.Println(task)
			return
		}
	}

	fmt.Println("task not found")
}

// UpdateTask updates an existing task
func UpdateTask(id string, description, status string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == ID {
			task = tasks[i]
			if description == "" {
				description = task.Description
			}

			if status == "" {
				status = task.Status
			}

			updatedTask := model.Task{
				ID:          task.ID,
				Description: description,
				Status:      status,
				CreatedAt:   task.CreatedAt,
				UpdatedAt:   time.Now().Format(time.DateTime),
			}
			tasks[i] = updatedTask
			return saveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

// DeleteTask removes a task by its ID
func DeleteTask(id string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return saveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

// ListTasks returns all tasks
func ListTasks() {
	task, err := loadTasks()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(task)
}
