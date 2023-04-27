package repository

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindById(id int) (*entity.Category, error)
	FindAll() ([]*entity.Category, error)
}

type categoryRepositoryImpl struct {
	db *gorm.DB
}

type CategoryRConfig struct {
	DB *gorm.DB
}

func NewCategoryRepository(c *CategoryRConfig) CategoryRepository {
	return &categoryRepositoryImpl{
		db: c.DB,
	}
}

func (r *categoryRepositoryImpl) FindById(id int) (category *entity.Category, err error) {
	err = r.db.Where("id = ?", id).Find(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepositoryImpl) FindAll() (categories []*entity.Category, err error) {
	err = r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}