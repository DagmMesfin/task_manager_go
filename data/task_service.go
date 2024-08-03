package data

import (
	"context"
	"errors"
	"log"
	"task-manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskManager struct {
	client *mongo.Client
}

func NewTaskManager(mongoClient *mongo.Client) *TaskManager {
	return &TaskManager{
		client: mongoClient,
	}

}

func (taskmgr *TaskManager) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task

	collection := taskmgr.client.Database("task-manager").Collection("tasks")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (taskmgr *TaskManager) GetTask(id string) (models.Task, error) {

	var task models.Task
	collection := taskmgr.client.Database("task-manager").Collection("tasks")
	ido, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(context.TODO(), bson.M{"_id": ido}).Decode(&task)
	return task, err

}

func (taskmgr *TaskManager) AddTask(task models.Task) error {
	collection := taskmgr.client.Database("task-manager").Collection("tasks")
	task.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), task)

	if err != nil {
		return err
	}

	return nil

}

func (taskmgr *TaskManager) SetTask(id string, updatedTask models.Task) error {
	collection := taskmgr.client.Database("task-manager").Collection("tasks")
	ido, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ido}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"status":      updatedTask.Status,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	log.Println(result, err, updatedTask)

	if err != nil || result.MatchedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (taskmgr *TaskManager) DeleteTask(id string) error {
	collection := taskmgr.client.Database("task-manager").Collection("tasks")
	ido, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ido}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil || result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
