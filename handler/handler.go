package handler

import "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
	courseUsecase usecase.CourseUsecase
}

type Config struct {
	UserUsecase usecase.UserUsecase
	CourseUsecase usecase.CourseUsecase
}

func New(c *Config) *Handler {
	return &Handler{
		userUsecase: c.UserUsecase,
		courseUsecase: c.CourseUsecase,
	}
}