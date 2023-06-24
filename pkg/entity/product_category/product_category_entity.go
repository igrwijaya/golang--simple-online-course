package product_category

import (
	"gorm.io/gorm"
	"time"
)

type ProductCategory struct {
	Id        int             `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Image     string          `json:"image,omitempty"`
	CreatedBy int             `json:"created_by,omitempty"`
	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedBy int             `json:"updated_by,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}
