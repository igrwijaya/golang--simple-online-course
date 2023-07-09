package get_admin

import (
	"golang-online-course/pkg/entity/admin"
	"golang-online-course/pkg/response"
)

type UseCase interface {
	Execute(page int, size int) GetAdminResponse
}

type getAdminUseCase struct {
	adminRepository admin.Repository
}

func (useCase getAdminUseCase) Execute(page int, size int) GetAdminResponse {
	totalRecord, entities := useCase.adminRepository.Get(page, size)

	var admins []AdminDto

	for _, entity := range entities {
		admins = append(admins, AdminDto{
			Id:    entity.Id,
			Email: entity.Email,
			Name:  entity.Name,
		})
	}

	return GetAdminResponse{
		Basic:       response.Success(),
		TotalRecord: totalRecord,
		CurrentPage: page,
		Limit:       size,
		Data:        admins,
	}
}

func NewGetAdminUseCase(adminRepository admin.Repository) UseCase {
	return &getAdminUseCase{
		adminRepository: adminRepository,
	}
}
