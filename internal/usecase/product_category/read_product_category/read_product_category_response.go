package read_product_category

import "golang-online-course/pkg/response"

type ReadProductCategoryResponse struct {
	response.Basic
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
}
