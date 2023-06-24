package order

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
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
