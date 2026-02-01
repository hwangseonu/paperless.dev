package resource

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	restful "github.com/hwangseonu/gin-restful"
	"github.com/hwangseonu/paperless.dev/internal/auth"
	"github.com/hwangseonu/paperless.dev/internal/common"
	"github.com/hwangseonu/paperless.dev/internal/database"
	"github.com/hwangseonu/paperless.dev/internal/redis"
	"github.com/hwangseonu/paperless.dev/internal/schema"
	redis2 "github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const verifyCodeLength = 6

type User struct {
	restful.Resource
	repository database.UserRepository
}

func NewUser(userRepo database.UserRepository) *User {
	user := new(User)
	user.repository = userRepo
	return user
}

func (resource *User) RequestBody(method string) interface{} {
	switch method {
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
// @Produce	json
// @Param	code	query	string	true	"verify code"
// @Success	204	{object}	object{user=schema.UserResponseSchema}
// @Failure 400 {object}	schema.Error
// @Failure 409 {object}	schema.Error
// @Failure 500 {object}	schema.Error
// @Router	/users [post]
func (resource *User) Create(_ interface{}, c *gin.Context) (gin.H, int, error) {
	code := c.Query("code")

	tmp, err := redis.GetTempUser(code)
	if errors.Is(err, redis2.Nil) {
		return nil, http.StatusBadRequest, common.ErrInvalidVerifyCode
	} else if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	doc, err := resource.repository.FindByEmail(tmp.Email)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, http.StatusInternalServerError, common.ErrInternal
		}
	}

	if doc != nil {
		return nil, http.StatusConflict, common.ErrUserConflict
	}

	result, err := resource.repository.Create(&schema.UserCreateSchema{
		Nickname: tmp.Nickname,
		Email:    tmp.Email,
		Password: tmp.Password,
	})

	if err != nil {
		return nil, http.StatusInternalServerError, common.ErrInternal
	}

	user := result.ResponseSchema()
	return gin.H{"user": user}, http.StatusCreated, nil
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
// @Security     BearerAuth
func (resource *User) Read(id string, c *gin.Context) (gin.H, int, error) {
	if id == "me" {
		credentials := auth.MustGetUserCredentials(c)
		userID := credentials.UserID

		objectID, err := bson.ObjectIDFromHex(userID)
		if err != nil {
			return nil, http.StatusBadRequest, common.ErrInvalidInput
		}

		user, err := resource.repository.FindByID(objectID)

		if err != nil {
			return nil, http.StatusNotFound, common.ErrUserNotFound
		}

		return gin.H{"user": user.ResponseSchema()}, http.StatusOK, nil
	}

	return nil, http.StatusForbidden, nil
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
// @Security     BearerAuth
func (resource *User) Update(id string, body interface{}, c *gin.Context) (gin.H, int, error) {
	if c.Request.Method == http.MethodPut {
		return nil, http.StatusNotFound, nil
	}

	credentials := auth.MustGetUserCredentials(c)

	if id != "me" {
		return nil, http.StatusForbidden, common.ErrAccessDenied
	}

	targetID := credentials.UserID
	objectID, err := bson.ObjectIDFromHex(targetID)
	if err != nil {
		return nil, http.StatusBadRequest, common.ErrInvalidInput
	}

	updateSchema := body.(*schema.UserUpdateSchema)
	updatedUser, err := resource.repository.Update(objectID, updateSchema)
	if err != nil {
		if errors.Is(err, common.ErrUserNotFound) {
			return nil, http.StatusNotFound, common.ErrUserNotFound
		}
		return nil, http.StatusInternalServerError, common.ErrInternal
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
// @Security     BearerAuth
func (resource *User) Delete(id string, c *gin.Context) (gin.H, int, error) {
	credentials := auth.MustGetUserCredentials(c)

	var targetID string
	if id == "me" {
		targetID = credentials.UserID
	} else {
		return nil, http.StatusForbidden, common.ErrAccessDenied
	}

	objectID, err := bson.ObjectIDFromHex(targetID)
	if err != nil {
		return nil, http.StatusBadRequest, common.ErrInvalidInput
	}
	_, err = resource.repository.DeleteByID(objectID)
	if err != nil {
		return nil, http.StatusInternalServerError, common.ErrInternal
	}

	return nil, http.StatusNoContent, nil
}
