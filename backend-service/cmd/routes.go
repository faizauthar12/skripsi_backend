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
		productRoutes.GET("/", app.product.GetMany)
		productRoutes.POST("/", app.product.CreateProduct)
		productRoutes.PATCH("/:productUUID", app.product.UpdateProduct)
		productRoutes.DELETE("/:productUUID", app.product.DeleteProduct)
	}

	customerRoutes := route.Group("/customer")
	{
		customerRoutes.POST("/", app.customer.CreateCustomer)
	}

	cartRoutes := route.Group("/cart")
	{
		cartRoutes.POST("/", app.cart.CreateCart)
	}

	orderRoutes := route.Group("/order")
	{
		orderRoutes.POST("/", app.order.CreateOrder)
	}

	return route
}
