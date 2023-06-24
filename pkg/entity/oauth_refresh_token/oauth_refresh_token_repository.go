package oauth_refresh_token

import (
	"golang-online-course/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate()
}

type oAuthRefreshTokenRepository struct {
	db *gorm.DB
}

func (repo *oAuthRefreshTokenRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&OauthRefreshToken{})

	if autoMigrateError != nil {
		panic("Can't migrate Oauth Refresh Token entity. " + autoMigrateError.Error())
	}
}

func NewRepository(appDb db.AppDb) Repository {
	return &oAuthRefreshTokenRepository{
		db: appDb.UseMysql(),
	}
}
