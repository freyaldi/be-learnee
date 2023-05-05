package dto

type AddToCartRequest struct {
	CourseId int `json:"course_id" validate:"required"`
}

type RemoveFromCartRequest struct {
	CourseId int `json:"course_id" validate:"required"`
}

type CartsResponse struct {
	CourseId     int    `json:"course_id"`
	Title        string `json:"title"`
	ImgThumbnail string `json:"img_thumbnail"`
	AuthorName   string `json:"author_name"`
	Price        float64 `json:"price"`
}
