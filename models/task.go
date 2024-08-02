package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	DueDate     time.Time          `bson:"due_date,omitempty" json:"due_date,omitempty"`
	Status      string             `bson:"status,omitempty" json:"status,omitempty"`
}
