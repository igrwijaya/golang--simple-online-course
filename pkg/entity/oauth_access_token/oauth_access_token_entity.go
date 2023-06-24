package oauth_access_token

import (
	"golang-online-course/pkg/entity/oauth_client"
	"golang-online-course/pkg/entity/user"
	"gorm.io/gorm"
	"time"
)

type OauthAccessToken struct {
	Id            int
	OauthClientId int
	UserId        int
	Token         string
	Scope         string
	ExpiredAt     *time.Time
	CreatedBy     int
	CreatedAt     *time.Time
	UpdatedBy     int
	UpdatedAt     *time.Time
	DeletedAt     *gorm.DeletedAt

	OauthClient oauth_client.OauthClient
	User        user.User
}
