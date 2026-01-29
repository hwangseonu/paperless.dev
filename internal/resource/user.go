package resource

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	restful "github.com/hwangseonu/gin-restful"
	"github.com/hwangseonu/paperless.dev/internal/auth"
	"github.com/hwangseonu/paperless.dev/internal/common"
	"github.com/hwangseonu/paperless.dev/internal/database"
	"github.com/hwangseonu/paperless.dev/internal/schema"
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

// Create *User.Create
// @Summary		create new user
// @Description	create new user
// @Tags	User
// @Accept	json
// @Produce	json
// @Param	user body	schema.UserCreateSchema	true	"initial values of user"
// @Success	201	{object}	object{user=schema.UserResponseSchema}
// @Failure 400 {object}	schema.Error
// @Failure 409 {object}	schema.Error
// @Failure 500 {object}	schema.Error
// @Router	/users [post]
func (resource *User) Create(body interface{}, _ *gin.Context) (gin.H, int, error) {
	user := body.(*schema.UserCreateSchema)

	if doc, err := resource.repository.FindByUsernameOrEmail(user.Username, user.Email); err != nil && !errors.Is(err, common.ErrUserNotFound) {
		return nil, http.StatusInternalServerError, common.ErrDatabase
	} else if doc != nil {
		return nil, http.StatusConflict, common.ErrUserConflict
	}

	var password []byte
	password, _ = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

	result, err := resource.repository.Create(user)

	if err != nil {
		return nil, http.StatusInternalServerError, common.ErrDatabase
	}

	return gin.H{
		"id":       result.ID,
		"username": user.Username,
		"email":    user.Email,
	}, http.StatusCreated, nil
}

// Read *User.Read
// @Summary	get user info by id
// @Description	get user info by id
// @Tags	User
// @Produce	json
// @Param	id	path	string	true	"User ID, Pass 'me' to retrieve your data."
// @Success 200 {object}	object{user=schema.UserResponseSchema}
// @Failure 400 {object} 	schema.Error
// @Failure 404 {object} 	schema.Error
// @Failure 500 {object} 	schema.Error
// @Router	/users/{id} [get]
func (resource *User) Read(id string, c *gin.Context) (gin.H, int, error) {
	if id == "me" {
		credentials := auth.MustGetUserCredentials(c)
		userID := credentials.UserID
		user, err := resource.repository.FindByID(userID)

		if err != nil {
			return nil, http.StatusNotFound, common.ErrUserNotFound
		}

		return gin.H{"user": user.ResponseSchema()}, http.StatusOK, nil
	}

	return nil, http.StatusOK, nil
}

func (resource *User) ReadAll(_ *gin.Context) (gin.H, int, error) {
	return nil, http.StatusNotFound, nil
}

// Update *User.Update
// @Summary	update user data by id
// @Description	update user data by id
// @Tags	User
// @Produce	json
// @Param	id	path	string	true	"User ID, Pass 'me' to delete your data."
// @Param	user body	schema.UserUpdateSchema	true	"update values of User"
// @Success 200 {object}	object{user=schema.UserResponseSchema}
// @Failure 400 {object} 	schema.Error
// @Failure 403 {object} 	schema.Error
// @Failure 404 {object} 	schema.Error
// @Failure 500 {object} 	schema.Error
// @Router	/users/{id} [PATCH]
func (resource *User) Update(id string, body interface{}, c *gin.Context) (gin.H, int, error) {
	if c.Request.Method == http.MethodPut {
		return nil, http.StatusNotFound, nil
	}

	credentials := auth.MustGetUserCredentials(c)

	if id != "me" {
		return nil, http.StatusForbidden, common.ErrAccessDenied
	}

	targetID := credentials.UserID
	updateSchema := body.(*schema.UserUpdateSchema)

	updatedUser, err := resource.repository.Update(targetID, updateSchema)
	if err != nil {
		if errors.Is(err, common.ErrUserNotFound) {
			return nil, http.StatusNotFound, common.ErrUserNotFound
		}
		return nil, http.StatusInternalServerError, common.ErrDatabase
	}

	return gin.H{
		"user": updatedUser.ResponseSchema(),
	}, http.StatusOK, nil
}

// Delete *User.Delete
// @Summary	delete user by id
// @Description	delete user by id
// @Tags	User
// @Produce	json
// @Param	id	path	string	true	"User ID, Pass 'me' to delete your data."
// @Success 204
// @Failure 400 {object} 	schema.Error
// @Failure 403 {object} 	schema.Error
// @Failure 404 {object} 	schema.Error
// @Failure 500 {object} 	schema.Error
// @Router	/users/{id} [DELETE]
func (resource *User) Delete(id string, c *gin.Context) (gin.H, int, error) {
	credentials := auth.MustGetUserCredentials(c)

	var targetID string
	if id == "me" {
		targetID = credentials.UserID
	} else {
		return nil, http.StatusForbidden, common.ErrAccessDenied
	}

	err := resource.repository.DeleteByID(targetID)
	if err != nil {
		return nil, http.StatusInternalServerError, common.ErrDatabase
	}

	return nil, http.StatusNoContent, nil
}
