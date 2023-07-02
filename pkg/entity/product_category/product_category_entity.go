package product_category

import (
	"golang-online-course/pkg/entity/core_entity"
)

type ProductCategory struct {
	core_entity.CoreEntity
	Name  string
	Image string
}
