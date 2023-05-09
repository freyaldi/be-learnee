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
	Login(*dto.UserLoginRequest) (string, error)
	Profile(id int) (*dto.UserDetailResponse, error)
	UpdateProfile(id int, request *dto.UserUpdateProfileRequest) error
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
	jwt            util.Auth
}

type UserUConfig struct {
	UserRepository repository.UserRepository
	JWT            util.Auth
}

func NewUserUsecase(c *UserUConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepository: c.UserRepository,
		jwt:            c.JWT,
	}
}

func (u *userUsecaseImpl) Register(request *dto.UserRegisterRequest) error {
	referral := util.GenerateReferral()

	hashedPassword, err := u.jwt.HashPassword(request.Password)
	if err != nil {
		return err
	}

	newUser := &entity.User{
		Email:       request.Email,
		Password:    hashedPassword,
		Fullname:    request.Fullname,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		RefReferral: &request.RefReferral,
		Referral:    referral,
		IsAdmin:     request.IsAdmin,
	}

	if !request.IsAdmin {
		level := &entity.Newbie
		newUser.Level = level
	}

	err = u.userRepository.Create(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecaseImpl) Login(request *dto.UserLoginRequest) (string, error) {
	user, err := u.userRepository.FindByEmail(request.Email)
	if err != nil {
		return "", er.ErrIncorrectCredentials
	}

	isPasswordCorrect := u.jwt.ComparePassword(user.Password, request.Password)
	if !isPasswordCorrect {
		return "", er.ErrIncorrectCredentials
	}

	token, err := u.jwt.GenerateToken(user.Id, user.IsAdmin)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userUsecaseImpl) Profile(id int) (*dto.UserDetailResponse, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	response := &dto.UserDetailResponse{
		Fullname:    user.Fullname,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		IsAdmin:     user.IsAdmin,
		Referral:    user.Referral,
	}

	if user.Level != nil {
		response.Level = string(*user.Level)
	}

	return response, nil
}

func (u *userUsecaseImpl) UpdateProfile(id int, request *dto.UserUpdateProfileRequest) error {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return err
	}

	isPasswordCorrect := u.jwt.ComparePassword(user.Password, request.PasswordConfirmation)
	if !isPasswordCorrect {
		return er.ErrIncorrectCredentials
	}

	user.Email = request.Email
	user.Fullname = request.Fullname
	user.Address = request.Address
	user.PhoneNumber = request.PhoneNumber

	err = u.userRepository.Update(user)
	if err != nil {
		return err
	}

	return nil
}
