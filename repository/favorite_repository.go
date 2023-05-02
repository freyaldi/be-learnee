package repository

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"gorm.io/gorm"
)

type FavoriteRepository interface {
	Insert(*entity.Favorite) error
	Delete(Id int) error
	FindAll(userId int) ([]*entity.Favorite, error)
	Count(courseId int) (int64, error)
	Find(userId int, courseId int) (*entity.Favorite, error)
}

type favoriteRepositoryImpl struct {
	db *gorm.DB
}

type FavoriteRConfig struct {
	DB *gorm.DB
}

func NewFavoriteRepository(c *FavoriteRConfig) FavoriteRepository {
	return &favoriteRepositoryImpl{
		db: c.DB,
	}
}

func (r *favoriteRepositoryImpl) Insert(favorite *entity.Favorite) error {
	err := r.db.Create(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *favoriteRepositoryImpl) Delete(Id int) error {
	err := r.db.Delete(&entity.Favorite{}, Id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *favoriteRepositoryImpl) FindAll(userId int) (favorites []*entity.Favorite, err error) {
	err = r.db.Where("user_id = ?", userId).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (r *favoriteRepositoryImpl) Count(CourseId int) (total int64, err error) {
	err = r.db.Model(&entity.Favorite{}).Where("course_id = ?", CourseId).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *favoriteRepositoryImpl) Find(userId int, courseId int) (favorite *entity.Favorite, err error) {
	r.db.Unscoped().Where("course_id = ?", courseId).Where("user_id = ?", userId).First(&favorite)
	if err != nil {
		return nil, err
	}
	return favorite, nil
}
