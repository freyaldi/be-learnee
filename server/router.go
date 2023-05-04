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
	FavoriteUsecase usecase.FavoriteUsecase
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handler.New(&handler.Config{
		UserUsecase:     c.UserUsecase,
		CourseUsecase:   c.CourseUsecase,
		CategoryUsecase: c.CategoryUsecase,
		TagUsecase:      c.TagUsecase,
		FavoriteUsecase: c.FavoriteUsecase,
	})

	api := router.Group("/api/v1")

	api.POST("/register", h.Register)
	api.POST("/login", h.Login)

	api.GET("/categories", h.Categories)
	api.GET("/tags", h.Tags)
	api.GET("courses", h.GetCourses)

	secured := api.Use(middleware.AuthorizeJWT(c.UserUsecase))
	{
		secured.GET("/profile", h.Profile)
		secured.GET("/courses/:slug", h.GetCourse)

		secured.POST("/favorites/add", h.Favorite)
		secured.POST("/favorites/remove", h.Unfavorite)
	}

	adminOnly := secured.Use(middleware.AdminOnly(c.UserUsecase))
	{
		adminOnly.POST("/courses/create", h.CreateCourse)
		adminOnly.POST("/courses/update/:id", h.UpdateCourse)
		adminOnly.POST("/courses/delete", h.DeleteCourse)
	}

	return router
}
