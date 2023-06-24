package user

import (
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/response"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
	Create(user User) (*int, *response.Error)
	FindByEmail(email string) (*User, *response.Error)
}

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) FindByEmail(email string) (*User, *response.Error) {
	var user User

	queryResult := repo.db.Where(&User{Email: email}).First(&user)

	if queryResult.Error != nil {
		return nil, &response.Error{
			Code:  404,
			Error: queryResult.Error,
		}
	}

	return &user, nil
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
