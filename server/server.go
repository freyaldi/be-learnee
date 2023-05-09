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
	voucherRepository := repository.NewVoucherRepository(&repository.VoucherRConfig{
		DB: db.Get(),
	})
	invoiceRepository := repository.NewInvoiceRepository(&repository.InvoiceRConfig{
		DB: db.Get(),
	})
	userCourseRepository := repository.NewUserCourseRepository(&repository.UserCourseRConfig{
		DB: db.Get(),
	})
	transactionRepository := repository.NewTransactionRepository(&repository.TransactionRConfig{
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
		UserCourseRepository: userCourseRepository,
	})
	voucherUsecase := usecase.NewVoucherUsecase(&usecase.VoucherUConfig{
		VoucherRepository: voucherRepository,
	})
	invoiceUsecase := usecase.NewInvoiceUsecase(&usecase.InvoiceUConfig{
		InvoiceRepository:     invoiceRepository,
		CartRepository:        cartRepository,
		VoucherRepository:     voucherRepository,
		TransactionRepository: transactionRepository,
		UserCourseRepository:  userCourseRepository,
	})
	userCourseUsecase := usecase.NewUserCourseUsecase(&usecase.UserCourseUConfig{
		UserCourseRepository: userCourseRepository,
	})

	return NewRouter(&RouterConfig{
		UserUsecase:       userUsecase,
		CourseUsecase:     courseUsecase,
		CategoryUsecase:   categoryUsecase,
		TagUsecase:        tagUsecase,
		FavoriteUsecase:   favoriteUsecase,
		CartUsecase:       cartUsecase,
		VoucherUsecase:    voucherUsecase,
		InvoiceUsecase:    invoiceUsecase,
		UserCourseUsecase: userCourseUsecase,
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
