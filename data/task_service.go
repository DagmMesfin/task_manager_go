package data

import (
	"errors"
	"task-manager/models"
	"time"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func GetAllTasks() []models.Task {
	return tasks
}

func GetTask(id string) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return *new(models.Task), errors.New("not found")
}

func FindTask(id string) error {
	for _, task := range tasks {
		if task.ID == id {
			return errors.New("task already exists")
		}
	}
	return nil

}

func AddTask(task models.Task) {
	tasks = append(tasks, task)
}

func SetTask(id string, updatedTask models.Task) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			return nil
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
