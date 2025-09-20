package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user does not exist"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("database error")})
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	access, err1 := GenerateToken(user.Username, "access")
	refresh, err2 := GenerateToken(user.Username, "refresh")

	if err1 != nil || err2 != nil {
		err = errors.Join(err1, err2)
		log.Println("an error accrued while generate tokens", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error accrued while generate tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
