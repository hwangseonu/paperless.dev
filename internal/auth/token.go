package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hwangseonu/paperless.dev/internal/common"
)

type Claims struct {
	UserID string   `json:"userid"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

const accessTokenDuration = time.Hour * 24
const refreshTokenDuration = time.Hour * 24 * 30

func GenerateToken(userID string, subject string) (string, error) {
	secret := []byte(common.GetConfig().JwtSecret)

	now := &jwt.NumericDate{Time: time.Now()}
	duration := &jwt.NumericDate{Time: time.Now()}
	if subject == "access" {
		duration.Time = duration.Add(accessTokenDuration)
	} else {
		duration.Time = duration.Add(refreshTokenDuration)
	}

	claims := Claims{
		UserID: userID,
		Roles:  []string{"user"},
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"paperless.dev"},
			ExpiresAt: duration,
			IssuedAt:  now,
			Issuer:    "paperless.dev",
			NotBefore: now,
			Subject:   subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.GetConfig().JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, common.ErrInvalidToken
	}

	return token.Claims.(*Claims), err
}
