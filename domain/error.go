package domain

import "errors"

// Define unexported error variables
var (
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrInternalServerError   = errors.New("internal server error")
	ErrInvalidInput          = errors.New("invalid input")
	ErrUnauthorized          = errors.New("unauthorized")
	ErrTokenExpired          = errors.New("token expired")
	ErrTokenInvalid          = errors.New("token invalid")
	ErrCodeUserNotFound 	 =errors.New("user not found")
)

// GetStatusCode returns the appropriate HTTP status code based on the error type.
func GetStatusCode(err error) int {
	switch err {
	case ErrUserNotFound:
		return 404
	case ErrInvalidCredentials:
		return 401
	case ErrEmailAlreadyExists:
		return 409
	case ErrUsernameAlreadyExists:
		return 409
	case ErrInvalidInput:
		return 400
	case ErrUnauthorized, ErrTokenExpired, ErrTokenInvalid:
		return 401
	case ErrInternalServerError:
		fallthrough // falls through to default if errInternalServerError is encountered
	default:
		return 500
	}
}
