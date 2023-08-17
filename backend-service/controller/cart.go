package controller

import (
	"fmt"
	"net/http"

	Cart "github.com/faizauthar12/skripsi/cart-gomod"
	CartItem "github.com/faizauthar12/skripsi/cart-item-gomod"
	Product "github.com/faizauthar12/skripsi/product-gomod"
	"github.com/gin-gonic/gin"
)

type CreateCartHTTPBody struct {
	CustomerUUID    string `json:"customeruuid" binding:"required"`
	ProductUUID     string `json:"productuuid" binding:"required"`
	ProductQuantity int64  `json:"productquantity" binding:"required"`
}

func (controller *Controller) CreateCart(c *gin.Context) {

	var createCartHTTPBody CreateCartHTTPBody
	errorBodyRequest := c.BindJSON(&createCartHTTPBody)

	if errorBodyRequest != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": 400, "message": errorBodyRequest.Error()})
		return
	}

	product, _, errorGetProduct := Product.Get(
		controller.ClientMongo,
		createCartHTTPBody.ProductUUID,
	)

	if errorGetProduct != nil {
		fmt.Println("CreateCart(): getProduct(): ERR: ", errorGetProduct.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_GET_PRODUCT,
			},
		)

		return
	}

	cartItem, errorCreateCartItem := CartItem.Create(
		controller.ClientMongo,
		createCartHTTPBody.CustomerUUID,
		product.UUID,
		product.ProductPrice,
		createCartHTTPBody.ProductQuantity,
	)

	if errorCreateCartItem != nil {
		fmt.Println("CreateCart(): CreateCartItem(): ERR: ", errorCreateCartItem.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_CREATE_CART_ITEM,
			},
		)

		return
	}

	cartItemMany, errorGetCartItemMany := CartItem.GetManyByCustomerUUID(
		controller.ClientMongo,
		cartItem.CustomerUUID,
		10000,
		1,
	)

	if errorGetCartItemMany != nil {
		fmt.Println("CreateCart(): cartItemMany(): ERR: ", errorGetCartItemMany.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_GET_MANY_CART_ITEM,
			},
		)

		return
	}

	var cartGrandTotal int64
	for _, item := range cartItemMany {
		cartGrandTotal = cartGrandTotal + item.ProductTotalPrice
	}

	cart, errorCreateCart := Cart.Create(
		controller.ClientMongo,
		createCartHTTPBody.CustomerUUID,
		cartGrandTotal,
	)

	if errorCreateCart != nil {
		fmt.Println("CreateCart() ERR: ", errorCreateCart.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_CREATE_CART,
			},
		)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_CREATE_CART,
		"data": gin.H{
			"cart": cart,
		},
	}

	c.JSON(http.StatusOK, successResponse)
}
