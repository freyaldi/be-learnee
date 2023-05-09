package dto

type UserRegisterRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	Fullname    string `json:"fullname" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,numeric,min=10,max=13"`
	RefReferral string `json:"ref_referrral" validate:"omitempty"`
	IsAdmin     bool   `json:"is_admin" validate:"omitempty"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserDetailResponse struct {
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	IsAdmin     bool   `json:"is_admin"`
	Level       string `json:"level"`
	Referral    string `json:"referral"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
