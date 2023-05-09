package error

import "errors"

var (
	ErrUserAlreadyExists                  = errors.New("user already exists")
	ErrIncorrectCredentials               = errors.New("incorrect email or password")
	ErrCourseAlreadyCarted                = errors.New("course already carted")
	ErrCourseAlreadyFavorited             = errors.New("course already favorited")
	ErrCourseAlreadyUnFavorited           = errors.New("course already unfavorited")
	ErrCourseNotFound                     = errors.New("course is not found")
	ErrCartIsEmpty                        = errors.New("cart is empty")
	ErrTransactionStatusAlreadyAsExpected = errors.New("transaction status already as expected")
)
