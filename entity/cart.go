package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Id       int `gorm:"primaryKey"`
	CourseId int
	UserId   int
	Price    float64
	Course   Course `gorm:"foreignKey:CourseId;references:Id"`
	User     User   `gorm:"foreignKey:UserId;references:Id"`
}
