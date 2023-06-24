package user

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
	Create(user User) int
}

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&User{})

	if autoMigrateError != nil {
		panic("Can't migrate User entity. " + autoMigrateError.Error())
	}
}

func (repo *userRepository) Create(user User) int {
	//TODO implement me

	panic("implement me")
}

func NewRepository(appDb db.AppDb) Repository {
	return &userRepository{
		db: appDb.UseMysql(),
	}
}
