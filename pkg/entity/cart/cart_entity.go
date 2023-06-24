package cart

import (
	"golang-online-course/pkg/entity/product"
	"golang-online-course/pkg/entity/user"
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	Id        int
	UserId    int
	ProductId int
	Quantity  int
	IsChecked bool
	CreatedBy int
	CreatedAt *time.Time
	UpdatedBy int
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt

	User    user.User
	Product product.Product
}
