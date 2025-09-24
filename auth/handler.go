package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/paperless.dev"
	"github.com/hwangseonu/paperless.dev/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := database.NewUserRepository().FindByUsername(credentials.Username)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": paperless.ErrUserNotFound})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": paperless.ErrDatabase})
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": paperless.ErrUnauthorized})
		return
	}

	access, err1 := GenerateToken(user.ID.Hex(), "access")
	refresh, err2 := GenerateToken(user.ID.Hex(), "refresh")

	if err1 != nil || err2 != nil {
		err = errors.Join(err1, err2)
		log.Println("an error occurred while generate tokens", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": paperless.ErrInvalidToken})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
