package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
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

	route.Use(cors.Default())

	route.LoadHTMLGlob("./templates/*/*.tmpl")

	userRoutes := route.Group("/user")
	{
		userRoutes.POST("/", app.controller.CreateUser)
		userRoutes.PATCH("/", app.controller.UpdateUser)
		userRoutes.POST("/login", app.controller.LoginUser)
	}

	productRoutes := route.Group("/product")
	{
		productRoutes.GET("/", app.controller.GetManyProduct)
		productRoutes.GET("/:productUUID", app.controller.GetProduct)
		productRoutes.POST("/", app.controller.CreateProduct)
		productRoutes.PATCH("/:productUUID", app.controller.UpdateProduct)
		productRoutes.DELETE("/:productUUID", app.controller.DeleteProduct)
	}

	customerRoutes := route.Group("/customer")
	{
		customerRoutes.POST("/", app.controller.CreateCustomer)
	}

	cartRoutes := route.Group("/cart")
	{
		cartRoutes.POST("/", app.controller.CreateCart)
	}

	orderRoutes := route.Group("/order")
	{
		orderRoutes.POST("/", app.controller.CreateOrder)
		orderRoutes.POST("/decode", app.controller.DecodeEthHash)
	}

	adminRoutes := route.Group("/admin")
	{
		adminRoutes.GET("/", app.controller.HomePage)
		adminRoutes.GET("/product", app.controller.ProductPage)
		adminRoutes.GET("/product/create", app.controller.ProductPageCreate)
		adminRoutes.POST("/product/create", app.controller.ProductPageCreatePost)
		adminRoutes.GET("/order", app.controller.OrderPage)
		adminRoutes.GET("/customer", app.controller.CustomerPage)
	}

	return route
}
