package repositories

import (
	"context"
	"errors"
	"log"
	domain "task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewTaskRepository(mongoClient *mongo.Client) domain.TaskRepository {
	return &TaskRepository{
		client:     mongoClient,
		database:   mongoClient.Database("task-manager"),
		collection: mongoClient.Database("task-manager").Collection("tasks"),
	}

}

func (taskrepo *TaskRepository) GetAllTasks(isadmin bool, userid primitive.ObjectID) ([]domain.Task, error) {
	var tasks []domain.Task

	collection := taskrepo.collection
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

func (taskrepo *TaskRepository) GetTask(id string, isadmin bool, userid string) (domain.Task, error) {

	var task domain.Task
	collection := taskrepo.collection
	ido, _ := primitive.ObjectIDFromHex(id)
	userido, _ := primitive.ObjectIDFromHex(userid)

	filter := bson.M{"_id": ido}
	if !isadmin {
		filter = bson.M{"_id": ido, "userid": userido}
	}

	err := collection.FindOne(context.TODO(), filter).Decode(&task)
	return task, err

}

func (taskrepo *TaskRepository) AddTask(task domain.Task) error {
	collection := taskrepo.collection
	task.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), task)

	if err != nil {
		return err
	}

	return nil

}

func (taskrepo *TaskRepository) SetTask(id string, updatedTask domain.Task, isadmin bool) error {
	collection := taskrepo.collection
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

func (taskrepo *TaskRepository) DeleteTask(id string, userid string, isadmin bool) error {
	collection := taskrepo.collection
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
