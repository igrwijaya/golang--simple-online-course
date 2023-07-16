package get_product_category

import (
	"golang-online-course/pkg/entity/product_category"
	"golang-online-course/pkg/response"
)

type UseCase interface {
	Execute(page int, size int) GetProductCategoryResponse
}

type getProductCategoryUseCase struct {
	productCategoryRepository product_category.Repository
}

func (useCase getProductCategoryUseCase) Execute(page int, size int) GetProductCategoryResponse {
	var productCategoryDtos []ProductCategoryDto

	totalRecord, productCategories := useCase.productCategoryRepository.Get(page, size)

	for _, productCategory := range productCategories {
		productCategoryDtos = append(productCategoryDtos, ProductCategoryDto{
			Id:    productCategory.Id,
			Name:  productCategory.Name,
			Image: productCategory.Image,
		})
	}

	return GetProductCategoryResponse{
		Pagination: response.Pagination{
			Code:        200,
			TotalRecord: totalRecord,
			CurrentPage: page,
			Limit:       size,
		},
		Data: productCategoryDtos,
	}
}

func NewUseCase(productCategoryRepository product_category.Repository) UseCase {
	return &getProductCategoryUseCase{productCategoryRepository}
}
