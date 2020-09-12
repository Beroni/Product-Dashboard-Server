package handlers

import (
	model "cms/src/models"
	utils "cms/src/util"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

const UserCollection = "users"

type sessionPostRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(c *gin.Context) {

	var result model.User

	client := *utils.MongoConnection("users")

	requestBody := sessionPostRequest{}

	loggedUser := model.User{
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	c.Bind(&loggedUser)

	filter := bson.D{{"email", loggedUser.Email}}

	error := client.FindOne(context.TODO(), filter).Decode(&result)

	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Invalid Credentials",
		})
		return
	}

	if result.Password != loggedUser.Password {
		c.JSON(http.StatusOK, gin.H{
			"message": "Invalid Credentials",
		})
		return
	}

	token, error := utils.CreateJWT(result.ID.Hex())

	c.JSON(http.StatusOK, gin.H{
		"user":  result,
		"token": token,
	})

}
