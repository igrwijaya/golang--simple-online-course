package create_product_category

type CreateProductCategoryRequest struct {
	Name  string `json:"name,omitempty" binding:"required,min=3"`
	Image string `json:"image,omitempty" binding:"required"`
}
