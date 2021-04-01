package user

import (
	"errors"
)

var (
	ErrUserIsExist  = errors.New("user is exist")
	ErrUserNotFound = errors.New("user not found")
)
