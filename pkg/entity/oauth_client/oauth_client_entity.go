package oauth_client

import (
	"gorm.io/gorm"
	"time"
)

type OauthClient struct {
	Id           int
	ClientId     string
	ClientSecret string
	Name         string
	Redirect     string
	Description  string
	Scope        string
	CreatedBy    int
	CreatedAt    *time.Time
	UpdatedBy    int
	UpdatedAt    *time.Time
	DeletedAt    *gorm.DeletedAt
}
