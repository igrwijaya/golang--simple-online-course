package order

import (
	"golang-online-course/pkg/entity/discount"
	"golang-online-course/pkg/entity/user"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	Id           int
	UserId       int
	DiscountId   int
	CheckoutLink string
	ExternalId   string
	Price        int
	TotalPrice   int
	Status       string
	CreatedBy    int
	CreatedAt    *time.Time
	UpdatedBy    int
	UpdatedAt    *time.Time
	DeletedAt    *gorm.DeletedAt

	User     user.User
	Discount discount.Discount
}
