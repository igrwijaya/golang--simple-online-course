package auth

import (
	"errors"
	"golang-online-course/internal/usecase/auth/dto"
	"golang-online-course/pkg/entity/user"
	"golang-online-course/pkg/response"
	"golang-online-course/pkg/utils"
	"gorm.io/gorm"
)

type UseCase interface {
	Login()
	Register(dto dto.RegisterRequestDto) response.Basic
}

type authUseCase struct {
	userRepository user.Repository
}

func (useCase *authUseCase) Login() {
	//TODO implement me
	panic("implement me")
}

func (useCase *authUseCase) Register(requestDto dto.RegisterRequestDto) response.Basic {
	existingUser, getUserError := useCase.userRepository.FindByEmail(requestDto.Email)

	if getUserError != nil && !errors.Is(getUserError.Error, gorm.ErrRecordNotFound) {
		return response.Basic{
			Code:  getUserError.Code,
			Error: getUserError.Error,
		}
	}

	if existingUser != nil {
		return response.Basic{
			Code:  403,
			Error: errors.New("email already registered"),
		}
	}

	newUser := user.User{
		Name:            requestDto.Name,
		Email:           requestDto.Email,
		Password:        utils.EncryptPassword(requestDto.Password),
		CodeVerified:    utils.RandString(32),
		EmailVerifiedAt: nil,
	}

	_, errCreateUser := useCase.userRepository.Create(newUser)

	if errCreateUser != nil {
		return response.Basic{
			Code:  errCreateUser.Code,
			Error: errCreateUser.Error,
		}
	}

	return response.Success()
}

func NewUseCase(userRepository user.Repository) UseCase {
	return &authUseCase{
		userRepository: userRepository,
	}
}
