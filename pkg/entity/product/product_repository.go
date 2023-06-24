package product

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
}

type productRepository struct {
	db *gorm.DB
}

func (repo *productRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&Product{})

	if autoMigrateError != nil {
		panic("Can't migrate Product entity. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &productRepository{
		db: appDb.UseMysql(),
	}
}
