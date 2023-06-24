package order_detail

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
}

type orderDetailRepository struct {
	db *gorm.DB
}

func (repo *orderDetailRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&OrderDetail{})

	if autoMigrateError != nil {
		panic("Can't migrate Order Detail entity. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &orderDetailRepository{
		db: appDb.UseMysql(),
	}
}
