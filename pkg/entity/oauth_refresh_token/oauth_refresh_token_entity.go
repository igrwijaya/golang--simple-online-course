package oauth_refresh_token

import (
	"golang-online-course/pkg/entity/oauth_access_token"
	"golang-online-course/pkg/entity/user"
	"gorm.io/gorm"
	"time"
)

type OauthRefreshToken struct {
	Id                 int
	OauthAccessTokenId int
	UserId             int
	Token              string
	ExpiredAt          *time.Time
	CreatedBy          int
	CreatedAt          *time.Time
	UpdatedBy          int
	UpdatedAt          *time.Time
	DeletedAt          *gorm.DeletedAt

	OauthAccessToken oauth_access_token.OauthAccessToken
	User             user.User
}
