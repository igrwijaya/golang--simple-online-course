package read_product_category

import (
	"errors"
	"golang-online-course/pkg/entity/product_category"
	"golang-online-course/pkg/response"
)

type UseCase interface {
	Execute(id uint) ReadProductCategoryResponse
}

type readProductCategoryUseCase struct {
	productCategoryRepository product_category.Repository
}

func (useCase readProductCategoryUseCase) Execute(id uint) ReadProductCategoryResponse {
	productCategory := useCase.productCategoryRepository.Read(id)

	if productCategory == nil {
		return ReadProductCategoryResponse{
			Basic: response.Basic{
				Code:  400,
				Error: errors.New("product category not found"),
			},
		}
	}

	return ReadProductCategoryResponse{
		Basic: response.Success(),
		Name:  productCategory.Name,
		Image: productCategory.Image,
	}
}

func NewUseCase(productCategoryRepository product_category.Repository) UseCase {
	return &readProductCategoryUseCase{
		productCategoryRepository,
	}
}
