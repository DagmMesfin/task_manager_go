package main

import (
	"task-manager/Delivery/controllers"
	"task-manager/Delivery/routers"
	infrastructure "task-manager/Infrastructure"
	repositories "task-manager/Repositories"
	usecase "task-manager/Usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	client := infrastructure.MongoDBInit() //mongodb initialization

	//initialization of the repositories
	task_repo := repositories.NewTaskRepository(client)
	user_repo := repositories.NewUserRepository(client)

	//set-up the controllers
	cont := controllers.NewTaskController(usecase.NewTaskUsecase(task_repo, 3*time.Minute), usecase.NewUserUsecase(user_repo, 3*time.Minute))

	//the router gateway
	router := gin.Default()
	routers.SetupRoutes(router, cont)
	router.Run(":8080")
}
