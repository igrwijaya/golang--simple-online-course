package forgot_password

import (
	"golang-online-course/pkg/entity/core_entity"
	"golang-online-course/pkg/entity/user"
	"time"
)

type ForgotPassword struct {
	core_entity.CoreEntity
	UserId    uint
	Code      string
	ExpiredAt *time.Time

	User user.User `gorm:"foreignKey:UserId;references:Id"`
}
