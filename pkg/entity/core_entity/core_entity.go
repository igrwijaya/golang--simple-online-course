package core_entity

import (
	"gorm.io/gorm"
	"time"
)

type CoreEntity struct {
	Id        uint `gorm:"primaryKey"`
	CreatedBy int
	CreatedAt *time.Time
	UpdatedBy int
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}
