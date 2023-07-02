package cart

import (
	"golang-online-course/pkg/entity/core_entity"
	"golang-online-course/pkg/entity/product"
	"golang-online-course/pkg/entity/user"
)

type Cart struct {
	core_entity.CoreEntity
	UserId    int
	ProductId int
	Quantity  int
	IsChecked bool

	User    user.User       `gorm:"foreignKey:UserId;references:Id"`
	Product product.Product `gorm:"foreignKey:ProductId;references:Id"`
}
