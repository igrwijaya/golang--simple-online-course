package forgot_password

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
}

type forgotPasswordRepository struct {
	db *gorm.DB
}

func (repo *forgotPasswordRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&ForgotPassword{})

	if autoMigrateError != nil {
		panic("Can't migrate Forgot Password entity. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &forgotPasswordRepository{
		db: appDb.UseMysql(),
	}
}
