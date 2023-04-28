package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
)

type CourseUsecase interface {
	CreateCourse(request *dto.CreateCourseRequest) error
	UpdateCourse(id int, request *dto.UpdateCourseRequest) error
	DeleteCourse(id int) error
	GetCourseBySlug(slug string) (*dto.CourseDetailResponse, error)
	GetCourses(query *dto.CourseRequestQuery) ([]*dto.CourseResponse, error)
}

type courseUsecaseImpl struct {
	courseRepository repository.CourseRepository
}
type CourseUConfig struct {
	CourseRepository repository.CourseRepository
}

func NewCourseUsecase(c *CourseUConfig) CourseUsecase {
	return &courseUsecaseImpl{
		courseRepository: c.CourseRepository,
	}
}

func (u *courseUsecaseImpl) GetCourseBySlug(slug string) (*dto.CourseDetailResponse, error) {
	course, err := u.courseRepository.FindBySlug(slug)
	if err != nil {
		return nil, err
	}

	response := &dto.CourseDetailResponse{
		Title:              course.Title,
		Slug:               course.Slug,
		SummaryDescription: course.SummaryDescription,
		Content:            course.Content,
		ImgThumbnail:       course.ImgThumbnail,
		ImgUrl:             course.ImgUrl,
		AuthorName:         course.AuthorName,
	}

	return response, nil
}

func (u *courseUsecaseImpl) CreateCourse(request *dto.CreateCourseRequest) error {
	course := &entity.Course{
		Title:              request.Title,
		Slug:               request.Slug,
		SummaryDescription: request.SummaryDescription,
		Content:            request.Content,
		ImgThumbnail:       request.ImgThumbnail,
		ImgUrl:             request.ImgUrl,
		AuthorName:         request.AuthorName,
		CategoryId:         request.CategoryId,
		TagId:              request.TagId,
	}

	err := u.courseRepository.Create(course)
	if err != nil {
		return err
	}

	return nil
}

func (u *courseUsecaseImpl) UpdateCourse(id int, request *dto.UpdateCourseRequest) error {

	course := &entity.Course{
		Title:              request.Title,
		Slug:               request.Slug,
		SummaryDescription: request.SummaryDescription,
		Content:            request.Content,
		ImgThumbnail:       request.ImgThumbnail,
		ImgUrl:             request.ImgUrl,
		AuthorName:         request.AuthorName,
		CategoryId:         request.CategoryId,
		TagId:              request.TagId,
	}

	err := u.courseRepository.Update(id, course)
	if err != nil {
		return err
	}

	return nil
}

func (u *courseUsecaseImpl) DeleteCourse(id int) error {
	course, err := u.courseRepository.FindById(id)
	if err != nil {
		return err
	}

	err = u.courseRepository.Delete(course)
	if err != nil {
		return err
	}

	return nil
}

func (u *courseUsecaseImpl) GetCourses(query *dto.CourseRequestQuery) (responses []*dto.CourseResponse, err error) {
	courses, err := u.courseRepository.Find(query)
	if err != nil {
		return nil, err
	}

	for _, c := range courses {
		response := &dto.CourseResponse{
			Title:        c.Title,
			Slug:         c.Slug,
			ImgThumbnail: c.ImgThumbnail,
			ImgUrl:       c.ImgUrl,
			AuthorName:   c.AuthorName,
		}
		responses = append(responses, response)
	}

	return responses, nil
}
