package entity

type Tag struct {
	Id   int `gorm:"primaryKey"`
	Name string
}
