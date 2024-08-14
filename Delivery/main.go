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
	mongoURI := infrastructure.DotEnvLoader("MONGODB_URI") //extracts the secret URI from the .env
	client := infrastructure.MongoDBInit(mongoURI)         //mongodb initialization
	db := client.Database("task-manager")
	//initialization of the repositories
	pass_service := infrastructure.NewPasswordService()
	task_repo := repositories.NewTaskRepository(db)
	user_repo := repositories.NewUserRepository(db, pass_service)

	//set-up the controllers
	cont := controllers.NewTaskController(usecase.NewTaskUsecase(task_repo, 3*time.Minute), usecase.NewUserUsecase(user_repo, 3*time.Minute))

	//the router gateway
	router := gin.Default()
	routers.SetupRoutes(router, cont)
	router.Run(":8080")
}
