package product_category

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
}

type productCategoryRepository struct {
	db *gorm.DB
}

func (repo *productCategoryRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&ProductCategory{})

	if autoMigrateError != nil {
		panic("Can't migrate Product Category entity. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &productCategoryRepository{
		db: appDb.UseMysql(),
	}
}
