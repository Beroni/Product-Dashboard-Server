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

	productsRoutes := r.Group("/products")
	{
		productsRoutes.GET("/", middlewares.JWT(), handlers.GetAllProducts)
		productsRoutes.GET("/:id", middlewares.JWT(), handlers.GetProductById)
		productsRoutes.POST("/", middlewares.JWT(), handlers.CreateProduct)
		productsRoutes.PUT("/:id", middlewares.JWT(), handlers.UpdateProduct)
		productsRoutes.DELETE("/:id", middlewares.JWT(), handlers.DeleteProduct)

	}

	r.Run(":3333")
}
