package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id              int
	Name            string
	Email           string
	Password        string
	CodeVerified    string
	EmailVerifiedAt *time.Time
	CreatedBy       int
	CreatedAt       *time.Time
	UpdatedBy       int
	UpdatedAt       *time.Time
	DeletedAt       *gorm.DeletedAt
}
