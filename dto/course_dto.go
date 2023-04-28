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

type UpdateCourseRequest struct {
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

type DeleteCourseRequest struct {
	Id int `json:"id" validate:"required"`
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

type CourseResponse struct {
	Title              string `json:"title"`
	Slug               string `json:"slug"`
	ImgThumbnail       string `json:"img_thumbnail"`
	ImgUrl             string `json:"img_url"`
	AuthorName         string `json:"author_name"`
}

type CourseRequestQuery struct {
	CourseFilterRequest
	SortBy   string `form:"sortBy"`
	Sort     string `form:"sort"`
	Limit    int    `form:"limit"`
	Page     int    `form:"page"`
}

type CourseFilterRequest struct {
	Search   string `form:"s"`	
	Category int    `form:"category"`
	Tag      int    `form:"tag"`
}
