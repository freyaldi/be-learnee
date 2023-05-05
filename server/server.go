package server

import (
	"log"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/db"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/usecase"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	userRepository := repository.NewUserRepository(&repository.UserRConfig{
		DB: db.Get(),
	})
	courseRepository := repository.NewCourseRepository(&repository.CourseRConfig{
		DB: db.Get(),
	})
	categoryRepository := repository.NewCategoryRepository(&repository.CategoryRConfig{
		DB: db.Get(),
	})
	tagRepository := repository.NewTagRepository(&repository.TagRConfig{
		DB: db.Get(),
	})
	favoriteRepository := repository.NewFavoriteRepository(&repository.FavoriteRConfig{
		DB: db.Get(),
	})
	cartRepository := repository.NewCartRepository(&repository.CartRConfig{
		DB: db.Get(),
	})

	jwt := util.NewAuth(&util.AuthConfig{})

	userUsecase := usecase.NewUserUsecase(&usecase.UserUConfig{
		UserRepository: userRepository,
		JWT:            jwt,
	})
	courseUsecase := usecase.NewCourseUsecase(&usecase.CourseUConfig{
		CourseRepository: courseRepository,
	})
	categoryUsecase := usecase.NewCategoryUsecase(&usecase.CategoryUConfig{
		CategoryRepository: categoryRepository,
	})
	tagUsecase := usecase.NewTagUsecase(&usecase.TagUConfig{
		TagRepository: tagRepository,
	})
	favoriteUsecase := usecase.NewFavoriteUsecase(&usecase.FavoriteUConfig{
		FavoriteRepository: favoriteRepository,
		CourseRepository:   courseRepository,
	})
	cartUsecase := usecase.NewCartUsecase(&usecase.CartUConfig{
		CartRepository:   cartRepository,
		CourseRepository: courseRepository,
	})

	return NewRouter(&RouterConfig{
		UserUsecase:     userUsecase,
		CourseUsecase:   courseUsecase,
		CategoryUsecase: categoryUsecase,
		TagUsecase:      tagUsecase,
		FavoriteUsecase: favoriteUsecase,
		CartUsecase:     cartUsecase,
	})
}

func Init() {
	r := createRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}
