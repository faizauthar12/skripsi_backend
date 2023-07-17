package token

import (
	"fmt"

	User "github.com/faizauthar12/skripsi/user-gomod"
	"github.com/gin-gonic/gin"
)

func ExtractToken(c *gin.Context) (User.User, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	user, errorExtractToken := User.ExtractToken(bearerToken)

	fmt.Println("ExtractToken: bearerToken: ", bearerToken)
	fmt.Println("ExtractToken: user: ", user)

	if errorExtractToken != nil {
		return User.User{}, errorExtractToken
	}

	return user, nil
}
