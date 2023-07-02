package oauth_refresh_token

import (
	"golang-online-course/pkg/entity/core_entity"
	"golang-online-course/pkg/entity/oauth_access_token"
	"golang-online-course/pkg/entity/user"
	"time"
)

type OauthRefreshToken struct {
	core_entity.CoreEntity
	OauthAccessTokenId uint
	UserId             uint
	Token              string
	ExpiredAt          *time.Time

	OauthAccessToken oauth_access_token.OauthAccessToken `gorm:"foreignKey:OauthAccessTokenId;references:Id"`
	User             user.User                           `gorm:"foreignKey:UserId;references:Id"`
}
