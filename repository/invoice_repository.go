package repository

import (
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
		return nil, err
	}
	
	var price float64 = 0
	for _, c := range carts {
		transaction := &entity.Transaction{
			InvoiceId: invoice.Id,
			CourseId:  c.CourseId,
			SoldPrice: c.Course.Price,
		}

		err = tx.Create(transaction).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.Delete(c).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		price += transaction.SoldPrice
	}

	discount := price * float64(invoice.BenefitDiscount)
	if voucher != nil {
		discount = discount + price * float64(voucher.Benefit)
	}
	cost := price - discount

	invoice.TotalPrice = price
	invoice.TotalDiscount = discount
	invoice.TotalCost = cost

	err = tx.Save(invoice).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return invoice, nil
}

func (r *invoiceRepositoryImpl) Update(invoice *entity.Invoice) error {
	tx := r.db.Begin()

	invoice.Status = entity.Success
	err := tx.Save(invoice).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *invoiceRepositoryImpl) FindAll() ([]*entity.Invoice, error) {
	panic("unimplemented")
}

func (r *invoiceRepositoryImpl) FindById(id int) (invoice *entity.Invoice, err error) {
	err = r.db.First(&invoice, id).Error
	if err != nil {
		return nil, err
	}
	return invoice, nil
}
