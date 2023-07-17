package controller

import (
	"fmt"
	"net/http"
	"strconv"

	token "github.com/faizauthar12/skripsi/backend-service/utils"
	Product "github.com/faizauthar12/skripsi/product-gomod"
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

type GetProductHTTPBody struct {
	UserName string `json:"username"`
}

type UpdateProductHTTPBody struct {
	ProductName        string `json:"productname"`
	ProductDescription string `json:"productdesc"`
	ProductCategory    string `json:"productcategory"`
	ProductPrice       int64  `json:"productprice"`
	ProductStock       int64  `json:"stock"`
}

type ProductController struct {
	Client *mongo.Client
}

func (controller *ProductController) CreateProduct(c *gin.Context) {

	var createProductHTTPBody CreateProductHTTPBody
	errorBodyRequest := c.BindJSON(&createProductHTTPBody)

	user, errorExtractToken := token.ExtractToken(c)

	if errorExtractToken != nil {
		c.JSON(http.StatusUnauthorized,
			gin.H{
				"status":  401,
				"code":    10000, // TODO check code
				"message": UNAUTHORIZED,
			},
		)

		c.Abort()

		return
	}

	if errorBodyRequest != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": 500, "message": errorBodyRequest.Error()})
		return
	}
	product, errorCreateProduct := Product.Create(
		controller.Client,
		user.UUID,
		user.Name,
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

func (controller *ProductController) GetMany(c *gin.Context) {

	// UserName := c.Query("userName")
	numItems, errorParsingNumItems := strconv.ParseInt(c.Query("numItems"), 10, 64)
	pages, errorParsingPages := strconv.ParseInt(c.Query("pages"), 10, 64)

	if errorParsingNumItems != nil {
		numItems = DEFAULT_NUM_ITEMS
	}

	if errorParsingPages != nil {
		pages = DEFAULT_PAGES
	}

	Products, errorGetProducts := Product.GetMany(
		controller.Client,
		numItems,
		pages,
	)

	if errorGetProducts != nil {

		fmt.Println("GetManyBookings() ERR: ", errorGetProducts.Error())
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_GET_PRODUCT,
			},
		)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_GET_PRODUCT,
		"data": gin.H{
			"products": Products,
		},
		"numItems": numItems,
		"pages":    pages,
	}

	c.JSON(http.StatusOK, successResponse)
}

func (controller *ProductController) UpdateProduct(c *gin.Context) {

	productUUID := c.Param("productUUID")

	user, _ := token.ExtractToken(c)

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
		user.UUID,
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

	user, errorExtractToken := token.ExtractToken(c)

	if errorExtractToken != nil {
		c.JSON(http.StatusUnauthorized,
			gin.H{
				"status":  401,
				"code":    10000, // TODO check code
				"message": UNAUTHORIZED,
			},
		)

		c.Abort()

		return
	}

	_, errorDeleteProduct := Product.Delete(
		controller.Client,
		user.UUID,
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
