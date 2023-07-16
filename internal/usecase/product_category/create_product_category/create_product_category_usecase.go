package create_product_category

import (
	"fmt"
	"golang-online-course/pkg/entity/product_category"
	"golang-online-course/pkg/response"
)

type UseCase interface {
	Execute(request CreateProductCategoryRequest) response.Basic
}

type createProductCategoryUseCase struct {
	productCategoryRepository product_category.Repository
}

func (useCase createProductCategoryUseCase) Execute(request CreateProductCategoryRequest) response.Basic {
	productCategory := product_category.ProductCategory{
		Name:  request.Name,
		Image: request.Image,
	}

	useCase.productCategoryRepository.Create(productCategory)

	fmt.Println(productCategory)

	return response.Success()
}

func NewUseCase(productCategoryRepository product_category.Repository) UseCase {
	return &createProductCategoryUseCase{
		productCategoryRepository: productCategoryRepository,
	}
}
