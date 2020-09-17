package routes

import (
	"cms/src/handlers"
	middlewares "cms/src/middleware"

	"github.com/gin-gonic/gin"
)

type Routes struct {
}

func (c Routes) StartGin() *gin.Engine {
	r := gin.New()

	r.Use(middlewares.CORS())

	r.POST("/users", handlers.SignUp)

	r.POST("/sessions", handlers.SignIn)

	r.GET("/products", middlewares.JWT(), handlers.GetAllProducts)
	r.GET("/products/:id", middlewares.JWT(), handlers.GetProductById)
	r.POST("/products", middlewares.JWT(), handlers.CreateProduct)
	r.PUT("/products/:id", middlewares.JWT(), handlers.UpdateProduct)
	r.DELETE("/products/:id", middlewares.JWT(), handlers.DeleteProduct)

	return r
}
