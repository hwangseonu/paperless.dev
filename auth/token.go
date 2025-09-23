package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hwangseonu/paperless.dev"
)

type Claims struct {
	UserID string   `json:"userid"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}

const accessTokenDuration = time.Hour * 24
const refreshTokenDuration = time.Hour * 24 * 30

func GenerateToken(userID string, subject string) (string, error) {
	secret := []byte(paperless.GetConfig().JwtSecret)

	now := time.Now()
	duration := now
	if subject == "access" {
		duration = duration.Add(accessTokenDuration)
	} else {
		duration = duration.Add(refreshTokenDuration)
	}

	claims := Claims{
		UserID: userID,
		Roles:  []string{"user"},
		StandardClaims: jwt.StandardClaims{
			Audience:  "paperless.dev",
			ExpiresAt: duration.Unix(),
			Id:        "",
			IssuedAt:  now.Unix(),
			Issuer:    "paperless.dev",
			NotBefore: now.Unix(),
			Subject:   subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(paperless.GetConfig().JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, paperless.ErrInvalidToken
	}

	return token.Claims.(*Claims), err
}
