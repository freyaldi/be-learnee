package error

import "errors"

var (
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrIncorrectCredentials = errors.New("incorrect email or password")
)