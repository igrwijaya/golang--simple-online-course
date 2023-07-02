package class_room

import (
	"golang-online-course/pkg/entity/core_entity"
	"golang-online-course/pkg/entity/product"
	"golang-online-course/pkg/entity/user"
)

type ClassRoom struct {
	core_entity.CoreEntity
	UserId    uint
	ProductId uint

	User    user.User       `gorm:"foreignKey:UserId;references:Id"`
	Product product.Product `gorm:"foreignKey:ProductId;references:Id"`
}
