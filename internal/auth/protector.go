package auth

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/paperless.dev/internal/common"
)

type Protector struct {
	protected map[string][]string
}

func NewProtector() *Protector {
	return &Protector{protected: make(map[string][]string)}
}

func (p *Protector) isProtected(method, path string) bool {
	paths := p.protected[method]
	return slices.Contains(paths, path)
}

func (p *Protector) Register(path string, methods ...string) {
	for _, method := range methods {
		p.protected[method] = append(p.protected[method], path)
	}
}

func (p *Protector) RegisterAny(path string) {
	p.Register(
		path,
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	)
}

func Authorize(c *gin.Context) (*UserCredentials, error) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		return nil, common.ErrUnauthorized
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, common.ErrInvalidInput
	}

	token := parts[1]
	claims, err := ParseToken(token)

	if err != nil {
		return nil, common.ErrInvalidToken
	}

	if claims.Subject != "access" {
		return nil, common.ErrInvalidToken
	}

	return &UserCredentials{
		UserID: claims.UserID,
		Roles:  claims.Roles,
	}, nil

}

func (p *Protector) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !p.isProtected(c.Request.Method, c.FullPath()) {
			c.Next()
			return
		}

		credentials, err := Authorize(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("credential", credentials)
		c.Next()
	}
}
