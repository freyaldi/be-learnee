package repository

import (
	"log"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	Insert(userId int, carts []*entity.Cart, voucher *entity.Voucher) (*entity.Invoice, error)
	Update(*entity.Invoice) error
	FindAll() ([]*entity.Invoice, error)
	FindById(id int) (*entity.Invoice, error)
}

type invoiceRepositoryImpl struct {
	db *gorm.DB
}

type InvoiceRConfig struct {
	DB *gorm.DB
}

func NewInvoiceRepository(c *InvoiceRConfig) InvoiceRepository {
	return &invoiceRepositoryImpl{
		db: c.DB,
	}
}

func (r *invoiceRepositoryImpl) Insert(userId int, carts []*entity.Cart, voucher *entity.Voucher) (*entity.Invoice, error) {
	tx := r.db.Begin()

	invoice := &entity.Invoice{Status: entity.Pending, UserId: userId}
	err := tx.Create(&invoice).Error
	if err != nil {
		tx.Rollback()
		log.Print(1)
		return nil, err
	}
	var price float64 = 0
	for _, c := range carts {
		transaction := &entity.Transaction{
			InvoiceId: invoice.Id,
			CourseId:  c.CourseId,
			SoldPrice: c.Price,
		}

		err = tx.Create(transaction).Error
		if err != nil {
			tx.Rollback()
			log.Print(2)
			return nil, err
		}

		err = tx.Delete(c).Error
		if err != nil {
			tx.Rollback()
			log.Print(4)
			return nil, err
		}

		price += transaction.SoldPrice
	}

	price = price * (1 - float64(invoice.BenefitDiscount))
	if voucher != nil {
		price = price * (1 - float64(voucher.Benefit))
	}
	invoice.Total = price

	err = tx.Save(invoice).Error
	if err != nil {
		tx.Rollback()
		log.Print(3)
		return nil, err
	}

	tx.Commit()
	return invoice, nil
}

func (r *invoiceRepositoryImpl) Update(*entity.Invoice) error {
	panic("unimplemented")
}

func (r *invoiceRepositoryImpl) FindAll() ([]*entity.Invoice, error) {
	panic("unimplemented")
}

func (r *invoiceRepositoryImpl) FindById(id int) (*entity.Invoice, error) {
	panic("unimplemented")
}
