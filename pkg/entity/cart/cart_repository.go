package cart

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
}

type cartRepository struct {
	db *gorm.DB
}

func (repo *cartRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&Cart{})

	if autoMigrateError != nil {
		panic("Can't migrate Cart entity. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &cartRepository{
		db: appDb.UseMysql(),
	}
}
