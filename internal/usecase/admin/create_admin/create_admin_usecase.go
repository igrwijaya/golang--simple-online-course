package create_admin

import (
	"errors"
	"golang-online-course/pkg/entity/admin"
	"golang-online-course/pkg/response"
	"golang-online-course/pkg/utils"
)

type UseCase interface {
	Execute(request CreateAdminRequest) response.Basic
}

type createAdminUseCase struct {
	adminRepository admin.Repository
}

func (useCase *createAdminUseCase) Execute(request CreateAdminRequest) response.Basic {

	if request.Password != request.ConfirmPassword {
		return response.Basic{
			Code:  400,
			Error: errors.New("confirm password not match"),
		}
	}

	existingAdmin := useCase.adminRepository.FindByEmail(request.Email)

	if existingAdmin != nil {
		return response.Basic{
			Code:  400,
			Error: errors.New("email already exist"),
		}
	}

	adminEntity := admin.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: utils.EncryptPassword(request.Password),
	}

	useCase.adminRepository.Create(adminEntity)

	return response.Success()
}

func NewCreateAdminUseCase(adminRepository admin.Repository) UseCase {
	return &createAdminUseCase{adminRepository: adminRepository}
}
