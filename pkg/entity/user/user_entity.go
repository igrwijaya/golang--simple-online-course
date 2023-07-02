package user

import (
	"golang-online-course/pkg/entity/core_entity"
	"time"
)

type User struct {
	core_entity.CoreEntity
	Name            string
	Email           string
	Password        string
	CodeVerified    string
	EmailVerifiedAt *time.Time
}
