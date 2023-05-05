package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
)

type VoucherUsecase interface {
	GetVouchers() ([]*entity.Voucher, error)
}

type voucherUsecaseImpl struct {
	voucherRepository repository.VoucherRepository
}
type VoucherUConfig struct {
	VoucherRepository repository.VoucherRepository
}

func NewVoucherUsecase(c *VoucherUConfig) VoucherUsecase {
	return &voucherUsecaseImpl{
		voucherRepository: c.VoucherRepository,
	}
}

func(u *voucherUsecaseImpl) GetVouchers() (vouchers []*entity.Voucher, err error) {
	vouchers, err = u.voucherRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}