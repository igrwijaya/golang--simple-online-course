package class_room

import (
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity/core_entity"
	"gorm.io/gorm"
)

type Repository interface {
	core_entity.BaseRepository
}

type classRoomRepository struct {
	db *gorm.DB
}

func (repo *classRoomRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&ClassRoom{})

	if autoMigrateError != nil {
		panic("Can't migrate Class Room entity. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &classRoomRepository{
		db: appDb.UseMysql(),
	}
}
