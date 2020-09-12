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
)

type userPostRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func SignUp(c *gin.Context) {

	var result model.User

	client := *utils.MongoConnection("users")

	requestBody := userPostRequest{}

	newUser := model.User{
		Email:    requestBody.Email,
		Name:     requestBody.Name,
		Password: requestBody.Password,
	}

	c.Bind(&newUser)

	filter := bson.D{{"email", newUser.Email}}

	error := client.FindOne(context.TODO(), filter).Decode(&result)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "E-mail already exists",
		})
		return
	}
	newUser.ID = primitive.NewObjectID()
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	// newUser.Password = utils.HashPassword(requestBody.Password)

	client.InsertOne(context.TODO(), newUser)

	c.JSON(http.StatusCreated, newUser)

}
