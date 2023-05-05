package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	er "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/error"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
)

type CartUsecase interface {
	AddToCart(userId int, courseId int) error
	RemoveFromCart(userId int, courseId int) error
}

type cartUsecaseImpl struct {
	cartRepository   repository.CartRepository
	courseRepository repository.CourseRepository
}
type CartUConfig struct {
	CartRepository   repository.CartRepository
	CourseRepository repository.CourseRepository
}

func NewCartUsecase(c *CartUConfig) CartUsecase {
	return &cartUsecaseImpl{
		cartRepository:   c.CartRepository,
		courseRepository: c.CourseRepository,
	}
}

const CoursePrice = 150_000

func(u *cartUsecaseImpl) AddToCart(userId int, courseId int) error {

	_, err := u.courseRepository.FindById(courseId)
	if err != nil {
		return err
	}

	cartedCourse, err := u.cartRepository.Find(userId, courseId)
	if err != nil {
		return err
	}
	
	if cartedCourse.Id != 0 {
		return er.ErrCourseAlreadyCarted
	}

	cartedCourse = &entity.Cart{
		UserId: userId,
		CourseId: courseId,
		Price: CoursePrice,
	}

	err = u.cartRepository.Insert(cartedCourse)
	if err != nil {
		return err
	}

	return nil
}

func(u *cartUsecaseImpl) RemoveFromCart(userId int, courseId int) error {

	cartedCourse, err := u.cartRepository.Find(userId, courseId)
	if err != nil {
		return err
	}

	err = u.cartRepository.Delete(cartedCourse.Id)
	if err != nil {
		return err
	}

	return nil
}