package data

import (
	"errors"
	"task-manager/models"
	"time"
)

type TaskManager struct {
	Tasks  []models.Task
	NextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: []models.Task{
			{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
			{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
			{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
		},
	}

}

func (taskmgr *TaskManager) GetAllTasks() []models.Task {
	return taskmgr.Tasks
}

func (taskmgr *TaskManager) GetTask(id string) (models.Task, error) {
	for _, task := range taskmgr.Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return *new(models.Task), errors.New("not found")
}

func (taskmgr *TaskManager) FindTask(id string) error {
	for _, task := range taskmgr.Tasks {
		if task.ID == id {
			return errors.New("task already exists")
		}
	}
	return nil

}

func (taskmgr *TaskManager) AddTask(task models.Task) {
	taskmgr.Tasks = append(taskmgr.Tasks, task)
}

func (taskmgr *TaskManager) SetTask(id string, updatedTask models.Task) error {
	for i, task := range taskmgr.Tasks {
		if task.ID == id {
			taskmgr.Tasks[i] = updatedTask
			return nil
		}
	}
	return errors.New("task not found")
}

func (taskmgr *TaskManager) DeleteTask(id string) error {
	for i, task := range taskmgr.Tasks {
		if task.ID == id {
			taskmgr.Tasks = append(taskmgr.Tasks[:i], taskmgr.Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
