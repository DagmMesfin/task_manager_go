package domain

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	GetAllTasks(isadmin bool, userid primitive.ObjectID) ([]Task, *AppError)
	GetTask(id string, isadmin bool, userid string) (Task, *AppError)
	AddTask(task Task) *AppError
	SetTask(id string, updatedTask Task, isadmin bool) *AppError
	DeleteTask(id string, userid string, isadmin bool) *AppError
}

type TaskUsecase interface {
	GetAllTasks(c context.Context, isadmin bool, userid primitive.ObjectID) ([]Task, *AppError)
	GetTask(c context.Context, id string, isadmin bool, userid string) (Task, *AppError)
	AddTask(c context.Context, task Task) *AppError
	SetTask(c context.Context, id string, updatedTask Task, isadmin bool) *AppError
	DeleteTask(c context.Context, id string, userid string, isadmin bool) *AppError
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
	RegisterUserDb(user User) *AppError
	LoginUserDb(user User) (string, interface{}, *AppError)
	DeleteUser(id string) *AppError
}

type UserUsecase interface {
	RegisterUserDb(c context.Context, user User) *AppError
	LoginUserDb(c context.Context, user User) (string, interface{}, *AppError)
	DeleteUser(c context.Context, id string) *AppError
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PasswordService interface {
	PasswordComparator(hash string, password string) bool
	PasswordHasher(password string) (string, error)
	TokenGenerator(id primitive.ObjectID, email string, isadmin bool) (string, error)
	TokenClaimer(tokenstr string) (*jwt.Token, error)
}
