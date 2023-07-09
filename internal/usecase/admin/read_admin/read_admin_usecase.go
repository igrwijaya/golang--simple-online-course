package read_admin

import (
	"errors"
	"golang-online-course/pkg/entity/admin"
	"golang-online-course/pkg/response"
)

type UseCase interface {
	Execute(id uint) ReadAdminResponse
}

type readAdminUseCase struct {
	adminRepository admin.Repository
}

func (useCase readAdminUseCase) Execute(id uint) ReadAdminResponse {
	adminEntity := useCase.adminRepository.Read(id)

	if adminEntity == nil {
		return ReadAdminResponse{
			Basic: response.Basic{
				Code:  400,
				Error: errors.New("admin not found"),
			},
		}
	}

	return ReadAdminResponse{
		Basic: response.Success(),
		Email: adminEntity.Email,
		Name:  adminEntity.Name,
	}
}

func NewReadAdminUseCase(adminRepository admin.Repository) UseCase {
	return &readAdminUseCase{
		adminRepository: adminRepository,
	}
}
