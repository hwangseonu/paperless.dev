package resource

import "github.com/gin-gonic/gin"

type Resume struct {
}

func NewResume() *Resume {
	return &Resume{}
}

func (r *Resume) RequestBody(method string) any {
	//TODO implement me
	panic("implement me")
}

func (r *Resume) Create(body interface{}, c *gin.Context) (gin.H, int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Resume) Read(id string, c *gin.Context) (gin.H, int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Resume) ReadAll(c *gin.Context) (gin.H, int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Resume) Update(id string, body interface{}, c *gin.Context) (gin.H, int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Resume) Delete(id string, c *gin.Context) (gin.H, int, error) {
	//TODO implement me
	panic("implement me")
}
