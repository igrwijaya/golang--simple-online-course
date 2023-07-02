package oauth_client

import (
	"errors"
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity/core_entity"
	"gorm.io/gorm"
)

type Repository interface {
	core_entity.BaseRepository
	FindClient(clientId string, clientSecret string) *OauthClient
}

type oauthClientRepository struct {
	db *gorm.DB
}

func (repo *oauthClientRepository) FindClient(clientId string, clientSecret string) *OauthClient {
	var oauthClient OauthClient

	queryResult := repo.db.Where(&OauthClient{ClientId: clientId, ClientSecret: clientSecret}).First(&oauthClient)

	if queryResult.Error != nil && !errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		panic("Cannot Find OAuth Client data. " + queryResult.Error.Error())
	}

	if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &oauthClient
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
