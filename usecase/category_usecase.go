package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
)

type CategoryUsecase interface {
	GetCategories() ([]*entity.Category, error)
}

type categoryUsecaseImpl struct {
	categoryRepository repository.CategoryRepository
}
type CategoryUConfig struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(c *CategoryUConfig) CategoryUsecase {
	return &categoryUsecaseImpl{
		categoryRepository: c.CategoryRepository,
	}
}

func(u *categoryUsecaseImpl) GetCategories() (categories []*entity.Category, err error) {
	categories, err = u.categoryRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}