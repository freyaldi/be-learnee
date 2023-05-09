package entity

type User struct {
	Id          int    `gorm:"primaryKey"`
	Email       string `gorm:"unique"`
	Password    string
	IsAdmin     bool
	Fullname    string
	Address     string
	PhoneNumber string
	Level       *level
	Referral    string
	RefReferral *string
}

type level string

var (
	Newbie level = "newbie"
	Junior level = "junior"
	Senior level = "senior"
	Master level = "master"
)
