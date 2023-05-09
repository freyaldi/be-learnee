package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
)

type InvoiceUsecase interface {
	CreateInvoice(userId int, checkout *dto.CheckoutRequest) (*dto.InvoiceResponse, error)
}

type invoiceUsecaseImpl struct {
	invoiceRepository repository.InvoiceRepository
	cartRepository    repository.CartRepository
}

type InvoiceUConfig struct {
	InvoiceRepository repository.InvoiceRepository
	CartRepository    repository.CartRepository
}

func NewInvoiceUsecase(c *InvoiceUConfig) InvoiceUsecase {
	return &invoiceUsecaseImpl{
		invoiceRepository: c.InvoiceRepository,
		cartRepository:    c.CartRepository,
	}
}

func (u *invoiceUsecaseImpl) CreateInvoice(userId int, checkout *dto.CheckoutRequest) (*dto.InvoiceResponse, error) {
	carts, err := u.cartRepository.FindSelectedCart(userId, checkout.CartId)
	if err != nil {
		return nil, err
	}

	invoice, err := u.invoiceRepository.Insert(userId, carts, &entity.Voucher{})
	if err != nil {
		return nil, err
	}

	invoiceResponse := &dto.InvoiceResponse{
		Id:              invoice.Id,
		Name:            invoice.User.Fullname,
		VoucherCode:     "",
		BenefitDiscount: 0,
		VoucherDiscount: 0,
		Status:          "",
		Discount:        0,
		Total:           invoice.Total,
	}

	return invoiceResponse, nil
}
