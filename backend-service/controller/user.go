package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

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
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserController struct {
	Client *mongo.Client
}

func (controller *UserController) CreateUser(c *gin.Context) {

	var createUserHTTPBody CreateUserHTTPBody
	errorBodyRequest := c.BindJSON(&createUserHTTPBody)

	if errorBodyRequest != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": 400, "message": errorBodyRequest.Error()})
		return
	}

	client := controller.Client

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

func (controller *UserController) LoginUser(c *gin.Context) {

	var loginUserBodyRequest LoginUserHTTPbody
	errorBodyRequest := c.BindJSON(&loginUserBodyRequest)

	if errorBodyRequest != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": 400, "message": errorBodyRequest.Error()})
		return
	}

	client := controller.Client

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

func (controller *UserController) UpdateUser(c *gin.Context) {

	client := controller.Client

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

	errorUpdateUser := User.ExecUpdate(
		client,
		updateList,
		updateUserBodyRequest.Email,
	)

	if len(updateList) == 0 {
		c.JSON(
			http.StatusOK,
			gin.H{"status": 200, "message": NOTHING_TO_UPDATE},
		)

		return
	}

	if errorUpdateUser != nil {
		if errorUpdateUser.Error() == User.USER_EMAIL_CANNOT_BLANK {
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
