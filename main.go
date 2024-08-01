package main

import (
	"task-manager/controllers"
	"task-manager/data"
	"task-manager/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	tasks := data.NewTaskManager()
	taskmgr := controllers.NewTaskController(*tasks)
	router := gin.Default()

	routers.SetupRoutes(router, taskmgr)

	router.Run()
}
