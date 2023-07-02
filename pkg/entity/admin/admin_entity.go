package admin

import (
	"golang-online-course/pkg/entity/core_entity"
)

type Admin struct {
	core_entity.CoreEntity
	Name     string
	Email    string
	Password string
}
