package usecase

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	er "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/error"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/repository"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
)

type UserUsecase interface {
	Register(*dto.UserRegisterRequest) error
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
}

type UserUConfig struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(c *UserUConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepository: c.UserRepository,
	}
}

func (u *userUsecaseImpl) Register(request *dto.UserRegisterRequest) error {

	user, err := u.userRepository.FindByEmail(request.Email)
	if err != nil {
		return err
	}
	if user.Id != 0 {
		return er.ErrUserAlreadyExists
	}

	referral := util.GenerateReferral()

	newUser := &entity.User{
		Email: request.Email,
		Password: request.Password,
		Fullname: request.Fullname,
		Address: request.Address,
		PhoneNumber: request.PhoneNumber,
		RefReferral: &request.RefReferral,
		Referral: referral,
		IsAdmin: false,
		Level: entity.Newbie,
	}

	err = u.userRepository.Create(newUser)
	if err != nil {
		return err
	}

	return nil
}
