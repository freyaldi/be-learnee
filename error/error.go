package error

import "errors"

var (
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrCourseAlreadyFavorited   = errors.New("course already favorited")
	ErrIncorrectCredentials = errors.New("incorrect email or password")
)