package forgot_password

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
	Create(entity ForgotPassword) int
}

type forgotPasswordRepository struct {
	db *gorm.DB
}

func (repo *forgotPasswordRepository) Create(entity ForgotPassword) int {
	createEntityResult := repo.db.Create(&entity)

	if createEntityResult.Error != nil {
		panic("Cannot create Forgot Password data. " + createEntityResult.Error.Error())
	}

	return entity.Id
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
