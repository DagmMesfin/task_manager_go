package repositories

import (
	"context"
	"log"
	domain "task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepository struct {
	database   domain.Database
	collection domain.Collection
}

func NewTaskRepository(mongoDatabase domain.Database) domain.TaskRepository {
	return &TaskRepository{
		database:   mongoDatabase,
		collection: mongoDatabase.Collection("tasks"),
	}

}

func (taskrepo *TaskRepository) GetAllTasks(isadmin bool, userid primitive.ObjectID) ([]domain.Task, *domain.AppError) {
	var tasks []domain.Task

	collection := taskrepo.collection
	filter := bson.M{}
	if !isadmin {
		filter = bson.M{"userid": userid}
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, domain.ErrInternalServerError
	}
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		if len(tasks) == 0 {
			return tasks, domain.ErrNoTasksFound
		}
		return nil, domain.ErrInternalServerError
	}
	return tasks, nil
}

func (taskrepo *TaskRepository) GetTask(id string, isadmin bool, userid string) (domain.Task, *domain.AppError) {

	var task domain.Task
	collection := taskrepo.collection
	ido, _ := primitive.ObjectIDFromHex(id)
	userido, _ := primitive.ObjectIDFromHex(userid)

	filter := bson.M{"_id": ido}
	if !isadmin {
		filter = bson.M{"_id": ido, "userid": userido}
	}

	err := collection.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil {
		return *new(domain.Task), domain.ErrTaskNotFound
	}

	return task, nil

}

func (taskrepo *TaskRepository) AddTask(task domain.Task) *domain.AppError {
	collection := taskrepo.collection
	task.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), task)

	if err != nil {
		return domain.ErrTaskInsertionFailed
	}

	return nil

}

func (taskrepo *TaskRepository) SetTask(id string, updatedTask domain.Task, isadmin bool) *domain.AppError {
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
		return domain.ErrTaskUpdateFailed
	}

	return nil
}

func (taskrepo *TaskRepository) DeleteTask(id string, userid string, isadmin bool) *domain.AppError {
	collection := taskrepo.collection
	ido, _ := primitive.ObjectIDFromHex(id)
	userido, _ := primitive.ObjectIDFromHex(userid)
	filter := bson.M{"_id": ido}

	if !isadmin {
		filter = bson.M{"_id": ido, "userid": userido}
	}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return domain.ErrInternalServerError
	}

	if result == 0 {
		return domain.ErrTaskDeletionFailed
	}
	return nil
}
