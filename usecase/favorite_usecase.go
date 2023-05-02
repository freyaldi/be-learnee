package usecase

import (
	"log"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	er "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/error"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
)

type FavoriteUsecase interface {
	AddFavorite(userId int, courseId int) error
	RemoveFavorite(userId int, courseId int) error
}

type favoriteUsecaseImpl struct {
	favoriteRepository repository.FavoriteRepository
	courseRepository   repository.CourseRepository
}
type FavoriteUConfig struct {
	FavoriteRepository repository.FavoriteRepository
	CourseRepository   repository.CourseRepository
}

func NewFavoriteUsecase(c *FavoriteUConfig) FavoriteUsecase {
	return &favoriteUsecaseImpl{
		favoriteRepository: c.FavoriteRepository,
		courseRepository:   c.CourseRepository,
	}
}

func (u *favoriteUsecaseImpl) AddFavorite(userId int, courseId int) error {
	_, err := u.courseRepository.FindById(courseId)
	if err != nil {
		return err
	}

	favorite, err := u.favoriteRepository.Find(userId, courseId)
	if err != nil {
		return err
	}
	if favorite.Id != 0 {
		return er.ErrCourseAlreadyFavorited
	}

	favorite = &entity.Favorite{
		UserId:   userId,
		CourseId: courseId,
	}

	err = u.favoriteRepository.Insert(favorite)
	if err != nil {
		return err
	}

	return nil
}

func (u *favoriteUsecaseImpl) RemoveFavorite(userId int, courseId int) error {

	favorite, err := u.favoriteRepository.Find(userId, courseId)
	if err != nil {
		return err
	}

	deletedAt, _ := favorite.DeletedAt.Value()
	log.Print(deletedAt)
	if deletedAt != nil {
		return er.ErrCourseAlreadyUnFavorited
	}

	err = u.favoriteRepository.Delete(favorite.Id)
	if err != nil {
		return err
	}

	return nil
}
