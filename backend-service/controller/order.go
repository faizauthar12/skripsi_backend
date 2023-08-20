package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	Order "github.com/faizauthar12/skripsi/order-gomod"
	Product "github.com/faizauthar12/skripsi/product-gomod"
)

type CreateOrderHTTPBody struct {
	CartItem            []Order.CartItem
	CartGrandTotal      int64
	CustomerUUID        string
	CustomerName        string
	CustomerEmail       string
	CustomerAddress     string
	CustomerPhoneNumber string
	Status              string
}

type DurianPayOrderHTTPBody struct {
	amount                  string
	currency                string
	is_payment_link         bool
	is_live                 bool
	is_notification_enabled bool
	sandbox_options         struct {
		force_fail bool
		delay_ms   int64
	}
	customer struct {
		email          string
		mobile         string
		given_name     string
		address_line_1 string
		city           string
		region         string
		postal_code    string
	}
	items []struct {
		name  string
		qty   int64
		price string
		logo  string
	}
}

func (controller *Controller) CreateOrder(c *gin.Context) {

	var createOrderHTTPBody CreateOrderHTTPBody
	errorBodyRequest := c.BindJSON(&createOrderHTTPBody)

	if errorBodyRequest != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": 400, "message": errorBodyRequest.Error()})
		return
	}

	// Calculate the Product Total Price
	var cartItems []Order.CartItem
	for _, item := range createOrderHTTPBody.CartItem {
		product, _, errorGetProduct := Product.Get(controller.ClientMongo, item.ProductUUID)

		if errorGetProduct != nil {
			fmt.Println("CreateOrder() ERR: ", errorGetProduct.Error())

			c.JSON(http.StatusInternalServerError,
				gin.H{
					"status":  500,
					"code":    10000,
					"message": SERVER_MALFUNCTION_CANNOT_GET_PRODUCT,
				},
			)

			return
		}

		cartItem := Order.CartItem{
			ProductUUID:       item.ProductUUID,
			ProductQuantity:   item.ProductQuantity,
			ProductTotalPrice: product.ProductPrice * item.ProductQuantity,
		}

		cartItems = append(cartItems, cartItem)
	}

	order, errorNewOrder := Order.New(
		cartItems,
		createOrderHTTPBody.CartGrandTotal,
		createOrderHTTPBody.CustomerUUID,
		createOrderHTTPBody.CustomerName,
		createOrderHTTPBody.CustomerEmail,
		createOrderHTTPBody.CustomerAddress,
		createOrderHTTPBody.CustomerPhoneNumber,
	)

	if errorNewOrder != nil {

		fmt.Println("NewOrder() ERR: ", errorNewOrder.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_CREATE_ORDER,
			},
		)

		return
	}

	auth := Order.GetAccountAuth(controller.ClientEth, controller.EthPrivateKey)

	address, errorGetAddress := Order.DeployApi(auth, controller.ClientEth)
	if errorGetAddress != nil {

		fmt.Println("CreateOrder(): DeployApi(): ERR: ", errorGetAddress.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_INITIALIZE_SMART_CONTRACT,
			},
		)

		return
	}

	orderContract, errorConnectEth := Order.NewApi(address, controller.ClientEth)
	if errorConnectEth != nil {

		fmt.Println("CreateOrder(): NewApi(): ERR: ", errorConnectEth.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_STORE_SMART_CONTRACT,
			},
		)

		return
	}

	tx, errorStoreData := Order.StoreDataToEth(order, orderContract, auth)
	if errorStoreData != nil {

		fmt.Println("CreateOrder(): StoreDataToEth(): ERR: ", errorStoreData.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_GET_PRODUCT,
			},
		)

		return
	}

	Order.AddEthAddress(&order, tx)

	errorCreateOrder := Order.Create(controller.ClientMongo, &order)
	if errorCreateOrder != nil {
		fmt.Println("CreateOrder() ERR: ", errorNewOrder.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_CREATE_ORDER,
			},
		)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_CREATE_PRODUCT,
		"data": gin.H{
			"order":            order,
			"eth_transaction":  tx,
			"payment_link_url": "asdasd",
		},
	}

	c.JSON(http.StatusOK, successResponse)
}
