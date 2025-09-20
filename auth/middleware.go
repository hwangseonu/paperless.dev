package auth

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func AccessTokenMiddleware(protected map[string][]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := protected[c.Request.Method]
		if !slices.Contains(p, c.FullPath()) {
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

		c.Set("credential", Credential{
			Username: claims.Username,
			Roles:    claims.Roles,
		})
		c.Next()
	}
}
