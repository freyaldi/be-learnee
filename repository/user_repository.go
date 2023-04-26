package repository

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	FindById(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

type UserRConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *UserRConfig) UserRepository {
	return &userRepositoryImpl{
		db: c.DB,
	}
}

func (r *userRepositoryImpl) Create(user *entity.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImpl) Update(user *entity.User) error {
	err := r.db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImpl) FindById(id int) (user *entity.User, err error) {
	err = r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) FindByEmail(email string) (user *entity.User, err error) {
	err = r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
