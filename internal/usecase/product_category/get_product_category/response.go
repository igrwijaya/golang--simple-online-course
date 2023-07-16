package get_product_category

import "golang-online-course/pkg/response"

type ProductCategoryDto struct {
	Id    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
}

type GetProductCategoryResponse struct {
	response.Pagination
	Data []ProductCategoryDto
}
