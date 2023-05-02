package dto

type AddFavoriteRequest struct {
	CourseId int `json:"course_id" validate:"required"`
}

type RemoveFavoriteRequest struct {
	CourseId int `json:"course_id" validate:"required"`
}