package controllers

import (
	"net/http"
	domain "task-manager/Domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	taskusecase domain.TaskUsecase
	userusecase domain.UserUsecase
}

// task-controller constructor
func NewTaskController(taskmgr domain.TaskUsecase, usermgr domain.UserUsecase) *TaskController {
	return &TaskController{
		taskusecase: taskmgr,
		userusecase: usermgr,
	}

}

/*
	================== The Task ======================
*/

// Get the tasks
func (controller *TaskController) GetTasks(c *gin.Context) {

	//get the required values from the context
	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	ido, _ := primitive.ObjectIDFromHex(userid)

	tasks, err := controller.taskusecase.GetAllTasks(c, role, ido)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

// Get the task by the id
func (controller *TaskController) GetTasksById(c *gin.Context) {
	id := c.Param("id") //get the path vairable "id"

	//get the required values from the context
	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	task, err := controller.taskusecase.GetTask(c, id, role, userid)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)

}

// Creates the tasks
func (controller *TaskController) PostTask(c *gin.Context) {
	//get the required values from the context
	userid := c.GetString("userid")

	var task domain.Task

	err := c.ShouldBindJSON(&task)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userido, _ := primitive.ObjectIDFromHex(userid)

	task.UserID = userido

	controller.taskusecase.AddTask(c, task)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "task created"})
}

// put the task
func (controller *TaskController) PutTask(c *gin.Context) {
	id := c.Param("id")

	//get the required values from the context
	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	var updatedTask domain.Task

	err := c.ShouldBindJSON(&updatedTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userido, _ := primitive.ObjectIDFromHex(userid)
	updatedTask.UserID = userido

	erro := controller.taskusecase.SetTask(c, id, updatedTask, role)

	if erro != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": erro.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "task updated"})
}

// delete the task
func (controller *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	//get the required values from the context
	role := c.GetBool("isadmin")
	userid := c.GetString("userid")

	erro := controller.taskusecase.DeleteTask(c, id, userid, role)
	if erro == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

/*
	================== The User ======================
*/

// register the user
func (controller *TaskController) RegisterUser(c *gin.Context) {
	var user domain.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	req_status, err := controller.userusecase.RegisterUserDb(c, user)
	if err != nil {
		c.JSON(req_status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
}

// Login the user
func (controller *TaskController) LoginUser(c *gin.Context) {
	var user domain.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	code, token, err := controller.userusecase.LoginUserDb(c, user)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully",
		"token": token})
}

// Delete the user (Admin-specific Operation)
func (controller *TaskController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if code, erro := controller.userusecase.DeleteUser(c, id); erro == nil {
		c.IndentedJSON(code, gin.H{"message": "user deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
