package routes

import (
	"cms/src/handlers"
	middlewares "cms/src/middleware"

	"github.com/gin-gonic/gin"
)

type Routes struct {
}

func (c Routes) StartGin() {
	r := gin.Default()

	usersRoutes := r.Group("/users")
	{
		usersRoutes.POST("/", handlers.SignUp)
	}

	sessionRoutes := r.Group("/sessions")
	{
		sessionRoutes.POST("/", handlers.SignIn)
	}

	testRoutes := r.Group("/test")
	{
		testRoutes.GET("/", middlewares.JWT(), handlers.Test)
	}

	r.Run(":3333")
}
