package order

import (
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity/core_entity"
	"gorm.io/gorm"
)

type Repository interface {
	core_entity.BaseRepository
}

type orderRepository struct {
	db *gorm.DB
}

func (repo *orderRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&Order{})

	if autoMigrateError != nil {
		panic("Can't migrate Order entity. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &orderRepository{
		db: appDb.UseMysql(),
	}
}
