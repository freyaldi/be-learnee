package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	er "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/error"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
)

type InvoiceUsecase interface {
	CreateInvoice(userId int, checkout *dto.CheckoutRequest) (*dto.InvoiceResponse, error)
	UpdateInvoice(request *dto.UpdateTransactionRequest) error
}

type invoiceUsecaseImpl struct {
	invoiceRepository     repository.InvoiceRepository
	cartRepository        repository.CartRepository
	voucherRepository     repository.VoucherRepository
	transactionRepository repository.TransactionRepository
	userCourseRepository  repository.UserCourseRepository
}

type InvoiceUConfig struct {
	InvoiceRepository     repository.InvoiceRepository
	CartRepository        repository.CartRepository
	VoucherRepository     repository.VoucherRepository
	TransactionRepository repository.TransactionRepository
	UserCourseRepository  repository.UserCourseRepository
}

func NewInvoiceUsecase(c *InvoiceUConfig) InvoiceUsecase {
	return &invoiceUsecaseImpl{
		invoiceRepository:     c.InvoiceRepository,
		cartRepository:        c.CartRepository,
		voucherRepository:     c.VoucherRepository,
		transactionRepository: c.TransactionRepository,
		userCourseRepository:  c.UserCourseRepository,
	}
}

func (u *invoiceUsecaseImpl) CreateInvoice(userId int, checkout *dto.CheckoutRequest) (*dto.InvoiceResponse, error) {
	carts, err := u.cartRepository.FindSelectedCart(userId, checkout.CartId)
	if err != nil {
		return nil, err
	}

	if len(carts) == 0 {
		return nil, er.ErrCartIsEmpty
	}

	var voucher = &entity.Voucher{}
	if checkout.VoucherCode != nil {
		voucher, _ = u.voucherRepository.FindByCode(*checkout.VoucherCode)
	}
	invoice, err := u.invoiceRepository.Insert(userId, carts, voucher)
	if err != nil {
		return nil, err
	}

	invoiceResponse := &dto.InvoiceResponse{
		Id:              invoice.Id,
		Name:            invoice.User.Fullname,
		BenefitDiscount: util.UserLevelBenefit(&carts[0].User),
		Status:          string(invoice.Status),
		Price:           invoice.TotalPrice,
		Discount:        invoice.TotalDiscount,
		Cost:            invoice.TotalCost,
	}

	if voucher.Id != 0 {
		invoiceResponse.VoucherCode = voucher.VoucherCode
		invoiceResponse.VoucherDiscount = voucher.Benefit
	}

	return invoiceResponse, nil
}

func (u *invoiceUsecaseImpl) UpdateInvoice(request *dto.UpdateTransactionRequest) error {
	invoice, err := u.invoiceRepository.FindById(request.InvoiceId)
	if err != nil {
		return err
	}

	if request.Status == string(invoice.Status) {
		return er.ErrTransactionStatusAlreadyAsExpected
	}

	if request.Status == string(entity.Success) {
		transactions, err := u.transactionRepository.FindByInvoiceId(invoice.Id)
		if err != nil {
			return err
		}

		for _, t := range transactions {
			userCourse := &entity.UserCourse{
				CourseId: t.CourseId,
				UserId:   invoice.UserId,
				Status:   entity.OnProgress,
			}

			err = u.userCourseRepository.Insert(userCourse)
			if err != nil {
				return err
			}
		}
	}

	err = u.invoiceRepository.Update(invoice)
	if err != nil {
		return err
	}

	return nil
}
