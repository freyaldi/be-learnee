package handler

import "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/usecase"

type Handler struct {
	userUsecase     usecase.UserUsecase
	courseUsecase   usecase.CourseUsecase
	categoryUsecase usecase.CategoryUsecase
	tagUsecase      usecase.TagUsecase
	favoriteUsecase usecase.FavoriteUsecase
	cartUsecase     usecase.CartUsecase
	voucherUsecase  usecase.VoucherUsecase
	invoiceUsecase  usecase.InvoiceUsecase
}

type Config struct {
	UserUsecase     usecase.UserUsecase
	CourseUsecase   usecase.CourseUsecase
	CategoryUsecase usecase.CategoryUsecase
	TagUsecase      usecase.TagUsecase
	FavoriteUsecase usecase.FavoriteUsecase
	CartUsecase     usecase.CartUsecase
	VoucherUsecase  usecase.VoucherUsecase
	InvoiceUsecase  usecase.InvoiceUsecase
}

func New(c *Config) *Handler {
	return &Handler{
		userUsecase:     c.UserUsecase,
		courseUsecase:   c.CourseUsecase,
		categoryUsecase: c.CategoryUsecase,
		tagUsecase:      c.TagUsecase,
		favoriteUsecase: c.FavoriteUsecase,
		cartUsecase:     c.CartUsecase,
		voucherUsecase:  c.VoucherUsecase,
		invoiceUsecase:  c.InvoiceUsecase,
	}
}
