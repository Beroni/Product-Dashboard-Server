package handlers

import (
	model "cms/src/models"
	utils "cms/src/util"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRequest struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity uint8   `json:"quantity"`
}

func GetAllProducts(c *gin.Context) {

	products := []*model.Product{}

	client := *utils.MongoConnection("products")

	filter := bson.D{{}}

	cur, _ := client.Find(context.TODO(), filter)

	for cur.Next(context.TODO()) {
		var p model.Product

		err := cur.Decode(&p)
		if err != nil {
			continue
		}

		products = append(products, &p)
	}

	cur.Close(context.TODO())

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})

}

func GetProductById(c *gin.Context) {

	product := model.Product{}

	client := *utils.MongoConnection("products")

	id := c.Param("id")

	objectId, error := primitive.ObjectIDFromHex(id)

	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Invalid ObjectID",
		})
		return
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	error = client.FindOne(context.TODO(), filter).Decode(&product)

	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product not Found",
		})
		return
	}

	c.JSON(http.StatusOK, product)

}

func CreateProduct(c *gin.Context) {

	client := *utils.MongoConnection("products")

	requestBody := productRequest{}

	createdProduct := model.Product{
		Name:     requestBody.Name,
		Price:    requestBody.Price,
		Quantity: requestBody.Quantity,
	}

	c.Bind(&createdProduct)

	createdProduct.ID = primitive.NewObjectID()
	createdProduct.CreatedAt = time.Now()
	createdProduct.UpdatedAt = time.Now()

	client.InsertOne(context.TODO(), createdProduct)

	c.JSON(http.StatusCreated, createdProduct)

}

func UpdateProduct(c *gin.Context) {

	client := *utils.MongoConnection("products")

	product := model.Product{}

	requestBody := productRequest{}

	updatedProduct := model.Product{
		Name:     requestBody.Name,
		Price:    requestBody.Price,
		Quantity: requestBody.Quantity,
	}

	c.Bind(&updatedProduct)

	id := c.Param("id")

	objectId, error := primitive.ObjectIDFromHex(id)

	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Invalid ObjectID",
		})
		return
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "name", Value: updatedProduct.Name},
		primitive.E{Key: "price", Value: updatedProduct.Price},
		primitive.E{Key: "quantity", Value: updatedProduct.Quantity},
		primitive.E{Key: "updated_at", Value: time.Now()},
	}},
	}

	after := options.After

	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	error = client.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&product)

	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product not Found",
		})
		return
	}

	c.JSON(http.StatusOK, product)

}

func DeleteProduct(c *gin.Context) {

	client := *utils.MongoConnection("products")

	id := c.Param("id")

	objectId, error := primitive.ObjectIDFromHex(id)

	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Invalid ObjectID",
		})
		return
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	client.FindOneAndDelete(context.TODO(), filter)

	c.Status(http.StatusNoContent)

}
