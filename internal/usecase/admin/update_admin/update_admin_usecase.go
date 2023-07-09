package update_admin

import (
	"errors"
	"golang-online-course/pkg/entity/admin"
	"golang-online-course/pkg/response"
)

type UseCase interface {
	Execute(id uint, request UpdateAdminRequest) response.Basic
}

type updateAdminUseCase struct {
	adminRepository admin.Repository
}

func (useCase updateAdminUseCase) Execute(id uint, request UpdateAdminRequest) response.Basic {
	existingAdmin := useCase.adminRepository.FindByEmail(request.Email)
	if existingAdmin != nil {
		return response.Basic{
			Code:  400,
			Error: errors.New("email is already exist"),
		}
	}

	useCase.adminRepository.Update(id, request.Name, request.Email)

	return response.Success()
}

func NewUpdateAdminUseCase(adminRepository admin.Repository) UseCase {
	return &updateAdminUseCase{adminRepository: adminRepository}
}
