package oauth_access_token

import (
	"golang-online-course/pkg/entity/core_entity"
	"golang-online-course/pkg/entity/oauth_client"
	"golang-online-course/pkg/entity/user"
	"time"
)

type OauthAccessToken struct {
	core_entity.CoreEntity
	OauthClientId uint
	UserId        uint
	Token         string
	Scope         string
	ExpiredAt     *time.Time

	OauthClient oauth_client.OauthClient `gorm:"foreignKey:OauthClientId;references:Id"`
	User        user.User                `gorm:"foreignKey:UserId;references:Id"`
}
