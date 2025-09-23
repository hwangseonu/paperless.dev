package resource

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/paperless.dev"
	"github.com/hwangseonu/paperless.dev/auth"
	"github.com/hwangseonu/paperless.dev/database"
	"github.com/hwangseonu/paperless.dev/schema"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
		return nil, http.StatusUnauthorized, paperless.ErrInvalidToken
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
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	res := new(schema.ResumeResponseSchema).FromModel(*doc)
	return gin.H{"resume": res}, http.StatusOK, nil
}

func (resource *Resume) Read(id string, c *gin.Context) (gin.H, int, error) {
	status, err := auth.Authorize(c)
	authorized := err == nil && status == http.StatusOK

	var userID string
	if authorized {
		credential := c.MustGet("credential")
		userCred := credential.(auth.Credential)
		userID = userCred.UserID
	}

	docID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, http.StatusBadRequest, paperless.ErrInvalidID
	}
	resume, err := resource.repository.FindByID(docID)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, http.StatusNotFound, paperless.ErrResumeNotFound
		}
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	if resume.Public || resume.UserID.Hex() == userID {
		res := new(schema.ResumeResponseSchema).FromModel(*resume)
		return gin.H{"resume": res}, http.StatusOK, nil
	}

	return nil, http.StatusForbidden, paperless.ErrAccessDenied
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
