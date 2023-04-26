package error

import "errors"

var (
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrUserNotFound         = errors.New("user not found")
	ErrIncorrectCredentials = errors.New("incorrect email or password")
)