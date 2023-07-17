package middlewares

import (
	"fmt"
	"net/http"

	"github.com/faizauthar12/skripsi/backend-service/controller"
	User "github.com/faizauthar12/skripsi/user-gomod"
	"github.com/gin-gonic/gin"
)

type Middlewares struct {
	UserInfo *User.User
}

func (middlewares *Middlewares) ExtractTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		user, errorExtractToken := User.ExtractToken(bearerToken)

		fmt.Println("ExtractToken: bearerToken: ", bearerToken)
		fmt.Println("ExtractToken: user: ", user)

		if errorExtractToken != nil {
			c.JSON(http.StatusUnauthorized,
				gin.H{
					"status":  401,
					"code":    10000, // TODO check code
					"message": controller.UNAUTHORIZED,
				},
			)

			c.Abort()

			return
		}

		middlewares.UserInfo = &User.User{
			UUID:  user.UUID,
			Name:  user.Name,
			Email: user.Email,
		}
	}
}
