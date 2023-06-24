package admin

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedBy int
	CreatedAt *time.Time
	UpdatedBy int
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt
}
