package oauth_access_token

import (
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity/core_entity"
	"gorm.io/gorm"
)

type Repository interface {
	core_entity.BaseRepository
	Create(entity OauthAccessToken) uint
}

type oauthAccessTokenRepository struct {
	db *gorm.DB
}

func (repo *oauthAccessTokenRepository) Create(entity OauthAccessToken) uint {

	var existingAccessToken OauthAccessToken
	findQueryResult := repo.db.Where(&OauthAccessToken{UserId: entity.UserId}).First(&existingAccessToken)

	if findQueryResult.Error == nil {
		repo.db.Delete(&existingAccessToken)
	}

	createEntityResult := repo.db.Create(&entity)

	if createEntityResult.Error != nil {
		panic("Cannot create OAuth Access Token. " + createEntityResult.Error.Error())
	}

	return entity.Id
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
