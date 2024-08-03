package routers

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(gino *gin.Engine, taskmgr *controllers.TaskController, usermgr *controllers.UserController) {
	gino.GET("/tasks", taskmgr.GetTasks)
	gino.GET("/tasks/:id", taskmgr.GetTasksById)
	gino.POST("/tasks", taskmgr.PostTask)
	gino.PUT("/tasks/:id", taskmgr.PutTask)
	gino.DELETE("/tasks/:id", taskmgr.DeleteTask)

	gino.POST("/register", usermgr.RegisterUser)
	gino.POST("/login", usermgr.LoginUser)

	gino.GET("/secure", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "This is a secure route"})
	})

}
