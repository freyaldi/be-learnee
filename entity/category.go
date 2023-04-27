package entity

type Category struct {
	Id   int `gorm:"primaryKey"`
	Name string
}
