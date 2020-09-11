package handler

import (
	"cms/platform/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserGet(users user.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {

		results := users.GetAll()

		c.JSON(http.StatusOK, results)
	}
}
