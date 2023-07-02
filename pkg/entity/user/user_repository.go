package user

import (
	"errors"
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity/core_entity"
	"golang-online-course/pkg/response"
	"gorm.io/gorm"
)

type Repository interface {
	core_entity.BaseRepository
	Create(user User) (*uint, *response.Error)
	FindByEmail(email string) *User
	ChangePassword(id uint, hashPassword string)
	Find(id uint) *User
}

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) Find(id uint) *User {
	var user User

	findQueryResult := repo.db.First(&user, id)

	if findQueryResult.Error == nil {
		return &user
	}

	if errors.Is(findQueryResult.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	panic("Cannot find User by Id. " + findQueryResult.Error.Error())
}

func (repo *userRepository) ChangePassword(id uint, hashPassword string) {
	var user User

	queryResult := repo.db.First(&user, id)

	if queryResult.Error != nil {
		panic("Cannot find User data by id. " + queryResult.Error.Error())
	}

	user.Password = hashPassword

	repo.db.Save(&user)
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

func (repo *userRepository) Create(user User) (*uint, *response.Error) {
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
