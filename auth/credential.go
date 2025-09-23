package auth

import (
	"github.com/gin-gonic/gin"
)

type UserCredentials struct {
	UserID string
	Roles  []string
}

func GetUserCredentials(c *gin.Context) *UserCredentials {
	cred, ok := c.Get("credential")
	if !ok {
		return nil
	}

	credential := cred.(UserCredentials)
	return &credential
}

func MustGetUserCredentials(c *gin.Context) *UserCredentials {
	cred := c.MustGet("credential")
	credential := cred.(UserCredentials)
	return &credential
}
