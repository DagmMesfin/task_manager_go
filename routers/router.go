package routers

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(gino *gin.Engine, taskmgr *controllers.TaskController) {

	gino.POST("/register", taskmgr.RegisterUser)
	gino.POST("/login", taskmgr.LoginUser)

	auth := gino.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/tasks", taskmgr.GetTasks)
		auth.GET("/tasks/:id", taskmgr.GetTasksById)
		auth.PUT("/tasks/:id", taskmgr.PutTask)
		auth.POST("/tasks", taskmgr.PostTask)
		auth.DELETE("/tasks/:id", taskmgr.DeleteTask)

		admin := auth.Group("/")
		admin.Use(middleware.AdminMiddleware())
		{
			//delete user
			admin.DELETE("/users/:id", taskmgr.DeleteUser)
		}
	}

}
