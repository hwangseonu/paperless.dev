package resource

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/paperless.dev/internal/auth"
	"github.com/hwangseonu/paperless.dev/internal/common"
	"github.com/hwangseonu/paperless.dev/internal/database"
	"github.com/hwangseonu/paperless.dev/internal/schema"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

// Create *Resume.Create
// @Summary		create new resume
// @Description	create new resume
// @Tags	Resume
// @Accept	json
// @Produce	json
// @Param	resume body	schema.ResumeCreateSchema	true	"initial values of resume"
// @Success	201	{object}	object{resume=schema.ResumeResponseSchema}
// @Failure 400 {object}	schema.Error
// @Failure 500 {object}	schema.Error
// @Router	/resumes [post]
// @Security     BearerAuth
func (resource *Resume) Create(body interface{}, c *gin.Context) (gin.H, int, error) {
	credentials := auth.MustGetUserCredentials(c)

	createSchema := body.(*schema.ResumeCreateSchema)
	createSchema.OwnerID = credentials.UserID

	resume, err := resource.repository.Create(createSchema)

	if err != nil {
		return nil, http.StatusInternalServerError, common.ErrInternal
	}

	res := gin.H{"resume": resume.ResponseSchema()}
	return res, http.StatusCreated, nil
}

// Read *Resume.Read
// @Summary	get resume by id
// @Description	get resume by id
// @Tags	Resume
// @Produce	json
// @Param	id	path	string	true	"Resume ID"
// @Success 200 {object}	object{resume=schema.ResumeResponseSchema}
// @Failure 400 {object} 	schema.Error
// @Failure 403 {object} 	schema.Error
// @Failure 404 {object} 	schema.Error
// @Failure 500 {object} 	schema.Error
// @Router	/resumes/{id} [get]
// @Security BearerAuth
func (resource *Resume) Read(id string, c *gin.Context) (gin.H, int, error) {
	credentials := auth.GetUserCredentials(c)
	userID := ""

	if credentials != nil {
		userID = credentials.UserID
	}

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, http.StatusBadRequest, common.ErrInvalidInput
	}

	resume, err := resource.repository.FindByID(objectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, http.StatusNotFound, common.ErrResumeNotFound
		}
		return nil, http.StatusInternalServerError, common.ErrInternal
	}

	if !resume.Public && resume.OwnerID.Hex() != userID {
		return nil, http.StatusForbidden, common.ErrAccessDenied
	}

	return gin.H{"resume": resume.ResponseSchema()}, http.StatusOK, nil
}

// ReadAll *Resume.ReadAll
// @Summary	get all resumes
// @Description	get all resumes
// @Tags	Resume
// @Produce	json
// @Param	user	query	string	false 	"Owner ID of resumes"
// @Success 200 {object}	object{resumes=[]schema.ResumeResponseSchema}
// @Failure 400 {object} 	schema.Error
// @Failure 403 {object} 	schema.Error
// @Failure 500 {object} 	schema.Error
// @Router	/resumes [get]
// @Security BearerAuth
func (resource *Resume) ReadAll(c *gin.Context) (gin.H, int, error) {
	credentials := auth.GetUserCredentials(c)
	userID := ""
	if credentials != nil {
		userID = credentials.UserID
	}

	targetOwner := c.Query("user")

	if targetOwner == "me" {
		targetOwner = userID
	} else if targetOwner == userID {
		return nil, http.StatusForbidden, common.ErrAccessDenied
	}

	objectID, err := bson.ObjectIDFromHex(targetOwner)
	if err != nil {
		return nil, http.StatusBadRequest, common.ErrInvalidInput
	}

	resumes, err := resource.repository.FindManyByOwnerID(objectID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res := make([]*schema.ResumeResponseSchema, 0)
	for _, resume := range resumes {
		res = append(res, resume.ResponseSchema())
	}

	return gin.H{"resumes": res}, http.StatusOK, nil
}

// Update *Resume.Update
// @Summary	update resume by id
// @Description	update resume by id
// @Tags	Resume
// @Produce	json
// @Param	id	path	string	true	"Resume ID"
// @Param	resume body	schema.ResumeUpdateSchema	true	"update values of resume"
// @Success 200 {object}	object{resume=schema.ResumeResponseSchema}
// @Failure 400 {object} 	schema.Error
// @Failure 403 {object} 	schema.Error
// @Failure 404 {object} 	schema.Error
// @Failure 500 {object} 	schema.Error
// @Router	/resumes/{id} [PATCH]
// @Security     BearerAuth
func (resource *Resume) Update(id string, body interface{}, c *gin.Context) (gin.H, int, error) {
	if c.Request.Method == http.MethodPut {
		return nil, http.StatusNotFound, nil
	}

	credentials := auth.MustGetUserCredentials(c)
	userID := credentials.UserID

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, http.StatusBadRequest, common.ErrInvalidInput
	}
	resumeDoc, err := resource.repository.FindByID(objectID)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, http.StatusNotFound, common.ErrResumeNotFound
		}
		return nil, http.StatusInternalServerError, common.ErrInternal
	}

	if resumeDoc.OwnerID.Hex() != userID {
		return nil, http.StatusForbidden, common.ErrAccessDenied
	}

	updateBody := body.(*schema.ResumeUpdateSchema)
	result, err := resource.repository.Update(objectID, updateBody)

	if err != nil {
		return nil, http.StatusInternalServerError, common.ErrInternal
	}

	return gin.H{"resume": result.ResponseSchema()}, http.StatusOK, nil
}

// Delete *Resume.Delete
// @Summary	delete resume by id
// @Description	delete resume by id
// @Tags	Resume
// @Produce	json
// @Param	id	path	string	true	"Resume ID"
// @Success 204
// @Failure 400 {object} 	schema.Error
// @Failure 403 {object} 	schema.Error
// @Failure 404 {object} 	schema.Error
// @Failure 500 {object} 	schema.Error
// @Router	/resumes/{id} [DELETE]
// @Security     BearerAuth
func (resource *Resume) Delete(id string, c *gin.Context) (gin.H, int, error) {
	credentials := auth.MustGetUserCredentials(c)
	userID := credentials.UserID

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, http.StatusBadRequest, common.ErrInvalidInput
	}
	resumeDoc, err := resource.repository.FindByID(objectID)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, http.StatusNotFound, common.ErrResumeNotFound
		}
		return nil, http.StatusInternalServerError, common.ErrInternal
	}

	if resumeDoc.OwnerID.Hex() != userID {
		return nil, http.StatusForbidden, common.ErrAccessDenied
	}

	_, err = resource.repository.DeleteByID(resumeDoc.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, common.ErrInternal
	}

	return nil, http.StatusNoContent, nil
}
