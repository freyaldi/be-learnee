package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
)

type CourseUsecase interface {
	GetCourseBySlug(slug string) (*dto.CourseDetailResponse, error)
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
