package repository

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"gorm.io/gorm"
)

type CartRepository interface {
	Insert(*entity.Cart) error
	Delete(Id int) error
	FindSelectedCart(userId int, courseIds []int) ([]*entity.Cart, error)
	FindAll(userId int) ([]*entity.Cart, error)
	Find(userId int, courseId int) (*entity.Cart, error)
}

type cartRepositoryImpl struct {
	db *gorm.DB
}

type CartRConfig struct {
	DB *gorm.DB
}

func NewCartRepository(c *CartRConfig) CartRepository {
	return &cartRepositoryImpl{
		db: c.DB,
	}
}

func (r *cartRepositoryImpl) Insert(cart *entity.Cart) error {
	err := r.db.Create(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *cartRepositoryImpl) Delete(Id int) error {
	err := r.db.Delete(&entity.Cart{}, Id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *cartRepositoryImpl) FindSelectedCart(userId int, courseIds []int) (carts []*entity.Cart, err error) {
	err = r.db.Where("user_id = ?", userId).Find(&carts, courseIds).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *cartRepositoryImpl) FindAll(userId int) (carts []*entity.Cart, err error) {
	err = r.db.Where("user_id = ?", userId).Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *cartRepositoryImpl) Count(CourseId int) (total int64, err error) {
	err = r.db.Model(&entity.Cart{}).Where("course_id = ?", CourseId).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *cartRepositoryImpl) Find(userId int, courseId int) (cart *entity.Cart, err error) {
	err = r.db.Where("course_id = ?", courseId).Where("user_id = ?", userId).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}
