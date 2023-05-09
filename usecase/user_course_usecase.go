package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	er "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/error"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
)

type UserCourseUsecase interface {
	GetUserCourses(userId int) ([]*dto.UserCourseResponse, error)
	CompleteCourse(userId int, courseId int) error
}

type userCourseUsecaseImpl struct {
	userCourseRepository repository.UserCourseRepository
}
type UserCourseUConfig struct {
	UserCourseRepository repository.UserCourseRepository
}

func NewUserCourseUsecase(c *UserCourseUConfig) UserCourseUsecase {
	return &userCourseUsecaseImpl{
		userCourseRepository: c.UserCourseRepository,
	}
}

func (u *userCourseUsecaseImpl) GetUserCourses(userId int) (responses []*dto.UserCourseResponse, err error) {
	courses, err := u.userCourseRepository.FindAll(userId)
	if err != nil {
		return nil, err
	}

	for _, c := range courses {
		response := &dto.UserCourseResponse{
			Title:              c.Course.Title,
			Slug:               c.Course.Slug,
			SummaryDescription: c.Course.SummaryDescription,
			Content:            c.Course.Content,
			ImgThumbnail:       c.Course.ImgThumbnail,
			ImgUrl:             c.Course.ImgUrl,
			AuthorName:         c.Course.AuthorName,
			Status:             string(c.Status),
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (u *userCourseUsecaseImpl) CompleteCourse(userId int, courseId int) error {
	course, err := u.userCourseRepository.Find(userId, courseId)
	if err != nil {
		return er.ErrCourseNotFound
	}

	course.Status = entity.Completed
	err = u.userCourseRepository.Update(course)
	if err != nil {
		return err
	}

	return nil
}
