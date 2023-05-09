package dto

type UserCourseResponse struct {
	Title              string `json:"title"`
	Slug               string `json:"slug"`
	SummaryDescription string `json:"summary_description"`
	Content            string `json:"content"`
	ImgThumbnail       string `json:"img_thumbnail"`
	ImgUrl             string `json:"img_url"`
	AuthorName         string `json:"author_name"`
	Status             string `json:"status"`
}
