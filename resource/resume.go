package resource

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/paperless.dev/auth"
	"github.com/hwangseonu/paperless.dev/database"
	"github.com/hwangseonu/paperless.dev/schema"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Resume struct {
	repository     *database.ResumeRepository
	userRepository *database.UserRepository
}

func NewResume() *Resume {
	return &Resume{
		repository:     database.NewResumeRepository(),
		userRepository: database.NewUserRepository(),
	}
}

func (resource *Resume) RequestBody(method string) any {
	switch method {
	case http.MethodPost:
		return new(schema.ResumeCreateSchema)
	default:
		return nil
	}
}

func (resource *Resume) Create(body interface{}, c *gin.Context) (gin.H, int, error) {
	credential := c.MustGet("credential")

	userCred := credential.(auth.Credential)
	userID, err := bson.ObjectIDFromHex(userCred.UserID)

	if err != nil {
		return nil, http.StatusUnauthorized, errors.New("invalid user id in token")
	}

	resume := body.(*schema.ResumeCreateSchema)

	doc, err := resource.repository.InsertOne(database.Resume{
		UserID:   userID,
		Title:    resume.Title,
		Bio:      resume.Bio,
		Public:   resume.Public,
		Template: resume.Template,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("database error")
	}

	res := new(schema.ResumeResponseSchema).FromModel(*doc)
	return gin.H{"resume": res}, http.StatusOK, nil
}

func (resource *Resume) Read(id string, c *gin.Context) (gin.H, int, error) {
	//TODO implement me
	panic("implement me")
}

func (resource *Resume) ReadAll(c *gin.Context) (gin.H, int, error) {
	//TODO implement me
	panic("implement me")
}

func (resource *Resume) Update(id string, body interface{}, c *gin.Context) (gin.H, int, error) {
	//TODO implement me
	panic("implement me")
}

func (resource *Resume) Delete(id string, c *gin.Context) (gin.H, int, error) {
	//TODO implement me
	panic("implement me")
}
