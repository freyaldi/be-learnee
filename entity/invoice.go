package entity

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	Id              int `gorm:"primaryKey"`
	Status          transactionStatus
	Total           float64
	BenefitDiscount float32
	UserId          int
	VoucherId       *int
	User            User    `gorm:"foreignKey:UserId;references:Id"`
	Voucher         Voucher `gorm:"foreignKey:VoucherId;references:Id"`
}

type transactionStatus string

const (
	Success transactionStatus = "success"
	Pending transactionStatus = "pending"
	Failed transactionStatus = "failed"
)
