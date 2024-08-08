package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
=========== The Models and Interfaces for Task =============
*/
type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	UserID      primitive.ObjectID `bson:"userid,omitempty" json:"-"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	DueDate     time.Time          `bson:"due_date,omitempty" json:"due_date,omitempty"`
	Status      string             `bson:"status,omitempty" json:"status,omitempty"`
}

type TaskRepository interface {
	GetAllTasks(isadmin bool, userid primitive.ObjectID) ([]Task, error)
	GetTask(id string, isadmin bool, userid string) (Task, error)
	AddTask(task Task) error
	SetTask(id string, updatedTask Task, isadmin bool) error
	DeleteTask(id string, userid string, isadmin bool) error
}

type TaskUsecase interface {
	GetAllTasks(c context.Context, isadmin bool, userid primitive.ObjectID) ([]Task, error)
	GetTask(c context.Context, id string, isadmin bool, userid string) (Task, error)
	AddTask(c context.Context, task Task) error
	SetTask(c context.Context, id string, updatedTask Task, isadmin bool) error
	DeleteTask(c context.Context, id string, userid string, isadmin bool) error
}

/*
=========== The Models and Interfaces for User =============
*/
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `json:"email"`
	Password string             `json:"password,omitempty"`
	IsAdmin  bool               `json:"isadmin,omitempty"`
}

type UserRepository interface {
	RegisterUserDb(user User) (int, error)
	LoginUserDb(user User) (int, string, error)
	DeleteUser(id string) (int, error)
}

type UserUsecase interface {
	RegisterUserDb(c context.Context, user User) (int, error)
	LoginUserDb(c context.Context, user User) (int, string, error)
	DeleteUser(c context.Context, id string) (int, error)
}
