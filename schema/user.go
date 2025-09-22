package schema

import (
	"time"

	"github.com/hwangseonu/paperless.dev/database"
)

type UserCreateSchema struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserUpdateSchema struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	Name     *string `json:"name,omitempty"`
	Bio      *string `json:"bio,omitempty"`
}

type UserResponseSchema struct {
	ID              string    `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Bio             string    `json:"bio"`
	ProfileImageURL string    `json:"profileImageURL"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func (s *UserResponseSchema) FromModel(user database.User) *UserResponseSchema {
	s.ID = user.ID.Hex()
	s.Username = user.Username
	s.Email = user.Email
	s.Name = user.Name
	s.Bio = user.Bio
	s.ProfileImageURL = user.ProfileImageURL
	s.CreatedAt = user.CreatedAt
	s.UpdatedAt = user.UpdatedAt
	return s
}
