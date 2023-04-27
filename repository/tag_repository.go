package repository

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"gorm.io/gorm"
)

type TagRepository interface {
	FindById(id int) (*entity.Tag, error)
	FindAll() ([]*entity.Tag, error)
}

type tagRepositoryImpl struct {
	db *gorm.DB
}

type TagRConfig struct {
	DB *gorm.DB
}

func NewTagRepository(c *TagRConfig) TagRepository {
	return &tagRepositoryImpl{
		db: c.DB,
	}
}

func (r *tagRepositoryImpl) FindById(id int) (tag *entity.Tag, err error) {
	err = r.db.Where("id = ?", id).Find(tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (r *tagRepositoryImpl) FindAll() (tags []*entity.Tag, err error) {
	err = r.db.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}