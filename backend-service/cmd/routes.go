package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "Hello world",
		})
	})

	userRoutes := route.Group("/user")
	{
		userRoutes.POST("/", app.user.CreateUser)
		userRoutes.PATCH("/", app.user.UpdateUser)

		userRoutes.POST("/login", app.user.LoginUser)
	}

	productRoutes := route.Group("/product")
	{
		productRoutes.POST("/", app.product.CreateProduct)
		productRoutes.GET("/:productUUID", app.product.GetProduct)
		productRoutes.DELETE("/:productUUID", app.product.DeleteProduct)
	}

	return route
}
