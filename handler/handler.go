package handler

import "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
}

type Config struct {
	UserUsecase usecase.UserUsecase
}

func New(c *Config) *Handler {
	return &Handler{
		userUsecase: c.UserUsecase,
	}
}