package auth

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

type Protected struct {
	protected map[string][]string
}

func NewProtected() *Protected {
	return &Protected{protected: make(map[string][]string)}
}

func (p *Protected) isProtected(method, path string) bool {
	paths := p.protected[method]
	return slices.Contains(paths, path)
}

func (p *Protected) Register(path string, methods ...string) {
	for _, method := range methods {
		list, ok := p.protected[method]
		if !ok {
			p.protected[method] = []string{path}
		}
		list = append(list, path)
		p.protected[method] = list
	}
}

func (p *Protected) RegisterAny(path string) {
	p.Register(
		path,
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	)
}

func (p *Protected) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !p.isProtected(c.Request.Method, c.FullPath()) {
			c.Next()
			return
		}

		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is empty"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be 'Bearer <token>'"})
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := ParseToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "an error occurred while parsing token"})
			c.Abort()
			return
		}

		if claims.Subject != "access" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "an error occurred while validating token"})
			c.Abort()
			return
		}

		c.Set("credential", Credential{
			UserID: claims.UserID,
			Roles:  claims.Roles,
		})
		c.Next()
	}
}
