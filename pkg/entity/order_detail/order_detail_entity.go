package order_detail

import (
	"golang-online-course/pkg/entity/order"
	"golang-online-course/pkg/entity/product"
	"gorm.io/gorm"
	"time"
)

type OrderDetail struct {
	Id        int
	OrderId   int
	ProductId int
	Price     int
	CreatedBy int
	CreatedAt *time.Time
	UpdatedBy int
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt

	Order   order.Order
	Product product.Product
}
