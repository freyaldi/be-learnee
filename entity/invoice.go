package entity

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	Id              int `gorm:"primaryKey"`
	Status          string
	Total           float64
	BenefitDiscount float32
	UserId          int
	VoucherId       int
	User            User    `gorm:"foreignKey:UserId;references:Id"`
	Voucher         Voucher `gorm:"foreignKey:VoucherId;references:Id"`
}
