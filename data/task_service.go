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

func (taskmgr *TaskManager) GetAllTasks(isadmin bool, userid primitive.ObjectID) ([]models.Task, error) {
	var tasks []models.Task

	collection := taskmgr.client.Database("task-manager").Collection("tasks")
	filter := bson.M{}
	if !isadmin {
		filter = bson.M{"userid": userid}
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		if len(tasks) == 0 {
			return tasks, errors.New("no tasks found")
		}
		return nil, err
	}
	return tasks, nil
}

func (taskmgr *TaskManager) GetTask(id string, isadmin bool, userid string) (models.Task, error) {

	var task models.Task
	collection := taskmgr.client.Database("task-manager").Collection("tasks")
	ido, _ := primitive.ObjectIDFromHex(id)
	userido, _ := primitive.ObjectIDFromHex(userid)

	filter := bson.M{"_id": ido}

	if !isadmin {
		filter = bson.M{"_id": ido, "userid": userido}
	}

	err := collection.FindOne(context.TODO(), filter).Decode(&task)
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

func (taskmgr *TaskManager) SetTask(id string, updatedTask models.Task, isadmin bool) error {
	collection := taskmgr.client.Database("task-manager").Collection("tasks")
	ido, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": ido}

	if !isadmin {
		filter = bson.M{"_id": ido, "userid": updatedTask.UserID}
	}

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
		return errors.New("task not found or you don't have the privilege to edit it")
	}

	return nil
}

func (taskmgr *TaskManager) DeleteTask(id string, userid string, isadmin bool) error {
	collection := taskmgr.client.Database("task-manager").Collection("tasks")
	ido, _ := primitive.ObjectIDFromHex(id)
	userido, _ := primitive.ObjectIDFromHex(userid)
	filter := bson.M{"_id": ido}

	if !isadmin {
		filter = bson.M{"_id": ido, "userid": userido}
	}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil || result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
