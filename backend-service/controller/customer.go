package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	Customer "github.com/faizauthar12/skripsi/customer-gomod"
)

type CreateCustomerHTTPBody struct {
	CustomerName        string `json:"customername"`
	CustomerEmail       string `json:"customeremail"`
	CustomerAddress     string `json:"customeraddress"`
	CustomerPhoneNumber string `json:"phonenumber"`
}

type CustomerController struct {
	Client *mongo.Client
}

func (controller *CustomerController) CreateCustomer(c *gin.Context) {

	var createCustomerHTTPBody CreateCustomerHTTPBody
	errorBodyRequest := c.BindJSON(&createCustomerHTTPBody)

	if errorBodyRequest != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": 400, "message": errorBodyRequest.Error()})
		return
	}

	customer, errorCreateCustomer := Customer.Create(
		controller.Client,
		createCustomerHTTPBody.CustomerName,
		createCustomerHTTPBody.CustomerEmail,
		createCustomerHTTPBody.CustomerAddress,
		createCustomerHTTPBody.CustomerPhoneNumber,
	)

	if errorCreateCustomer != nil {
		fmt.Println("CreateCustomer() ERR: ", errorCreateCustomer.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status": 500,
				"code":   10000,
				// "message": SERVER_MALFUNCTION_CANNOT_CREATE_PRODUCT,
			},
		)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_CREATE_CUSTOMER,
		"data": gin.H{
			"customer": customer,
		},
	}

	c.JSON(http.StatusOK, successResponse)
}
