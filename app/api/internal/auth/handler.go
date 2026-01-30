package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/paperless.dev/internal/common"
	"github.com/hwangseonu/paperless.dev/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// LoginHandler
// @Summary		login
// @Description	get tokens
// @Tags	Auth
// @Accept	json
// @Produce	json
// @Param	credentials body	LoginCredentials	true	"login credentials info"
// @Success	200	{object}	LoginResponse
// @Failure 400 {object}	schema.Error
// @Failure 401 {object}	schema.Error
// @Failure 404 {object}	schema.Error
// @Failure 500 {object}	schema.Error
// @Router	/auth/login [post]
func LoginHandler(c *gin.Context) {
	var credentials LoginCredentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := database.NewUserRepository().FindByUsername(credentials.Username)

	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, common.ErrUserNotFound) {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": common.ErrUnauthorized})
		return
	}

	access, err1 := GenerateToken(user.ID.Hex(), "access")
	refresh, err2 := GenerateToken(user.ID.Hex(), "refresh")

	if err1 != nil || err2 != nil {
		err = errors.Join(err1, err2)
		log.Println("an error occurred while generate tokens", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ErrInvalidToken})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
