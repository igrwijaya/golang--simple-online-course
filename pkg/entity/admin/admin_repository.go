package admin

import (
	"errors"
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity/core_entity"
	"gorm.io/gorm"
)

type Repository interface {
	core_entity.BaseRepository
	Create(entity Admin)
	Read(id uint) *Admin
	Update(id uint, name string, email string)
	Delete(id uint)
	Get(page int, size int) (int64, []Admin)
	FindByEmail(email string) *Admin
}

type adminRepository struct {
	db *gorm.DB
}

func (repo *adminRepository) FindByEmail(email string) *Admin {
	var entity Admin

	findQueryResult := repo.db.Where(&Admin{Email: email}).First(&entity)

	if findQueryResult.Error == nil {
		return &entity
	}

	if findQueryResult.Error != nil && errors.Is(findQueryResult.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	panic("cannot find admin data by email. " + findQueryResult.Error.Error())
}

func (repo *adminRepository) Get(page int, size int) (int64, []Admin) {
	var entities []Admin
	var totalRecord int64

	getEntitiesResult := repo.db.Scopes(db.Paginate(page, size)).Find(&entities)
	countEntityResult := repo.db.Model(&Admin{}).Count(&totalRecord)

	if getEntitiesResult.Error == nil && countEntityResult.Error == nil {
		return totalRecord, entities
	}

	if getEntitiesResult.Error != nil {
		panic("Cannot get Admin pagination data. " + getEntitiesResult.Error.Error())
	}

	panic("Cannot get Admin pagination data. " + countEntityResult.Error.Error())
}

func (repo *adminRepository) Create(entity Admin) {
	createEntityResult := repo.db.Create(&entity)

	if createEntityResult.Error != nil {
		panic("Cannot create Admin. " + createEntityResult.Error.Error())
	}
}

func (repo *adminRepository) Read(id uint) *Admin {
	var entity Admin

	readEntityResult := repo.db.First(&entity, id)

	if readEntityResult.Error == nil {
		return &entity
	}

	if readEntityResult.Error != nil && errors.Is(readEntityResult.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	panic("cannot read Admin. " + readEntityResult.Error.Error())
}

func (repo *adminRepository) Update(id uint, name string, email string) {
	existingAdmin := repo.Read(id)

	if existingAdmin == nil {
		panic("cannot update admin (not found)")
	}

	existingAdmin.Name = name
	existingAdmin.Email = email

	updateEntityResult := repo.db.Save(&existingAdmin)

	if updateEntityResult.Error != nil {
		panic("Cannot save Admin data. " + updateEntityResult.Error.Error())
	}
}

func (repo *adminRepository) Delete(id uint) {
	entity := repo.Read(id)

	if entity == nil {
		panic("Execute Admin data not found")
	}

	deleteEntityResult := repo.db.Delete(&entity)

	if deleteEntityResult.Error != nil {
		panic("Cannot delete Admin data. " + deleteEntityResult.Error.Error())
	}
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
