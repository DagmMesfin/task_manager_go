package routers

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(gino *gin.Engine, taskmgr *controllers.TaskController, usermgr *controllers.UserController) {

	auth := gino.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/tasks", taskmgr.GetTasks)
		auth.GET("/tasks/:id", taskmgr.GetTasksById)

		admin := auth.Group("/")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.POST("/tasks", taskmgr.PostTask)
			admin.PUT("/tasks/:id", taskmgr.PutTask)
			admin.DELETE("/tasks/:id", taskmgr.DeleteTask)
		}
	}
	gino.POST("/register", usermgr.RegisterUser)
	gino.POST("/login", usermgr.LoginUser)

}
