package product_category

import (
	"errors"
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity/core_entity"
	"gorm.io/gorm"
)

type Repository interface {
	core_entity.BaseRepository
	Create(entity ProductCategory)
	Read(id uint) *ProductCategory
	Update(id uint, name string, image string)
	Delete(id uint)
	Get(page int, size int) (int64, []ProductCategory)
}

type productCategoryRepository struct {
	db *gorm.DB
}

func (repo *productCategoryRepository) Create(entity ProductCategory) {
	createEntityResult := repo.db.Create(&entity)

	if createEntityResult.Error != nil {
		panic("cannot create Product Category data. " + createEntityResult.Error.Error())
	}
}

func (repo *productCategoryRepository) Read(id uint) *ProductCategory {
	var productCategory ProductCategory

	readQueryResult := repo.db.First(id, &productCategory)

	if readQueryResult.Error == nil {
		return &productCategory
	}

	if readQueryResult.Error != nil && errors.Is(readQueryResult.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	panic("cannot read Product Category data. " + readQueryResult.Error.Error())
}

func (repo *productCategoryRepository) Update(id uint, name string, image string) {
	productCategory := repo.Read(id)

	if productCategory == nil {
		panic("cannot update Product Category. Data not found")
	}

	productCategory.Name = name
	productCategory.Image = image

	updateQueryResult := repo.db.Save(&productCategory)

	if updateQueryResult.Error != nil {
		panic("cannot update Product Category. " + updateQueryResult.Error.Error())
	}
}

func (repo *productCategoryRepository) Delete(id uint) {
	productCategory := repo.Read(id)

	if productCategory == nil {
		panic("cannot delete Product Category. Data not found")
	}

	deleteQueryResult := repo.db.Delete(&productCategory)

	if deleteQueryResult.Error != nil {
		panic("cannot delete Product Category. " + deleteQueryResult.Error.Error())
	}
}

func (repo *productCategoryRepository) Get(page int, size int) (int64, []ProductCategory) {
	var productCategories []ProductCategory
	var totalRecord int64

	getQueryResult := repo.db.Scopes(db.Paginate(page, size)).Find(&productCategories)
	countQueryResult := repo.db.Model(&ProductCategory{}).Count(&totalRecord)

	if getQueryResult.Error != nil {
		panic("cannot Get Product Category. " + getQueryResult.Error.Error())
	}

	if countQueryResult.Error != nil {
		panic("cannot Get Product Category. " + countQueryResult.Error.Error())
	}

	return totalRecord, productCategories
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
