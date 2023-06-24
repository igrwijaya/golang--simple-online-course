package class_room

import (
	"golang-online-course/pkg/entity/product"
	"golang-online-course/pkg/entity/user"
	"gorm.io/gorm"
	"time"
)

type ClassRoom struct {
	Id        int
	UserId    int
	ProductId int
	CreatedBy int
	CreatedAt *time.Time
	UpdatedBy int
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt

	User    user.User
	Product product.Product
}
