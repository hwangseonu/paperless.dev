package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/paperless.dev/database"
	"github.com/hwangseonu/paperless.dev/schema"
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
	//TODO implement me
	panic("implement me")
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
