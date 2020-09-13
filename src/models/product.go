package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Product -> Product Model
type Product struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Price     float32            `json:"price" bson:"price"`
	Quantity  uint8              `json:"quantity" bson:"quantity"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

//Products -> Array of Products
type Products []Product

//ValidateCreation -> Validator to the post request.
func (product Product) ValidateCreation() error {

	return validation.ValidateStruct(&product,
		validation.Field(&product.Name, validation.Required),
		validation.Field(&product.Price, validation.Required, is.Float),
		validation.Field(&product.Quantity, validation.Required, is.Int),
	)

}

//ValidateUpdate -> Validator to the put request.
func (product Product) ValidateUpdate() error {

	return validation.ValidateStruct(&product,
		validation.Field(&product.Name),
		validation.Field(&product.Price, is.Float),
		validation.Field(&product.Quantity, is.Int),
	)

}
