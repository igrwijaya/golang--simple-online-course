package product

import (
	"golang-online-course/pkg/entity/product_category"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id                int
	ProductCategoryId int
	Title             string
	Image             string
	Video             string
	Description       string
	IsHighlighted     string
	Price             int
	CreatedBy         int
	CreatedAt         *time.Time
	UpdatedBy         int
	UpdatedAt         *time.Time
	DeletedAt         *gorm.DeletedAt

	ProductCategory product_category.ProductCategory
}
