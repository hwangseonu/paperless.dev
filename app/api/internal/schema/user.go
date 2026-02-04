package schema

import (
	"time"
)

type UserCreateSchema struct {
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserUpdateSchema struct {
	Nickname     *string `json:"nickname,omitempty"`
	ProfileImage *string `json:"profile_image,omitempty"`
}

type UserResponseSchema struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
