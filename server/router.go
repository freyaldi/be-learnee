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
	CartUsecase     usecase.CartUsecase
	VoucherUsecase  usecase.VoucherUsecase
	InvoiceUsecase  usecase.InvoiceUsecase
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handler.New(&handler.Config{
		UserUsecase:     c.UserUsecase,
		CourseUsecase:   c.CourseUsecase,
		CategoryUsecase: c.CategoryUsecase,
		TagUsecase:      c.TagUsecase,
		FavoriteUsecase: c.FavoriteUsecase,
		CartUsecase:     c.CartUsecase,
		VoucherUsecase:  c.VoucherUsecase,
		InvoiceUsecase:  c.InvoiceUsecase,
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

		secured.GET("/carts", h.Carts)
		secured.POST("/carts/add", h.AddToCart)
		secured.POST("/carts/remove", h.RemoveFromCart)

		secured.GET("/vouchers", h.Vouchers)

		secured.POST("/transactions/checkout", h.Checkout)
	}

	adminOnly := secured.Use(middleware.AdminOnly(c.UserUsecase))
	{
		adminOnly.POST("/courses/create", h.CreateCourse)
		adminOnly.POST("/courses/update/:id", h.UpdateCourse)
		adminOnly.POST("/courses/delete", h.DeleteCourse)
	}

	return router
}
