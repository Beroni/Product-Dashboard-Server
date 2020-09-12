package model

import (
	"time"
)

//User
type User struct {
	ID        string    `json:"_id" bson:"_id"`
	Email     string    `json:"email" bson:"email"`
	Name      string    `json:"name" bson:"name"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"created_at" bson:"updated_at"`
}

//Users
type Users []User
