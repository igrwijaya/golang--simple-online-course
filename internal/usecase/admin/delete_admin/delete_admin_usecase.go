package delete_admin

import (
	"golang-online-course/pkg/entity/admin"
	"golang-online-course/pkg/response"
)

type UseCase interface {
	Execute(id uint) response.Basic
}

type deleteAdminUseCase struct {
	adminRepository admin.Repository
}

func (useCase deleteAdminUseCase) Execute(id uint) response.Basic {
	useCase.adminRepository.Delete(id)

	return response.Success()
}

func NewDeleteAdminUseCase(adminRepository admin.Repository) UseCase {
	return &deleteAdminUseCase{
		adminRepository: adminRepository,
	}
}
