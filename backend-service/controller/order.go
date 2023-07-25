package controller

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	Order "github.com/faizauthar12/skripsi/order-gomod"
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

type OrderController struct {
	Client        *mongo.Client
	ClientEth     *ethclient.Client
	EthPrivateKey string
}

func (controller *OrderController) CreateOrder(c *gin.Context) {

	var createOrderHTTPBody CreateOrderHTTPBody
	errorBodyRequest := c.BindJSON(&createOrderHTTPBody)

	if errorBodyRequest != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": 400, "message": errorBodyRequest.Error()})
		return
	}

	order, errorCreateOrder := Order.Create(
		controller.Client,
		createOrderHTTPBody.CartItem,
		createOrderHTTPBody.CartGrandTotal,
		createOrderHTTPBody.CustomerUUID,
		createOrderHTTPBody.CustomerName,
		createOrderHTTPBody.CustomerEmail,
		createOrderHTTPBody.CustomerAddress,
		createOrderHTTPBody.CustomerPhoneNumber,
	)

	if errorCreateOrder != nil {

		fmt.Println("CreateOrder() ERR: ", errorCreateOrder.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status": 500,
				"code":   10000,
				// "message": SERVER_MALFUNCTION_CANNOT_CREATE_PRODUCT,
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
				"status": 500,
				"code":   10000,
				// "message": SERVER_MALFUNCTION_CANNOT_GET_PRODUCT,
			},
		)

		return
	}

	orderContract, errorConnectEth := Order.NewApi(address, controller.ClientEth)
	if errorConnectEth != nil {

		fmt.Println("CreateOrder(): NewApi(): ERR: ", errorConnectEth.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status": 500,
				"code":   10000,
				// "message": SERVER_MALFUNCTION_CANNOT_GET_PRODUCT,
			},
		)

		return
	}

	tx, errorStoreData := Order.StoreDataToEth(order, orderContract, auth)
	if errorStoreData != nil {

		fmt.Println("CreateOrder(): StoreDataToEth(): ERR: ", errorStoreData.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status": 500,
				"code":   10000,
				// "message": SERVER_MALFUNCTION_CANNOT_GET_PRODUCT,
			},
		)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_CREATE_PRODUCT,
		"data": gin.H{
			"order":           order,
			"eth_transaction": tx,
		},
	}

	c.JSON(http.StatusOK, successResponse)
}
