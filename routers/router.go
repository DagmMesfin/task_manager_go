package routers

import (
	"task-manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(gino *gin.Engine) {
	gino.GET("/tasks", controllers.GetTasks)
	gino.GET("/tasks/:id", controllers.GetTasksById)
	gino.POST("/tasks", controllers.PostTask)
	gino.PUT("/tasks/:id", controllers.PutTask)
	gino.DELETE("/tasks/:id", controllers.DeleteTask)
}
