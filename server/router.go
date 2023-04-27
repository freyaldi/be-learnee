package server

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/handler"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/middleware"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/usecase"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase     usecase.UserUsecase
	CourseUsecase   usecase.CourseUsecase
	CategoryUsecase usecase.CategoryUsecase
	TagUsecase      usecase.TagUsecase
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handler.New(&handler.Config{
		UserUsecase:     c.UserUsecase,
		CourseUsecase:   c.CourseUsecase,
		CategoryUsecase: c.CategoryUsecase,
		TagUsecase:      c.TagUsecase,
	})

	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	router.GET("/categories", h.Categories)
	router.GET("/tags", h.Tags)

	router.Use(middleware.AuthorizeJWT(c.UserUsecase))
	{
		router.GET("/profile", h.Profile)
		router.GET("/courses/:slug", h.GetCourse)

		router.POST("/courses/create", h.CreateCourse)
		router.POST("/courses/update/:id", h.UpdateCourse)
		router.POST("/courses/delete", h.DeleteCourse)
	}

	return router
}
