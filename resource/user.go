package resource

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	restful "github.com/hwangseonu/gin-restful"
	"github.com/hwangseonu/paperless.dev"
	"github.com/hwangseonu/paperless.dev/auth"
	"github.com/hwangseonu/paperless.dev/database"
	"github.com/hwangseonu/paperless.dev/schema"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	restful.Resource
	repository database.UserRepository
}

func NewUser() *User {
	user := new(User)
	user.repository = database.NewUserRepository()
	return user
}

func (resource *User) RequestBody(method string) interface{} {
	switch method {
	case http.MethodPost:
		return new(schema.UserCreateSchema)
	case http.MethodPut, http.MethodPatch:
		return new(schema.UserUpdateSchema)
	default:
		return nil
	}
}

func (resource *User) Create(body interface{}, _ *gin.Context) (gin.H, int, error) {
	user := body.(*schema.UserCreateSchema)

	if doc, err := resource.repository.FindByUsernameOrEmail(user.Username, user.Email); err != nil && !errors.Is(err, paperless.ErrUserNotFound) {
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	} else if doc != nil {
		return nil, http.StatusConflict, paperless.ErrUserConflict
	}

	var password []byte
	password, _ = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

	result, err := resource.repository.Create(user)

	if err != nil {
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	return gin.H{
		"id":       result.ID,
		"username": user.Username,
		"email":    user.Email,
	}, http.StatusCreated, nil
}

func (resource *User) Read(id string, c *gin.Context) (gin.H, int, error) {
	if id == "me" {
		credentials := auth.MustGetUserCredentials(c)
		userID := credentials.UserID
		user, err := resource.repository.FindByID(userID)

		if err != nil {
			return nil, http.StatusNotFound, paperless.ErrUserNotFound
		}

		return gin.H{"user": user.ResponseSchema()}, http.StatusOK, nil
	}

	return nil, http.StatusOK, nil
}

func (resource *User) ReadAll(_ *gin.Context) (gin.H, int, error) {
	return nil, http.StatusNotFound, nil
}

func (resource *User) Update(id string, body interface{}, c *gin.Context) (gin.H, int, error) {
	if c.Request.Method == http.MethodPut {
		return nil, http.StatusNotFound, nil
	}

	credentials := auth.MustGetUserCredentials(c)

	if id != "me" {
		return nil, http.StatusForbidden, paperless.ErrAccessDenied
	}

	targetID := credentials.UserID
	updateSchema := body.(*schema.UserUpdateSchema)

	updatedUser, err := resource.repository.Update(targetID, updateSchema)
	if err != nil {
		if errors.Is(err, paperless.ErrUserNotFound) {
			return nil, http.StatusNotFound, paperless.ErrUserNotFound
		}
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	return gin.H{
		"user": updatedUser.ResponseSchema(),
	}, http.StatusOK, nil
}

func (resource *User) Delete(id string, c *gin.Context) (gin.H, int, error) {
	credentials := auth.MustGetUserCredentials(c)

	var targetID string
	if id == "me" {
		targetID = credentials.UserID
	} else {
		return nil, http.StatusForbidden, paperless.ErrAccessDenied
	}

	err := resource.repository.DeleteByID(targetID)
	if err != nil {
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	return nil, http.StatusNoContent, nil
}
