package admin

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
}

type adminRepository struct {
	db *gorm.DB
}

func (repo *adminRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&Admin{})

	if autoMigrateError != nil {
		panic("Can't migrate Admin entity. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &adminRepository{
		db: appDb.UseMysql(),
	}
}
