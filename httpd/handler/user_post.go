package handler

import (
	"cms/platform/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userPostRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"passwword"`
}

func UserPost(users user.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := userPostRequest{}

		c.Bind(&requestBody)

		newUser := user.User{
			Email:    requestBody.Email,
			Name:     requestBody.Name,
			Password: requestBody.Password,
		}

		users.Add(newUser)

		c.Status(http.StatusNoContent)
	}
}
