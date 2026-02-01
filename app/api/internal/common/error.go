package common

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ErrorCode int

// Server Error
const (
	CodeInternalError = ErrorCode(1000 + iota)
	CodeInvalidInput
)

// Auth Error
const (
	CodeUnauthorized = ErrorCode(2000 + iota)
	CodeAccessDenied
	CodeInvalidToken
	CodeExpiredToken
)

// User Error
const (
	CodeUserNotFound = ErrorCode(3000 + iota)
	CodeUserConflict
	CodeInvalidVerifyCode
)

// Resume Error
const (
	CodeResumeNotFound = ErrorCode(4000 + iota)
)

var errorMessages = map[ErrorCode]string{
	CodeInternalError: "Internal server error",
	CodeInvalidInput:  "Invalid input",

	CodeUnauthorized: "Unauthorized",
	CodeAccessDenied: "Access denied",
	CodeInvalidToken: "Token is invalid",
	CodeExpiredToken: "Token is expired",

	CodeUserNotFound:      "User not found",
	CodeUserConflict:      "User conflict",
	CodeInvalidVerifyCode: "Invalid verify code",

	CodeResumeNotFound: "Resume not found",
}

func NewError(code ErrorCode) *Error {
	return &Error{
		Message: errorMessages[code],
		Code:    code,
	}
}

var (
	ErrInternal     = NewError(CodeInternalError)
	ErrInvalidInput = NewError(CodeInvalidInput)
)

var (
	ErrUnauthorized = NewError(CodeUnauthorized)
	ErrAccessDenied = NewError(CodeAccessDenied)
	ErrInvalidToken = NewError(CodeInvalidToken)
)

var (
	ErrUserNotFound      = NewError(CodeUserNotFound)
	ErrUserConflict      = NewError(CodeUserConflict)
	ErrInvalidVerifyCode = NewError(CodeInvalidVerifyCode)
)

var (
	ErrResumeNotFound = NewError(CodeResumeNotFound)
)

type Error struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
}

func (err *Error) Error() string {
	return err.Message
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		for _, err := range c.Errors {
			log.Printf("[ERROR_LOG] error: %v\n", err.Err)
		}
	}
}
