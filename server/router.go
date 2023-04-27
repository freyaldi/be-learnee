package server

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/handler"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/middleware"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/usecase"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase usecase.UserUsecase
	CourseUsecase usecase.CourseUsecase
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handler.New(&handler.Config{
		UserUsecase: c.UserUsecase,
		CourseUsecase: c.CourseUsecase,
	})

	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	router.Use(middleware.AuthorizeJWT(c.UserUsecase))
	{
		router.GET("/profile", h.Profile)
		router.GET("/courses/:slug", h.GetCourse)
	}

	return router
}
