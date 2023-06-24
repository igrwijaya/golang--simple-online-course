package forgot_password

import (
	"golang-online-course/pkg/entity/user"
	"gorm.io/gorm"
	"time"
)

type ForgotPassword struct {
	Id        int
	UserId    int
	Valid     int8
	Code      string
	ExpiredAt *time.Time
	CreatedBy int
	CreatedAt *time.Time
	UpdatedBy int
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt

	User user.User
}
