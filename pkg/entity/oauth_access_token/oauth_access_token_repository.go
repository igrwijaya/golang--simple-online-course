package oauth_access_token

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
}

type oauthAccessTokenRepository struct {
	db *gorm.DB
}

func (repo *oauthAccessTokenRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&OauthAccessToken{})

	if autoMigrateError != nil {
		panic("Can't migrate OAuth Access Token. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &oauthAccessTokenRepository{
		db: appDb.UseMysql(),
	}
}
