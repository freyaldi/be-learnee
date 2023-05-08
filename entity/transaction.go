package entity

type Transaction struct {
	Id        int `gorm:"primaryKey"`
	InvoiceId int
	CourseId  int
	SoldPrice float64
	Invoice   Invoice `gorm:"foreignKey:InvoiceId;references:Id"`
	Course    Course  `gorm:"foreignKey:CourseId;references:Id"`
}
