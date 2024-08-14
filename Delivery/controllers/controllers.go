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
		c.JSON(err.Status(), gin.H{"error": err.Message()})
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
		c.IndentedJSON(err.Status(), gin.H{"error": err.Message()})
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

	erro := controller.taskusecase.AddTask(c, task)

	if erro != nil {
		c.IndentedJSON(erro.Status(), gin.H{"error": erro.Message()})
		return
	}

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
		c.IndentedJSON(erro.Status(), gin.H{"error": erro.Message()})
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

	c.IndentedJSON(erro.Status(), gin.H{"error": erro.Message()})
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

	erro := controller.userusecase.RegisterUserDb(c, user)
	if erro != nil {
		c.JSON(erro.Status(), gin.H{"error": erro.Message()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login the user
func (controller *TaskController) LoginUser(c *gin.Context) {
	var user domain.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	token, result, erro := controller.userusecase.LoginUserDb(c, user)

	if erro != nil {
		c.JSON(erro.Status(), gin.H{"error": erro.Message()})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully",
		"token": token, "user": result})
}

// Delete the user (Admin-specific Operation)
func (controller *TaskController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if erro := controller.userusecase.DeleteUser(c, id); erro != nil {
		c.IndentedJSON(erro.Status(), gin.H{"error": erro.Message()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted"})

}
