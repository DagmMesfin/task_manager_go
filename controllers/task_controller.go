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
	service  data.TaskManager
	userdata data.UserManager
}

func NewTaskController(taskmgr data.TaskManager, usermgr data.UserManager) *TaskController {
	return &TaskController{
		service:  taskmgr,
		userdata: usermgr,
	}

}

func (controller *TaskController) GetTasks(c *gin.Context) {
	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	ido, _ := primitive.ObjectIDFromHex(userid)

	tasks, _ := controller.service.GetAllTasks(role, ido)

	log.Println(tasks)

	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (controller *TaskController) GetTasksById(c *gin.Context) {
	id := c.Param("id")

	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	task, err := controller.service.GetTask(id, role, userid)

	log.Println(task)

	if err == nil {
		c.IndentedJSON(http.StatusOK, task)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func (controller *TaskController) PostTask(c *gin.Context) {
	userid := c.GetString("userid")

	log.Println("userid ", userid)

	var task models.Task

	err := c.ShouldBindJSON(&task)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userido, _ := primitive.ObjectIDFromHex(userid)

	task.UserID = userido

	controller.service.AddTask(task)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "task created"})
}

func (controller *TaskController) PutTask(c *gin.Context) {
	id := c.Param("id")

	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	var updatedTask models.Task

	err := c.ShouldBindJSON(&updatedTask)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userido, _ := primitive.ObjectIDFromHex(userid)

	updatedTask.UserID = userido

	log.Println(updatedTask)

	erro := controller.service.SetTask(id, updatedTask, role)

	log.Println(erro)

	if erro == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task updated"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}
}

func (controller *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	if erro := controller.service.DeleteTask(id, userid, role); erro == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func (controller *TaskController) RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	req_status, err := controller.userdata.RegisterUserDb(user)

	if err != nil {
		c.JSON(req_status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func (controller *TaskController) LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	code, err, token := controller.userdata.LoginUserDb(user)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully",
		"token": token})
}

func (controller *TaskController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if code, erro := controller.userdata.DeleteUser(id); erro == nil {
		c.IndentedJSON(code, gin.H{"message": "user deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
