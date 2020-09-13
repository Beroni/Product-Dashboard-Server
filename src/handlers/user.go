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

type SignUpCredentials struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignUp(c *gin.Context) {

	isUserDuplicated := model.User{}
	user := model.User{}

	client := *utils.MongoConnection("users")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := user.ValidateSignUp(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Body"})
		return
	}

	filter := bson.D{{"email", user.Email}}

	error := client.FindOne(context.TODO(), filter).Decode(&isUserDuplicated)

	if error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "E-mail already exists",
		})
		return
	}
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// newUser.Password = utils.HashPassword(requestBody.Password)

	client.InsertOne(context.TODO(), user)

	c.JSON(http.StatusCreated, user)

}
