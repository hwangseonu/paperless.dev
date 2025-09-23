package auth

import (
	"errors"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/paperless.dev"
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

func Authorize(c *gin.Context) (int, error) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		return http.StatusUnauthorized, errors.New("no authorization header")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return http.StatusUnauthorized, errors.New("authorization header format must be 'Bearer <token>'")
	}

	token := parts[1]
	claims, err := ParseToken(token)

	if err != nil {
		return http.StatusUnauthorized, paperless.ErrInvalidToken
	}

	if claims.Subject != "access" {
		return http.StatusUnauthorized, paperless.ErrInvalidToken
	}

	c.Set("credential", UserCredentials{
		UserID: claims.UserID,
		Roles:  claims.Roles,
	})

	return http.StatusOK, nil

}

func (p *Protector) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !p.isProtected(c.Request.Method, c.FullPath()) {
			c.Next()
			return
		}

		status, err := Authorize(c)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
