package controller

import (
	"fmt"
	"net/http"

	Product "github.com/faizauthar12/skripsi/product-gomod"
	User "github.com/faizauthar12/skripsi/user-gomod"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateProductHTTPBody struct {
	ProductName        string `json:"productname" binding:"required"`
	ProductDescription string `json:"productdesc" binding:"required"`
	ProductCategory    string `json:"productcategory" binding:"required"`
	ProductPrice       int64  `json:"productprice" binding:"required"`
	ProductStock       int64  `json:"stock" binding:"required"`
}

type UpdateProductHTTPBody struct {
	ProductName        string `json:"productname"`
	ProductDescription string `json:"productdesc"`
	ProductCategory    string `json:"productcategory"`
	ProductPrice       int64  `json:"productprice"`
	ProductStock       int64  `json:"stock"`
}

type ProductController struct {
	Client   *mongo.Client
	UserInfo *User.User
}

func (controller *ProductController) CreateProduct(c *gin.Context) {

	var createProductHTTPBody CreateProductHTTPBody
	errorBodyRequest := c.BindJSON(&createProductHTTPBody)

	if errorBodyRequest != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": 500, "message": errorBodyRequest.Error()})
		return
	}
	product, errorCreateProduct := Product.Create(
		controller.Client,
		controller.UserInfo.UUID,
		createProductHTTPBody.ProductName,
		createProductHTTPBody.ProductDescription,
		createProductHTTPBody.ProductCategory,
		createProductHTTPBody.ProductPrice,
		createProductHTTPBody.ProductStock,
	)

	if errorCreateProduct != nil {

		fmt.Println("CreateProduct() ERR: ", errorCreateProduct.Error())
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_CREATE_PRODUCT,
			},
		)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_CREATE_PRODUCT,
		"data": gin.H{
			"product": product,
		},
	}

	c.JSON(http.StatusOK, successResponse)
}

func (controller *ProductController) GetProduct(c *gin.Context) {

	productUUID := c.Param("productUUID")

	var updateProductHTTPBody UpdateProductHTTPBody
	errorBodyRequest := c.BindJSON(&updateProductHTTPBody)

	if errorBodyRequest != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": 500, "message": errorBodyRequest.Error()})
		return
	}

	var updateList []Product.UpdateCandidate

	if updateProductHTTPBody.ProductName != "" {
		updateList, _ = Product.UpdateName(
			updateList,
			updateProductHTTPBody.ProductName,
		)
	}

	if updateProductHTTPBody.ProductDescription != "" {
		updateList, _ = Product.UpdateDescription(
			updateList,
			updateProductHTTPBody.ProductDescription,
		)
	}

	if updateProductHTTPBody.ProductCategory != "" {
		updateList, _ = Product.UpdateCategory(
			updateList,
			updateProductHTTPBody.ProductCategory,
		)
	}

	if updateProductHTTPBody.ProductPrice > 0 {
		updateList, _ = Product.UpdatePrice(
			updateList,
			updateProductHTTPBody.ProductPrice,
		)
	}

	if updateProductHTTPBody.ProductStock > 0 {
		updateList, _ = Product.UpdateStock(
			updateList,
			updateProductHTTPBody.ProductStock,
		)
	}

	errorUpdateProduct := Product.ExecUpdate(
		controller.Client,
		updateList,
		controller.UserInfo.UUID,
		productUUID,
	)

	if len(updateList) == 0 {
		c.JSON(
			http.StatusOK,
			gin.H{"status": 200, "message": "Nothing to update"},
		)

		return
	}

	if errorUpdateProduct != nil {
		if errorUpdateProduct.Error() == Product.UUID_DOESNT_MATCH {

			errorResponse := gin.H{
				"status":  401,
				"code":    10006,
				"message": UNAUTHORIZED,
			}

			c.JSON(http.StatusUnauthorized,
				errorResponse,
			)

			return
		}

		if errorUpdateProduct.Error() == Product.PRODUCT_NOT_FOUND {

			errorResponse := gin.H{
				"status":  404,
				"code":    10005,
				"message": Product.PRODUCT_NOT_FOUND,
			}

			c.JSON(http.StatusUnauthorized,
				errorResponse,
			)

			return
		}

		fmt.Println("errorUpdateService.Error() ", errorUpdateProduct.Error())

		errorResponse := gin.H{
			"status":  500,
			"code":    10001,
			"message": SERVER_MALFUNCTION_CANNOT_UPDATE_PRODUCT,
		}

		c.JSON(http.StatusInternalServerError, errorResponse)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_UPDATE_PRODUCT,
	}

	c.JSON(http.StatusOK, successResponse)
}

func (controller *ProductController) DeleteProduct(c *gin.Context) {

	productUUID := c.Param("productUUID")

	_, errorDeleteProduct := Product.Delete(
		controller.Client,
		"asdasd",
		productUUID,
	)

	if errorDeleteProduct != nil {

		fmt.Println("DeleteProduct() ERR: ", errorDeleteProduct.Error())

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10010,
				"message": SERVER_MALFUNCTION_CANNOT_DELETE_PRODUCT,
			},
		)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_DELETE_SERVICE,
	}

	c.JSON(http.StatusOK, successResponse)
}
