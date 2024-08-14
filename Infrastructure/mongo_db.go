package infrastructure

import (
	"context"
	"log"
	domain "task-manager/Domain"
	"task-manager/database"
	"time"
)

// initializing the working structure of mongo-db
func MongoDBInit(mongoURI string) domain.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := database.NewClient(mongoURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
