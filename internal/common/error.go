package common

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeDatabaseError = 1001
	CodeInvalidInput  = 1002
	CodeUnauthorized  = 1003
	CodeAccessDenied  = 1004

	CodeUserNotFound  = 2001
	CodeInvalidUserID = 2002
	CodeInvalidToken  = 2003
	CodeUserConflict  = 2004

	CodeResumeNotFound  = 3001
	CodeInvalidResumeID = 3002
)

var (
	ErrDatabase     = &Error{"database error", CodeDatabaseError}
	ErrInvalidInput = &Error{"invalid input", CodeInvalidInput}
	ErrUnauthorized = &Error{"unauthorized", CodeUnauthorized}
	ErrAccessDenied = &Error{"access denied", CodeAccessDenied}

	ErrUserNotFound  = &Error{"user not found", CodeUserNotFound}
	ErrInvalidUserID = &Error{"invalid user id", CodeInvalidUserID}
	ErrInvalidToken  = &Error{"invalid token", CodeInvalidToken}
	ErrUserConflict  = &Error{"user conflict", CodeUserConflict}

	ErrResumeNotFound  = &Error{"resume not found", CodeResumeNotFound}
	ErrInvalidResumeID = &Error{"invalid resume id", CodeInvalidResumeID}
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (err *Error) Error() string {
	return err.Message
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	lastErr := c.Errors.Last()

	if lastErr != nil {
		var err *Error
		if errors.As(lastErr.Err, &err) {
			log.Println(err.Message)

			status := http.StatusInternalServerError

			switch err.Code {
			case CodeInvalidInput, CodeInvalidUserID, CodeInvalidResumeID:
				status = http.StatusBadRequest
			case CodeInvalidToken:
				status = http.StatusUnauthorized
			case CodeAccessDenied:
				status = http.StatusForbidden
			case CodeUserNotFound, CodeResumeNotFound:
				status = http.StatusNotFound
			case CodeUserConflict:
				status = http.StatusConflict
			}

			c.AbortWithStatusJSON(status, gin.H{"error": err})
		}
	}
}
