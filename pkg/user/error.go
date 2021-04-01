package user

import (
	"errors"
)

// errors
var (
	ErrUserIsExist  = errors.New("user is exist")
	ErrUserNotFound = errors.New("user not found")

	ErrFiledAuthenticate = errors.New("failed authenticate")
	ErrTokenExpired      = errors.New("token expired")
)
