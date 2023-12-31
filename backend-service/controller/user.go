package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/faizauthar12/skripsi/backend-service/utils"
	User "github.com/faizauthar12/skripsi/user-gomod"
)

type CreateUserHTTPBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserHTTPbody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserHTTPBody struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
}

func (controller *Controller) CreateUser(c *gin.Context) {

	var createUserHTTPBody CreateUserHTTPBody
	errorBodyRequest := c.BindJSON(&createUserHTTPBody)

	if errorBodyRequest != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": 400, "message": errorBodyRequest.Error()})
		return
	}

	client := controller.ClientMongo

	userUUID, errorCreateUser := User.Create(
		client,
		createUserHTTPBody.Name,
		createUserHTTPBody.Email,
		createUserHTTPBody.Password,
	)

	if errorCreateUser != nil {

		fmt.Println("CreateUser() ERR: ", errorCreateUser.Error())
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_CREATE_USER,
			},
		)

		return
	}

	userPayload := User.User{
		UUID:  userUUID,
		Name:  createUserHTTPBody.Password,
		Email: createUserHTTPBody.Email,
	}

	userToken, errorGenerateToken := User.GenerateToken(
		userPayload,
		"",
		false,
	)

	if errorGenerateToken != nil {
		fmt.Println("GenerateToken() ERR: ", errorGenerateToken.Error())
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_CREATE_TOKEN,
			},
		)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_CREATE_USER,
		"data": gin.H{
			"userUUID": userUUID,
			"token":    userToken,
		},
	}

	c.JSON(http.StatusOK, successResponse)
}

func (controller *Controller) LoginUser(c *gin.Context) {

	var loginUserBodyRequest LoginUserHTTPbody
	errorBodyRequest := c.BindJSON(&loginUserBodyRequest)

	if errorBodyRequest != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": 400, "message": errorBodyRequest.Error()})
		return
	}

	client := controller.ClientMongo

	logedUser, isAuthenticated := User.NativeAuthenticate(
		client,
		loginUserBodyRequest.Email,
		loginUserBodyRequest.Password,
	)

	if !isAuthenticated {

		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  404,
				"code":    10000, // check
				"message": USER_NOT_FOUND,
			},
		)

		return
	}

	userPayload := User.User{
		UUID:  logedUser.UUID,
		Name:  logedUser.Name,
		Email: logedUser.Email,
	}

	userToken, errorGenerateToken := User.GenerateToken(
		userPayload,
		"",
		false,
	)

	if errorGenerateToken != nil {
		fmt.Println("GenerateToken() ERR: ", errorGenerateToken.Error())
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  500,
				"code":    10000,
				"message": SERVER_MALFUNCTION_CANNOT_CREATE_TOKEN,
			},
		)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_LOGIN_USER,
		"data": gin.H{
			"user":  logedUser,
			"token": userToken,
		},
	}

	c.JSON(http.StatusOK, successResponse)
}

func (controller *Controller) UpdateUser(c *gin.Context) {

	client := controller.ClientMongo

	user, errorExtractToken := utils.ExtractToken(c)

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

	var updateUserBodyRequest UpdateUserHTTPBody
	errorBodyRequest := c.BindJSON(&updateUserBodyRequest)

	if errorBodyRequest != nil {
		errorResponse := gin.H{
			"status":  400,
			"message": errorBodyRequest.Error(),
		}
		c.JSON(http.StatusBadRequest, errorResponse)

		return
	}

	var updateList []User.UpdateCandidate

	if updateUserBodyRequest.Name != "" {
		updateList, _ = User.UpdateName(
			updateList,
			updateUserBodyRequest.Name,
		)
	}

	if updateUserBodyRequest.Email != "" {
		updateList, _ = User.UpdateEmail(
			updateList,
			updateUserBodyRequest.Email,
		)
	}

	if updateUserBodyRequest.Password != "" {
		updateList, _ = User.UpdatePassword(
			updateList,
			updateUserBodyRequest.Password,
		)
	}

	if updateUserBodyRequest.Address != "" {
		updateList, _ = User.UpdateAddress(
			updateList,
			updateUserBodyRequest.Address,
		)
	}

	if updateUserBodyRequest.PhoneNumber != "" {
		updateList, _ = User.UpdatePhoneNumber(
			updateList,
			updateUserBodyRequest.PhoneNumber,
		)
	}

	errorUpdateUser := User.ExecUpdate(
		client,
		updateList,
		user.UUID,
	)

	if len(updateList) == 0 {
		c.JSON(
			http.StatusOK,
			gin.H{"status": 200, "message": NOTHING_TO_UPDATE},
		)

		return
	}

	if errorUpdateUser != nil {
		if errorUpdateUser.Error() == User.USER_UUID_CANNOT_BLANK {
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

		fmt.Println("errorUpdateUser.Error() ", errorUpdateUser.Error())

		errorResponse := gin.H{
			"status":  500,
			"code":    10001,
			"message": SERVER_MALFUNCTION_CANNOT_UPDATE_USER,
		}

		c.JSON(http.StatusInternalServerError, errorResponse)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_UPDATE_USER,
	}

	c.JSON(http.StatusOK, successResponse)
}

func (controller *Controller) EnableMerchant(c *gin.Context) {

	client := controller.ClientMongo

	user, errorExtractToken := utils.ExtractToken(c)

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

	var updateList []User.UpdateCandidate

	updateList, _ = User.EnableMerchant(
		client,
		updateList,
	)

	errorUpdateUser := User.ExecUpdate(
		client,
		updateList,
		user.UUID,
	)

	if errorUpdateUser != nil {
		if errorUpdateUser.Error() == User.USER_UUID_CANNOT_BLANK {
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

		fmt.Println("errorUpdateUser.Error() ", errorUpdateUser.Error())

		errorResponse := gin.H{
			"status":  500,
			"code":    10001,
			"message": SERVER_MALFUNCTION_CANNOT_UPDATE_USER,
		}

		c.JSON(http.StatusInternalServerError, errorResponse)

		return
	}

	successResponse := gin.H{
		"status":  200,
		"message": SUCCESS_ENABLE_MERCHANT,
	}

	c.JSON(http.StatusOK, successResponse)
}
