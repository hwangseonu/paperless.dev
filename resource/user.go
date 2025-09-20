package resource

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	restful "github.com/hwangseonu/gin-restful"
	"github.com/hwangseonu/paperless.dev/database"
	"github.com/hwangseonu/paperless.dev/schema"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	restful.Resource
	repository *database.UserRepository
}

func NewUser() *User {
	user := new(User)
	user.repository = database.NewUserRepository()
	return user
}

func (resource *User) RequestBody(_ string) interface{} {
	return new(schema.User)
}

func (resource *User) Create(body interface{}, _ *gin.Context) (gin.H, int, error) {
	user := body.(*schema.User)

	if doc, err := resource.repository.FindByUsernameOrEmail(user.Username, user.Email); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, http.StatusInternalServerError, errors.New("database error")
	} else if doc != nil {
		return nil, http.StatusConflict, errors.New("user already exists")
	}

	var password []byte
	password, _ = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	doc := database.User{
		Username:  user.Username,
		Password:  string(password),
		Email:     user.Email,
		CreatedAt: time.Now(),
	}

	result, err := resource.repository.InsertOne(doc)

	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError, errors.New("database error")
	}

	return gin.H{
		"id":       result.ID,
		"username": user.Username,
		"email":    user.Email,
	}, http.StatusCreated, nil
}

func (resource *User) Read(_ string, _ *gin.Context) (gin.H, int, error) {
	return nil, http.StatusOK, nil
}

func (resource *User) ReadAll(_ *gin.Context) (gin.H, int, error) {
	return nil, http.StatusNotFound, nil
}

func (resource *User) Update(_ string, _ interface{}, _ *gin.Context) (gin.H, int, error) {
	return nil, http.StatusNoContent, nil
}

func (resource *User) Delete(_ string, _ *gin.Context) (gin.H, int, error) {
	return nil, http.StatusNoContent, nil

}
