package main

import (
	"cms/httpd/handler"
	"cms/platform/user"

	"github.com/gin-gonic/gin"
)

func main() {
	users := user.New()

	r := gin.Default()

	r.GET("/users", handler.UserGet(users))
	r.POST("/users", handler.UserPost(users))

	r.Run()

}
