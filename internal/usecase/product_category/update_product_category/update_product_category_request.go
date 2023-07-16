package update_product_category

type UpdateProductCategoryBodyRequest struct {
	Name  string `json:"name,omitempty" binding:"required,mix=3"`
	Image string `json:"image,omitempty" binding:"required"`
}
