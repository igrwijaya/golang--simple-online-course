package forgot_password

import (
	"errors"
	"fmt"
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
	Create(entity ForgotPassword) uint
	FindByCode(code string) *ForgotPassword
	Delete(id uint)
}

type forgotPasswordRepository struct {
	db *gorm.DB
}

func (repo *forgotPasswordRepository) Delete(id uint) {
	var entity ForgotPassword

	findEntityResult := repo.db.First(&entity, id)

	if findEntityResult.Error != nil {
		panic("Deleting Forgot Password data is error. " + findEntityResult.Error.Error())
	}

	deleteEntityResult := repo.db.Delete(&ForgotPassword{}, id)

	if deleteEntityResult.Error != nil {
		panic("Cannot delete Forgot Password data. " + deleteEntityResult.Error.Error())
	}
}

func (repo *forgotPasswordRepository) FindByCode(code string) *ForgotPassword {
	var forgotPasswordEntity ForgotPassword

	findEntityResult := repo.db.Where(&ForgotPassword{Code: code}).First(&forgotPasswordEntity)

	if findEntityResult.Error != nil && !errors.Is(findEntityResult.Error, gorm.ErrRecordNotFound) {
		panic("Cannot find Forgot Password data. " + findEntityResult.Error.Error())
	}

	if errors.Is(findEntityResult.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	fmt.Println(forgotPasswordEntity)

	return &forgotPasswordEntity
}

func (repo *forgotPasswordRepository) Create(entity ForgotPassword) uint {
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
