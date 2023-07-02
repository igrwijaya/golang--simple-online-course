package product

import (
	"golang-online-course/pkg/entity/core_entity"
	"golang-online-course/pkg/entity/product_category"
)

type Product struct {
	core_entity.CoreEntity
	ProductCategoryId uint
	Title             string
	Image             string
	Video             string
	Description       string
	IsHighlighted     string
	Price             int

	ProductCategory product_category.ProductCategory `gorm:"foreignKey:ProductCategoryId;references:Id"`
}
