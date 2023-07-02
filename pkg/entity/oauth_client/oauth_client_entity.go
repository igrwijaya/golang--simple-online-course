package oauth_client

import (
	"golang-online-course/pkg/entity/core_entity"
)

type OauthClient struct {
	core_entity.CoreEntity
	ClientId     string
	ClientSecret string
	Name         string
	Redirect     string
	Description  string
	Scope        string
}
