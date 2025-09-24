package paperless

const (
	CodeDatabaseError  = 1001
	CodeInvalidInput   = 1002
	CodeUnauthorized   = 1003
	CodeAccessDenied   = 1004
	CodeEntityConflict = 1005

	CodeUserNotFound  = 2001
	CodeInvalidUserID = 2002
	CodeInvalidToken  = 2003

	CodeResumeNotFound  = 3001
	CodeInvalidResumeID = 3002
)

var (
	ErrDatabase       = NewError("database error", CodeDatabaseError)
	ErrInvalidInput   = NewError("invalid input", CodeInvalidInput)
	ErrUnauthorized   = NewError("unauthorized", CodeUnauthorized)
	ErrAccessDenied   = NewError("access denied", CodeAccessDenied)
	ErrEntityConflict = NewError("entity conflict", CodeEntityConflict)

	ErrUserNotFound  = NewError("user not found", CodeUserNotFound)
	ErrInvalidUserID = NewError("invalid user id", CodeInvalidUserID)
	ErrInvalidToken  = NewError("invalid token", CodeInvalidToken)

	ErrResumeNotFound  = NewError("resume not found", CodeResumeNotFound)
	ErrInvalidResumeID = NewError("invalid resume id", CodeInvalidResumeID)
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (err *Error) Error() string {
	return err.Message
}

func NewError(message string, code int) error {
	return &Error{message, code}
}
