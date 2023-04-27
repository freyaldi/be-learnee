package dto

type CreateCourseRequest struct {
	Title              string `json:"title" validate:"required"`
	Slug               string `json:"slug" validate:"required"`
	SummaryDescription string `json:"summary_description" validate:"required"`
	Content            string `json:"content" validate:"required"`
	ImgThumbnail       string `json:"img_thumbnail" validate:"required,url"`
	ImgUrl             string `json:"img_url" validate:"required,url"`
	AuthorName         string `json:"author_name" validate:"required"`
	CategoryId         int    `json:"category_id" validate:"required"`
	TagId              int    `json:"tag_id" validate:"required"`
}

type CourseDetailResponse struct {
	Title              string `json:"title"`
	Slug               string `json:"slug"`
	SummaryDescription string `json:"summary_description"`
	Content            string `json:"content"`
	ImgThumbnail       string `json:"img_thumbnail"`
	ImgUrl             string `json:"img_url"`
	AuthorName         string `json:"author_name"`
}
