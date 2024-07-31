package controllers

import (
	"net/http"
	"task-manager/data"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks := data.GetAllTasks()
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTasksById(c *gin.Context) {
	id := c.Param("id")

	task, err := data.GetTask(id)

	if err == nil {
		c.JSON(http.StatusOK, task)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func PostTask(c *gin.Context) {
	var task models.Task

	err := c.ShouldBindJSON(&task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.AddTask(task)

	c.JSON(http.StatusCreated, gin.H{"message": "task created"})
}

func PutTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask models.Task

	err := c.ShouldBindJSON(&updatedTask)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	erro := data.SetTask(id, updatedTask)

	if erro == nil {
		c.JSON(http.StatusOK, gin.H{"message": "task updated"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if erro := data.DeleteTask(id); erro == nil {
		c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
