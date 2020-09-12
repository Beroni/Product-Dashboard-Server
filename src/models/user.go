package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User
type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Email     string             `bson:"email"`
	Name      string             `bson:"name"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//Users
type Users []User
