package errors

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

const (
	ErrFetchUser = iota + 1
	ErrUnmarshal
	ErrInvalidData
	ErrInvalidEmail
	ErrMarshalItem
	ErrDeleteItem
	ErrDynamoPutItem
	ErrUserExists
	ErrUserNotFound
	ErrMethodNotAllowed
)

func NewCustomError(code int, message string) error {
	return &CustomError{Code: code, Message: message}
}

// Defined errors
var (
	ErrorFailedToFetchUser       = NewCustomError(ErrFetchUser, "failed to fetch user record")
	ErrorFailedToUnmarshalRecord = NewCustomError(ErrUnmarshal, "failed to unmarshal record")
	ErrorInvalidUserData         = NewCustomError(ErrInvalidData, "invalid user data")
	ErrorInvalidEmail            = NewCustomError(ErrInvalidEmail, "invalid email")
	ErrorCouldNotMarshalItem     = NewCustomError(ErrMarshalItem, "could not marshal item")
	ErrorCouldNotDeleteItem      = NewCustomError(ErrDeleteItem, "could not delete item")
	ErrorCouldNotDynamoPutItem   = NewCustomError(ErrDynamoPutItem, "could not dynamo put item")
	ErrorUserAlreadyExists       = NewCustomError(ErrUserExists, "user already exists")
	ErrorUserDoesNotExist        = NewCustomError(ErrUserNotFound, "user does not exist")
	ErrorMethodNotAllowed        = NewCustomError(ErrMethodNotAllowed, "method not allowed")
)
