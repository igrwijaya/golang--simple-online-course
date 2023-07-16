package login_admin

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang-online-course/pkg/entity/admin"
	"golang-online-course/pkg/entity/oauth_access_token"
	"golang-online-course/pkg/entity/oauth_client"
	"golang-online-course/pkg/entity/oauth_refresh_token"
	"golang-online-course/pkg/response"
	"golang-online-course/pkg/utils"
	"golang-online-course/pkg/utils/jwt_utils"
	"time"
)

type UseCase interface {
	Login(request LoginAdminRequest) (*LoginAdminResponse, *response.Error)
}

type loginAdminUseCase struct {
	adminRepository             admin.Repository
	oauthClientRepository       oauth_client.Repository
	oauthAccessTokenRepository  oauth_access_token.Repository
	oauthRefreshTokenRepository oauth_refresh_token.Repository
}

func (useCase loginAdminUseCase) Login(request LoginAdminRequest) (*LoginAdminResponse, *response.Error) {
	oauthClient := useCase.oauthClientRepository.FindClient(request.ClientId, request.ClientSecret)

	if oauthClient == nil {
		return nil, &response.Error{
			Code:  400,
			Error: errors.New("invalid client"),
		}
	}

	adminEntity := useCase.adminRepository.FindByEmail(request.Email)

	if adminEntity == nil {
		return nil, &response.Error{
			Code:  400,
			Error: errors.New("email or password is invalid"),
		}
	}

	hasValidPassword := jwt_utils.HasValidPassword(adminEntity.Password, request.Password)

	if !hasValidPassword {
		return nil, &response.Error{
			Code:  400,
			Error: errors.New("email or password is invalid"),
		}
	}

	expirationTime := time.Now().UTC().Add(time.Hour)

	userClaims := jwt_utils.AppClaims{
		Id:      adminEntity.Id,
		Name:    adminEntity.Name,
		Email:   adminEntity.Email,
		IsAdmin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	accessToken := jwt_utils.CreateJwtToken(userClaims)

	oauthAccessTokenEntity := oauth_access_token.OauthAccessToken{
		OauthClientId: oauthClient.Id,
		UserId:        adminEntity.Id,
		Token:         accessToken,
		Scope:         "*",
		ExpiredAt:     &expirationTime,
	}

	createAccessTokenResult := useCase.oauthAccessTokenRepository.Create(oauthAccessTokenEntity)

	refreshTokenExpiredAt := time.Now().UTC().Add(time.Hour * 2)
	refreshTokenEntity := oauth_refresh_token.OauthRefreshToken{
		OauthAccessTokenId: createAccessTokenResult,
		UserId:             adminEntity.Id,
		Token:              utils.RandString(128),
		ExpiredAt:          &refreshTokenExpiredAt,
	}

	useCase.oauthRefreshTokenRepository.Create(refreshTokenEntity)

	return &LoginAdminResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenEntity.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime,
		Scope:        "*",
	}, nil
}

func NewUseCase(
	adminRepository admin.Repository,
	oauthClientRepository oauth_client.Repository,
	oauthAccessTokenRepository oauth_access_token.Repository,
	oauthRefreshTokenRepository oauth_refresh_token.Repository) UseCase {

	return &loginAdminUseCase{
		adminRepository:             adminRepository,
		oauthClientRepository:       oauthClientRepository,
		oauthAccessTokenRepository:  oauthAccessTokenRepository,
		oauthRefreshTokenRepository: oauthRefreshTokenRepository,
	}

}
