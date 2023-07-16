package update_product_category

import (
	"golang-online-course/pkg/entity/product_category"
	"golang-online-course/pkg/response"
)

type UseCase interface {
	Execute(id uint, bodyRequest UpdateProductCategoryBodyRequest) response.Basic
}

type updateProductCategoryUseCase struct {
	productCategoryRepository product_category.Repository
}

func (useCase updateProductCategoryUseCase) Execute(id uint, bodyRequest UpdateProductCategoryBodyRequest) response.Basic {
	useCase.productCategoryRepository.Update(id, bodyRequest.Name, bodyRequest.Image)

	return response.Success()
}

func NewUseCase(productCategoryRepository product_category.Repository) UseCase {
	return &updateProductCategoryUseCase{productCategoryRepository}
}
