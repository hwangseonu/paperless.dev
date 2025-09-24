package resource

import (
	"errors"
	"net/http"
	"time"

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
	userRepository database.UserRepository
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
	case http.MethodPatch:
		return new(schema.ResumeUpdateSchema)
	default:
		return nil
	}
}

func (resource *Resume) Create(body interface{}, c *gin.Context) (gin.H, int, error) {
	credential := c.MustGet("credential")

	userCred := credential.(auth.UserCredentials)
	userID, err := bson.ObjectIDFromHex(userCred.UserID)

	if err != nil {
		return nil, http.StatusUnauthorized, paperless.ErrInvalidToken
	}

	resume := body.(*schema.ResumeCreateSchema)

	doc, err := resource.repository.InsertOne(database.Resume{
		UserID:    userID,
		Title:     resume.Title,
		Bio:       resume.Bio,
		Public:    resume.Public,
		Template:  resume.Template,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	return gin.H{"resume": doc.FromModel()}, http.StatusOK, nil
}

func (resource *Resume) Read(id string, c *gin.Context) (gin.H, int, error) {
	credentials := auth.GetUserCredentials(c)
	userID := ""

	if credentials != nil {
		userID = credentials.UserID
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
		return gin.H{"resume": resume.FromModel()}, http.StatusOK, nil
	}

	return nil, http.StatusForbidden, paperless.ErrAccessDenied
}

func (resource *Resume) ReadAll(_ *gin.Context) (gin.H, int, error) {
	//TODO implement me
	panic("implement me")
}

func (resource *Resume) Update(id string, body interface{}, c *gin.Context) (gin.H, int, error) {
	if c.Request.Method == http.MethodPut {
		return nil, http.StatusNotFound, nil
	}

	credentialContext, ok := c.Get("credential")
	if !ok {
		return nil, http.StatusUnauthorized, paperless.ErrUnauthorized
	}
	credential := credentialContext.(auth.UserCredentials)
	ownerID, _ := bson.ObjectIDFromHex(credential.UserID)

	resumeID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, http.StatusNotFound, paperless.ErrInvalidID
	}

	resumeDoc, err := resource.repository.FindByID(resumeID)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, http.StatusNotFound, paperless.ErrResumeNotFound
		}
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	if resumeDoc.UserID != ownerID {
		return nil, http.StatusForbidden, paperless.ErrAccessDenied
	}

	updateBody := body.(*schema.ResumeUpdateSchema)

	updateFields := bson.M{}
	if updateBody.Title != nil {
		updateFields["title"] = *updateBody.Title
	}
	if updateBody.Bio != nil {
		updateFields["bio"] = *updateBody.Bio
	}
	if updateBody.Public != nil {
		updateFields["public"] = *updateBody.Public
	}
	if updateBody.Template != nil {
		updateFields["template"] = *updateBody.Template
	}
	if updateBody.Skills != nil {
		updateFields["skills"] = *updateBody.Skills
	}
	if updateBody.Experiences != nil {
		updateFields["experiences"] = *updateBody.Experiences
	}
	if updateBody.Educations != nil {
		updateFields["educations"] = *updateBody.Educations
	}
	if updateBody.Projects != nil {
		updateFields["projects"] = *updateBody.Projects
	}

	if len(updateFields) == 0 {
		return nil, http.StatusBadRequest, paperless.ErrNoChanges
	}

	updateFields["updatedAt"] = time.Now()

	result, err := resource.repository.UpdateOne(resumeID, updateFields)
	if err != nil {
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	if result.MatchedCount == 0 {
		return nil, http.StatusNotFound, paperless.ErrResumeNotFound
	}
	if result.ModifiedCount == 0 {
		return nil, http.StatusNotModified, paperless.ErrNoChanges
	}

	return nil, http.StatusNoContent, nil
}

func (resource *Resume) Delete(_ string, _ *gin.Context) (gin.H, int, error) {
	//TODO implement me
	panic("implement me")
}
