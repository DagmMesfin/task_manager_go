package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `json:"email"`
	Password string             `json:"password,omitempty"`
	IsAdmin  bool               `json:"isadmin,omitempty"`
}
