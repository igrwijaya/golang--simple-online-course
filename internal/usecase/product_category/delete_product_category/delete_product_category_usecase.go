package delete_product_category

import (
	"golang-online-course/pkg/entity/product_category"
	"golang-online-course/pkg/response"
)

type UseCase interface {
	Execute(id uint) response.Basic
}

type deleteProductCategoryUseCase struct {
	productCategoryRepository product_category.Repository
}

func (useCase deleteProductCategoryUseCase) Execute(id uint) response.Basic {
	useCase.productCategoryRepository.Delete(id)

	return response.Success()
}

func NewUseCase(productCategoryRepository product_category.Repository) UseCase {
	return &deleteProductCategoryUseCase{productCategoryRepository}
}
