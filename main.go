package main

import (
	"task-manager/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.SetupRoutes(router)
	router.Run()
}
