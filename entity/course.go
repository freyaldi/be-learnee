package entity

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Id                 int `gorm:"primaryKey"`
	Title              string
	Slug               string
	SummaryDescription string
	Content            string
	ImgThumbnail       string
	ImgUrl             string
	AuthorName         string
	CategoryId         int
	TagId              int
	Category           Category `gorm:"foreignKey:CategoryId;references:Id"`
	Tag                Tag      `gorm:"foreignKey:TagId;references:Id"`
}
