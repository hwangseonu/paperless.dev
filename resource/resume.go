package resource

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/paperless.dev"
	"github.com/hwangseonu/paperless.dev/auth"
	"github.com/hwangseonu/paperless.dev/database"
	"github.com/hwangseonu/paperless.dev/schema"
)

type Resume struct {
	repository     database.ResumeRepository
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
	credentials := auth.MustGetUserCredentials(c)

	createSchema := body.(*schema.ResumeCreateSchema)
	createSchema.OwnerID = credentials.UserID

	resume, err := resource.repository.Create(createSchema)

	if err != nil {
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	return gin.H{"resume": resume.ResponseSchema()}, http.StatusOK, nil
}

func (resource *Resume) Read(id string, c *gin.Context) (gin.H, int, error) {
	credentials := auth.GetUserCredentials(c)
	userID := ""

	if credentials != nil {
		userID = credentials.UserID
	}

	resume, err := resource.repository.FindByID(id)

	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, paperless.ErrResumeNotFound) {
			status = http.StatusNotFound
		}
		return nil, status, err
	}

	if resume.Public || resume.OwnerID.Hex() == userID {
		return gin.H{"resume": resume.ResponseSchema()}, http.StatusOK, nil
	}

	return nil, http.StatusForbidden, paperless.ErrAccessDenied
}

func (resource *Resume) ReadAll(c *gin.Context) (gin.H, int, error) {
	credentials := auth.GetUserCredentials(c)
	userID := ""
	if credentials != nil {
		userID = credentials.UserID
	}

	targetOwner := c.Query("user")

	if targetOwner == userID {
		return nil, http.StatusBadRequest, paperless.ErrAccessDenied
	}

	resumes, err := resource.repository.FindManyByOwnerID(targetOwner)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res := make([]*schema.ResumeResponseSchema, 0)
	for _, resume := range resumes {
		res = append(res, resume.ResponseSchema())
	}

	return gin.H{"resumes": res}, http.StatusOK, nil

}

func (resource *Resume) Update(id string, body interface{}, c *gin.Context) (gin.H, int, error) {
	if c.Request.Method == http.MethodPut {
		return nil, http.StatusNotFound, nil
	}

	credentials := auth.MustGetUserCredentials(c)
	userID := credentials.UserID

	resumeDoc, err := resource.repository.FindByID(id)

	if err != nil {
		if errors.Is(err, paperless.ErrResumeNotFound) {
			return nil, http.StatusNotFound, paperless.ErrResumeNotFound
		}
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	if resumeDoc.OwnerID.Hex() != userID {
		return nil, http.StatusForbidden, paperless.ErrAccessDenied
	}

	updateBody := body.(*schema.ResumeUpdateSchema)
	result, err := resource.repository.Update(id, updateBody)

	if err != nil {
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	return gin.H{"resume": result.ResponseSchema()}, http.StatusOK, nil
}

func (resource *Resume) Delete(id string, c *gin.Context) (gin.H, int, error) {
	credentials := auth.MustGetUserCredentials(c)
	userID := credentials.UserID
	resumeDoc, err := resource.repository.FindByID(id)

	if err != nil {
		if errors.Is(err, paperless.ErrResumeNotFound) {
			return nil, http.StatusNotFound, paperless.ErrResumeNotFound
		}
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	if resumeDoc.OwnerID.Hex() != userID {
		return nil, http.StatusForbidden, paperless.ErrAccessDenied
	}

	err = resource.repository.DeleteByID(id)
	if err != nil {
		return nil, http.StatusInternalServerError, paperless.ErrDatabase
	}

	return nil, http.StatusNoContent, nil
}
