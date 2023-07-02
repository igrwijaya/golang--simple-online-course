package discount

import (
	"golang-online-course/pkg/entity/core_entity"
	"time"
)

type Discount struct {
	core_entity.CoreEntity
	Name              string
	Code              string
	Quantity          int
	RemainingQuantity int
	Type              string
	Value             int
	StartDate         *time.Time
	EndDate           *time.Time
}
