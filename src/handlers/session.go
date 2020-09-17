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

type SignInCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignIn(c *gin.Context) {

	findedUser := model.User{}
	user := model.User{}

	client := *utils.MongoConnection("users")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing values"})
		return
	}

	if err := user.ValidateSignIn(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Body"})
		return
	}

	filter := bson.D{{"email", user.Email}}

	error := client.FindOne(context.TODO(), filter).Decode(&findedUser)

	if error != nil || findedUser.Password != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Credentials",
		})
		return
	}

	token, error := utils.CreateJWT(findedUser.ID.Hex())

	c.JSON(http.StatusOK, gin.H{
		"user":  findedUser,
		"token": token,
	})

}
