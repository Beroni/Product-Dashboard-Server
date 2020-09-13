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
	Name     string  `json:"name" binding:"required"`
	Price    float32 `json:"price" binding:"required"`
	Quantity uint8   `json:"quantity" binding:"required"`
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

	product := model.Product{}

	client := *utils.MongoConnection("products")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing values"})
		return
	}

	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	client.InsertOne(context.TODO(), product)

	c.JSON(http.StatusCreated, product)

}

func UpdateProduct(c *gin.Context) {

	product := model.Product{}

	client := *utils.MongoConnection("products")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing values"})
		return
	}

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
		primitive.E{Key: "name", Value: product.Name},
		primitive.E{Key: "price", Value: product.Price},
		primitive.E{Key: "quantity", Value: product.Quantity},
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
