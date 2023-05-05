package entity

type Voucher struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	VoucherCode string
	Benefit     float32
}
