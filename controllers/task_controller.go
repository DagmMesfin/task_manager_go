package controllers

import (
	"net/http"
	"task-manager/data"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service data.TaskManager
}

func NewTaskController(taskmgr data.TaskManager) *TaskController {
	return &TaskController{
		service: taskmgr,
	}

}

func (controller *TaskController) GetTasks(c *gin.Context) {
	tasks := controller.service.GetAllTasks()
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (controller *TaskController) GetTasksById(c *gin.Context) {
	id := c.Param("id")

	task, err := controller.service.GetTask(id)

	if err == nil {
		c.IndentedJSON(http.StatusOK, task)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func (controller *TaskController) PostTask(c *gin.Context) {
	var task models.Task

	err := c.ShouldBindJSON(&task)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if exists := controller.service.FindTask(task.ID); exists != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": exists.Error()})
		return
	}

	controller.service.AddTask(task)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "task created"})
}

func (controller *TaskController) PutTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask models.Task

	err := c.ShouldBindJSON(&updatedTask)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	erro := controller.service.SetTask(id, updatedTask)

	if erro == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task updated"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}
}

func (controller *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if erro := controller.service.DeleteTask(id); erro == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
