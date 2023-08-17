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
		userRoutes.POST("/", app.user.CreateUser)
		userRoutes.PATCH("/", app.user.UpdateUser)
		userRoutes.POST("/login", app.user.LoginUser)
	}

	productRoutes := route.Group("/product")
	{
		productRoutes.GET("/", app.product.GetMany)
		productRoutes.GET("/:productUUID", app.product.Get)
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

	adminRoutes := route.Group("/admin")
	{
		adminRoutes.GET("/", app.admin.HomePage)
		adminRoutes.GET("/product", app.admin.ProductPage)
		adminRoutes.GET("/order", app.admin.OrderPage)
		adminRoutes.GET("/customer", app.admin.CustomerPage)
	}

	return route
}
