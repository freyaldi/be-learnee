package error

import "errors"

var (
	ErrUserAlreadyExists    = errors.New("user already exists")
)