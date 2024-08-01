package routers

import (
	"task-manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(gino *gin.Engine, taskmgr *controllers.TaskController) {
	gino.GET("/tasks", taskmgr.GetTasks)
	gino.GET("/tasks/:id", taskmgr.GetTasksById)
	gino.POST("/tasks", taskmgr.PostTask)
	gino.PUT("/tasks/:id", taskmgr.PutTask)
	gino.DELETE("/tasks/:id", taskmgr.DeleteTask)
}
