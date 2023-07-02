package oauth_refresh_token

import (
	"errors"
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity/core_entity"
	"gorm.io/gorm"
)

type Repository interface {
	core_entity.BaseRepository
	Create(entity OauthRefreshToken) uint
	FindByToken(refreshToken string) *OauthRefreshToken
}

type oAuthRefreshTokenRepository struct {
	db *gorm.DB
}

func (repo *oAuthRefreshTokenRepository) FindByToken(refreshToken string) *OauthRefreshToken {
	var refreshTokenEntity OauthRefreshToken

	findQueryResult := repo.db.
		Preload("OauthAccessToken").
		Where(&OauthRefreshToken{Token: refreshToken}).
		First(&refreshTokenEntity)

	if findQueryResult.Error == nil {
		return &refreshTokenEntity
	}

	if errors.Is(findQueryResult.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	panic("Cannot find Refresh Token data. " + findQueryResult.Error.Error())
}

func (repo *oAuthRefreshTokenRepository) Create(entity OauthRefreshToken) uint {
	var existingRefreshToken OauthRefreshToken

	findQueryResult := repo.db.Where(&OauthRefreshToken{UserId: entity.UserId}).First(&existingRefreshToken)

	if findQueryResult.Error == nil {
		repo.db.Delete(&existingRefreshToken)
	}

	createEntityResult := repo.db.Create(&entity)

	if createEntityResult.Error != nil {
		panic("Cannot create OAuth Refresh Token data. " + createEntityResult.Error.Error())
	}

	return entity.Id
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
