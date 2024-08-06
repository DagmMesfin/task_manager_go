package controllers

import (
	"net/http"
	"task-manager/data"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service data.UserManager
}

func NewUserController(usermgr data.UserManager) *UserController {
	return &UserController{
		service: usermgr,
	}

}

func (controller *UserController) RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	req_status, err := controller.service.RegisterUserDb(user)

	if err != nil {
		c.JSON(req_status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func (controller *UserController) LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	code, err, token := controller.service.LoginUserDb(user)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully",
		"token": token})
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if code, erro := controller.service.DeleteUser(id); erro == nil {
		c.IndentedJSON(code, gin.H{"message": "user deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
