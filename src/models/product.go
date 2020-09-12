package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Product
type Product struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Price     float32            `json:"price" bson:"price"`
	Quantity  uint8              `json:"quantity" bson:"quantity"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

//Products
type Products []Product
