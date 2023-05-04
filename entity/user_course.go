package entity

import "gorm.io/gorm"

type UserCourse struct {
	gorm.Model
	Id       int `gorm:"primaryKey"`
	CourseId int
	UserId   int
	Status   courseStatus
	Course   Course `gorm:"foreignKey:CourseId;references:Id"`
	User     User   `gorm:"foreignKey:UserId;references:Id"`
}

type courseStatus string

const (
	Completed  courseStatus = "completed"
	OnProgress courseStatus = "on progress"
)
