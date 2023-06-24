package discount

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
}

type discountRepository struct {
	db *gorm.DB
}

func (repo *discountRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&Discount{})

	if autoMigrateError != nil {
		panic("Can't migrate Discount entity. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &discountRepository{
		db: appDb.UseMysql(),
	}
}
