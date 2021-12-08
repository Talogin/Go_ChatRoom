package model

import (
	"errors"
)

var (
	ErrUserNotExist    = errors.New("User not exist...")
	ErrInvalidPassword = errors.New("Password or username not right...")
	ErrInvalidParams   = errors.New("Invalid params")
	ErrUserExist       = errors.New("user exist")
)
