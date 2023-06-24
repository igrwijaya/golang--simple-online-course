package user

import (
	"errors"
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/response"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
	Create(user User) (*int, *response.Error)
	FindByEmail(email string) *User
}

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) FindByEmail(email string) *User {
	var user User

	queryResult := repo.db.Where(&User{Email: email}).First(&user)

	if queryResult.Error != nil && !errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		panic("Cannot find User data by email. " + queryResult.Error.Error())
	}

	if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}

func (repo *userRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&User{})

	if autoMigrateError != nil {
		panic("Can't migrate User entity. " + autoMigrateError.Error())
	}
}

func (repo *userRepository) Create(user User) (*int, *response.Error) {
	createUserResult := repo.db.Create(&user)

	if createUserResult.Error != nil {
		return nil, &response.Error{
			Code:  500,
			Error: createUserResult.Error,
		}
	}

	return &user.Id, nil
}

func NewRepository(appDb db.AppDb) Repository {
	return &userRepository{
		db: appDb.UseMysql(),
	}
}
