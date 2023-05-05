package dto

type AddToCartRequest struct {
	CourseId int `json:"course_id" validate:"required"`
}