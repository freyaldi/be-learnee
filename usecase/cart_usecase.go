package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	er "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/error"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
	"gorm.io/gorm"
)

type CartUsecase interface {
	AddToCart(userId int, courseId int) error
	RemoveFromCart(userId int, courseId int) error
	GetCarts(userId int) ([]*dto.CartsResponse, error)
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

	if cartedCourse.Id == 0 {
		return gorm.ErrRecordNotFound
	}

	err = u.cartRepository.Delete(cartedCourse.Id)
	if err != nil {
		return err
	}

	return nil
}

func(u *cartUsecaseImpl) GetCarts(userId int) (carts []*dto.CartsResponse, err error) {
	cartedCourses, err := u.cartRepository.FindAll(userId)
	if err != nil {
		return nil, err
	}

	for _, cc := range cartedCourses {
		cart := &dto.CartsResponse{
			CourseId: cc.CourseId,
			Title: cc.Course.Title,
			ImgThumbnail: cc.Course.ImgThumbnail,
			AuthorName: cc.Course.AuthorName,
			Price: cc.Course.Price,
		}
		carts = append(carts, cart)
	}
	return carts, nil
}