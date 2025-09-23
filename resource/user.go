package resource

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	restful "github.com/hwangseonu/gin-restful"
	"github.com/hwangseonu/paperless.dev"
	"github.com/hwangseonu/paperless.dev/auth"
	"github.com/hwangseonu/paperless.dev/database"
	"github.com/hwangseonu/paperless.dev/schema"
	"go.mongodb.org/mongo-driver/v2/bson"
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

	if doc, err := resource.repository.FindByUsernameOrEmail(user.Username, user.Email); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
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
		userID, err := bson.ObjectIDFromHex(credentials.UserID)

		if err != nil {
			return nil, http.StatusUnauthorized, paperless.ErrInvalidToken
		}

		user, err := resource.repository.FindByID(userID)

		if err != nil {
			return nil, http.StatusNotFound, paperless.ErrUserNotFound
		}

		res := new(schema.UserResponseSchema).FromModel(*user)
		return gin.H{"user": res}, http.StatusOK, nil
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
	var targetUserID bson.ObjectID

	if id == "me" {
		objID, err := bson.ObjectIDFromHex(credentials.UserID)
		if err != nil {
			return nil, http.StatusUnauthorized, paperless.ErrInvalidToken
		}
		targetUserID = objID
	} else {
		// TODO: Admin permission is required to update other users.
		return nil, http.StatusForbidden, paperless.ErrAccessDenied
	}

	updateBody := body.(*schema.UserUpdateSchema)
	updateFields := bson.M{}

	if updateBody.Username != nil {
		updateFields["username"] = *updateBody.Username
	}
	if updateBody.Email != nil {
		updateFields["email"] = *updateBody.Email
	}
	if updateBody.Name != nil {
		updateFields["name"] = *updateBody.Name
	}
	if updateBody.Bio != nil {
		updateFields["bio"] = *updateBody.Bio
	}

	if len(updateFields) == 0 {
		return nil, http.StatusBadRequest, errors.New("no fields to update")
	}

	updateFields["updatedAt"] = time.Now()

	result, err := resource.repository.UpdateOne(targetUserID, updateFields)

	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	if result.MatchedCount == 0 {
		return nil, http.StatusNotFound, paperless.ErrUserNotFound
	}
	if result.ModifiedCount == 0 {
		return nil, http.StatusNotModified, errors.New("no changes made")
	}

	updatedUser, err := resource.repository.FindByID(targetUserID)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("failed to retrieve updated user")
	}

	responseSchema := new(schema.UserResponseSchema).FromModel(*updatedUser)

	return gin.H{
		"user": responseSchema,
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

	userID, err := bson.ObjectIDFromHex(targetID)
	if err != nil {
		return nil, http.StatusUnauthorized, paperless.ErrInvalidToken
	}

	result, err := resource.repository.DeleteByID(userID)
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	if result.DeletedCount == 0 {
		return nil, http.StatusNotFound, paperless.ErrUserNotFound
	}

	return nil, http.StatusNoContent, nil
}
