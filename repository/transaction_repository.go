package repository

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindByInvoiceId(invoiceId int) ([]*entity.Transaction, error)
}

type transactionRepositoryImpl struct {
	db *gorm.DB
}

type TransactionRConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(c *TransactionRConfig) TransactionRepository {
	return &transactionRepositoryImpl{
		db: c.DB,
	}
}

func (r *transactionRepositoryImpl) FindByInvoiceId(invoiceId int) (transactions []*entity.Transaction, err error) {
	err = r.db.Where("invoice_id = ?", invoiceId).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
