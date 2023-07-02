package order_detail

import (
	"golang-online-course/pkg/entity/core_entity"
	"golang-online-course/pkg/entity/order"
	"golang-online-course/pkg/entity/product"
)

type OrderDetail struct {
	core_entity.CoreEntity
	OrderId   uint
	ProductId uint
	Price     int

	Order   order.Order     `gorm:"foreignKey:OrderId;references:Id"`
	Product product.Product `gorm:"foreignKey:ProductId;references:Id"`
}
