package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/paperless.dev/internal/common"
	"github.com/hwangseonu/paperless.dev/internal/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
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
// @Success	200	{object}	TokenResponse
// @Failure 400 {object}	schema.Error
// @Failure 401 {object}	schema.Error
// @Failure 404 {object}	schema.Error
// @Failure 500 {object}	schema.Error
// @Router	/auth/login [post]
func LoginHandler(c *gin.Context) {
	var credentials LoginCredentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": common.ErrInvalidInput})
		return
	}

	user, err := database.NewUserRepository().FindByEmail(credentials.Email)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": common.ErrUserNotFound})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal})
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": common.ErrUnauthorized})
		return
	}

	access, err1 := GenerateToken(user.Email, "access")
	refresh, err2 := GenerateToken(user.Email, "refresh")

	if err1 != nil || err2 != nil {
		err = errors.Join(err1, err2)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	})
}

// RefreshHandler
// @Summary    refresh tokens
// @Description get new access and refresh tokens using refresh token
// @Tags    Auth
// @Accept  json
// @Produce json
// @Param   Authorization header string true "Bearer {refresh_token}"
// @Success 200    {object}   TokenResponse
// @Failure 401 {object}    schema.Error
// @Failure 500 {object}    schema.Error
// @Router  /auth/refresh [post]
func RefreshHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": common.ErrUnauthorized})
		return
	}

	tokenString := authHeader
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	}

	claims, err := ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrTokenInvalid})
		return
	}

	if claims.Subject != "refresh" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrTokenInvalid})
		return
	}

	access, err1 := GenerateToken(claims.UserID, "access")
	refresh, err2 := GenerateToken(claims.UserID, "refresh")

	if err1 != nil || err2 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": common.ErrInvalidToken})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	})
}
