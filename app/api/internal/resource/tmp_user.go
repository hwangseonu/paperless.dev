package resource

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/paperless.dev/internal/common"
	"github.com/hwangseonu/paperless.dev/internal/database"
	"github.com/hwangseonu/paperless.dev/internal/mail"
	"github.com/hwangseonu/paperless.dev/internal/redis"
	"github.com/hwangseonu/paperless.dev/internal/schema"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var mailClient = mail.NewClient(common.GetConfig().SMTP)

// CreateTempUserHandler
// @Summary		Initiate user registration & send verification email
// @Description	Starts the sign-up process by creating a temporary user profile in the cache and dispatching a verification code to the provided email address.
// @Tags	User
// @Accept	json
// @Param	user body	schema.UserCreateSchema	true	"initial values of User"
// @Success	201
// @Failure 400 {object}	schema.Error
// @Failure 409 {object}	schema.Error
// @Failure 500 {object}	schema.Error
// @Router	/users/register [post]
func CreateTempUserHandler(userRepo database.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := new(schema.UserCreateSchema)
		if err := c.ShouldBindJSON(body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": common.ErrInvalidInput})
			return
		}

		doc, err := userRepo.FindByEmail(body.Email)
		if err != nil {
			if !errors.Is(err, mongo.ErrNoDocuments) {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal})
				return
			}
		}
		if doc != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": common.ErrUserConflict})
			return
		}

		var password []byte
		password, _ = bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

		tmpUser := redis.TempUser{
			Nickname: body.Nickname,
			Email:    body.Email,
			Password: string(password),
		}

		verifyCode, err := sendVerifyCode(body.Email)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal})
			return
		}

		tmpUser.VerifyCode = verifyCode
		err = redis.SaveTempUser(tmpUser)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": common.ErrInternal})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func sendVerifyCode(email string) (string, error) {
	b := make([]byte, verifyCodeLength/2)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	verifyCode := hex.EncodeToString(b)
	body, err := mail.VerifyCodeTemplate(verifyCode)
	if err != nil {
		return "", err
	}

	content := mail.Content{
		To:      []string{email},
		Subject: "[PAPERLESS.DEV] 인증코드",
		Body:    body,
	}

	if err := mailClient.SendMail(content); err != nil {
		return "", err
	}
	return verifyCode, nil
}
