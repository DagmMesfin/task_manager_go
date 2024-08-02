package controllers

import (
	"log"
	"net/http"
	"task-manager/data"
	"task-manager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	tasks, _ := controller.service.GetAllTasks()

	log.Println(tasks)

	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (controller *TaskController) GetTasksById(c *gin.Context) {
	id := c.Param("id")

	ido, _ := primitive.ObjectIDFromHex(id)

	task, err := controller.service.GetTask(ido)

	log.Println(task)

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

	task.ID = primitive.NewObjectID()

	if _, exists := controller.service.GetTask(task.ID); exists == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Task Already Exists"})
		return
	}

	controller.service.AddTask(task)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "task created"})
}

func (controller *TaskController) PutTask(c *gin.Context) {
	id := c.Param("id")

	ido, _ := primitive.ObjectIDFromHex(id)

	var updatedTask models.Task

	err := c.ShouldBindJSON(&updatedTask)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(updatedTask)

	erro := controller.service.SetTask(ido, updatedTask)

	log.Println(erro)

	if erro == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task updated"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}
}

func (controller *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	ido, _ := primitive.ObjectIDFromHex(id)

	if erro := controller.service.DeleteTask(ido); erro == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
