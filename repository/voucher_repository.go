package repository

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"gorm.io/gorm"
)

type VoucherRepository interface {
	FindById(id int) (*entity.Voucher, error)
	FindByCode(voucherCode string) (*entity.Voucher, error)
	FindAll() ([]*entity.Voucher, error)
}

type voucherRepositoryImpl struct {
	db *gorm.DB
}

type VoucherRConfig struct {
	DB *gorm.DB
}

func NewVoucherRepository(c *VoucherRConfig) VoucherRepository {
	return &voucherRepositoryImpl{
		db: c.DB,
	}
}

func (r *voucherRepositoryImpl) FindById(id int) (voucher *entity.Voucher, err error) {
	err = r.db.Where("id = ?", id).Find(voucher).Error
	if err != nil {
		return nil, err
	}
	return voucher, nil
}

func (r *voucherRepositoryImpl) FindByCode(voucherCode string) (voucher *entity.Voucher, err error) {
	err = r.db.Where("voucher_code = ?", voucherCode).Find(voucher).Error
	if err != nil {
		return nil, err
	}
	return voucher, nil
}

func (r *voucherRepositoryImpl) FindAll() (vouchers []*entity.Voucher, err error) {
	err = r.db.Find(&vouchers).Error
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}
