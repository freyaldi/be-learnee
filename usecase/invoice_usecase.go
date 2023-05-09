package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
)

type InvoiceUsecase interface {
	CreateInvoice(userId int, checkout *dto.CheckoutRequest) (*dto.InvoiceResponse, error)
}

type invoiceUsecaseImpl struct {
	invoiceRepository repository.InvoiceRepository
	cartRepository    repository.CartRepository
	voucherRepository repository.VoucherRepository
}

type InvoiceUConfig struct {
	InvoiceRepository repository.InvoiceRepository
	CartRepository    repository.CartRepository
	VoucherRepository repository.VoucherRepository
}

func NewInvoiceUsecase(c *InvoiceUConfig) InvoiceUsecase {
	return &invoiceUsecaseImpl{
		invoiceRepository: c.InvoiceRepository,
		cartRepository:    c.CartRepository,
		voucherRepository: c.VoucherRepository,
	}
}

func (u *invoiceUsecaseImpl) CreateInvoice(userId int, checkout *dto.CheckoutRequest) (*dto.InvoiceResponse, error) {
	carts, err := u.cartRepository.FindSelectedCart(userId, checkout.CartId)
	if err != nil {
		return nil, err
	}

	voucher, _ := u.voucherRepository.FindByCode(*checkout.VoucherCode)
	invoice, err := u.invoiceRepository.Insert(userId, carts, voucher)
	if err != nil {
		return nil, err
	}

	invoiceResponse := &dto.InvoiceResponse{
		Id:              invoice.Id,
		Name:            invoice.User.Fullname,
		VoucherCode:     voucher.VoucherCode,
		BenefitDiscount: util.UserLevelBenefit(&carts[0].User),
		VoucherDiscount: voucher.Benefit,
		Status:          string(invoice.Status),
		Price:           invoice.TotalPrice,
		Discount:        invoice.TotalDiscount,
		Cost:            invoice.TotalCost,
	}

	return invoiceResponse, nil
}
