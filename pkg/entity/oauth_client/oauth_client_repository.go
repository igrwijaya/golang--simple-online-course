package oauth_client

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
}

type oauthClientRepository struct {
	db *gorm.DB
}

func (repo *oauthClientRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&OauthClient{})

	if autoMigrateError != nil {
		panic("Can't migrate Oauth Client entity. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &oauthClientRepository{
		db: appDb.UseMysql(),
	}
}
