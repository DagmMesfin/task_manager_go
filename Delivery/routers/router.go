package routers

import (
	"task-manager/Delivery/controllers"
	infrastructure "task-manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(gino *gin.Engine, taskmgr *controllers.TaskController) {

	//public routes
	gino.POST("/register", taskmgr.RegisterUser)
	gino.POST("/login", taskmgr.LoginUser)

	//protected routes group
	auth := gino.Group("/")
	auth.Use(infrastructure.AuthMiddleware)
	{
		auth.GET("/tasks", taskmgr.GetTasks)          //get all tasks
		auth.GET("/tasks/:id", taskmgr.GetTasksById)  //get a specific task
		auth.PUT("/tasks/:id", taskmgr.PutTask)       //update a task
		auth.POST("/tasks", taskmgr.PostTask)         //create a task
		auth.DELETE("/tasks/:id", taskmgr.DeleteTask) //delete a task

		//admin-specific endpoint group
		admin := auth.Group("/")
		admin.Use(infrastructure.AdminMiddleware)
		{
			//delete user
			admin.DELETE("/users/:id", taskmgr.DeleteUser)
		}
	}

}
