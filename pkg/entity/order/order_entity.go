package order

import (
	"golang-online-course/pkg/entity/core_entity"
	"golang-online-course/pkg/entity/discount"
	"golang-online-course/pkg/entity/user"
)

type Order struct {
	core_entity.CoreEntity
	UserId       uint
	DiscountId   uint
	CheckoutLink string
	ExternalId   string
	Price        int
	TotalPrice   int
	Status       string

	User     user.User         `gorm:"foreignKey:UserId;references:Id"`
	Discount discount.Discount `gorm:"foreignKey:DiscountId;references:Id"`
}
