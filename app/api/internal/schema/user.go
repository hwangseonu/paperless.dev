package schema

import (
	"time"
)

type UserCreateSchema struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserUpdateSchema struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
}

type UserResponseSchema struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
