package discount

import (
	"gorm.io/gorm"
	"time"
)

type Discount struct {
	Id                int
	Name              string
	Code              string
	Quantity          int
	RemainingQuantity int
	Type              string
	Value             int
	StartDate         *time.Time
	EndDate           *time.Time
	CreatedBy         int
	CreatedAt         *time.Time
	UpdatedBy         int
	UpdatedAt         *time.Time
	DeletedAt         *gorm.DeletedAt
}
