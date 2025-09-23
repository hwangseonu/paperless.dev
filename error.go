package paperless

import "errors"

var (
	ErrDatabase       = errors.New("database error")
	ErrInvalidToken   = errors.New("invalid token")
	ErrUserNotFound   = errors.New("user not found")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrAccessDenied   = errors.New("access denied")
	ErrResumeNotFound = errors.New("resume not found")
	ErrInvalidID      = errors.New("invalid id")
	ErrNoChanges      = errors.New("no fields to update")
)
