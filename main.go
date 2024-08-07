package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"task-manager/controllers"
	"task-manager/data"
	"task-manager/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGODB_URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	tasks := data.NewTaskManager(client)
	userdata := data.NewUserManager(client)

	taskmgr := controllers.NewTaskController(*tasks, *userdata)

	router := gin.Default()

	routers.SetupRoutes(router, taskmgr)

	router.Run(":8080")
}
