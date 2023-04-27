package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
)

type TagUsecase interface {
	GetTags() ([]*entity.Tag, error)
}

type tagUsecaseImpl struct {
	tagRepository repository.TagRepository
}
type TagUConfig struct {
	TagRepository repository.TagRepository
}

func NewTagUsecase(c *TagUConfig) TagUsecase {
	return &tagUsecaseImpl{
		tagRepository: c.TagRepository,
	}
}

func(u *tagUsecaseImpl) GetTags() (tags []*entity.Tag, err error) {
	tags, err = u.tagRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tags, nil
}