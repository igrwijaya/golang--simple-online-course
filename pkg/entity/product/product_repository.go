package product

import (
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity/core_entity"
	"gorm.io/gorm"
)

type Repository interface {
	core_entity.BaseRepository
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
