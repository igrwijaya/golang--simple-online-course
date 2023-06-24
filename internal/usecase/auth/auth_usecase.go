package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang-online-course/internal/usecase/auth/dto"
	"golang-online-course/pkg/entity/oauth_access_token"
	"golang-online-course/pkg/entity/oauth_client"
	"golang-online-course/pkg/entity/oauth_refresh_token"
	"golang-online-course/pkg/entity/user"
	"golang-online-course/pkg/response"
	"golang-online-course/pkg/utils"
	"time"
)

type UseCase interface {
	Login(dto dto.LoginRequestDto) (*dto.AuthResponseDto, *response.Error)
	Register(dto dto.RegisterRequestDto) response.Basic
}

type authUseCase struct {
	userRepository              user.Repository
	oauthClientRepository       oauth_client.Repository
	oauthAccessTokenRepository  oauth_access_token.Repository
	oauthRefreshTokenRepository oauth_refresh_token.Repository
}

func NewUseCase(
	userRepository user.Repository,
	oauthClientRepository oauth_client.Repository,
	oauthAccessTokenRepository oauth_access_token.Repository,
	oauthRefreshTokenRepository oauth_refresh_token.Repository) UseCase {

	return &authUseCase{
		userRepository:              userRepository,
		oauthClientRepository:       oauthClientRepository,
		oauthAccessTokenRepository:  oauthAccessTokenRepository,
		oauthRefreshTokenRepository: oauthRefreshTokenRepository,
	}
}

func (useCase *authUseCase) Login(requestDto dto.LoginRequestDto) (*dto.AuthResponseDto, *response.Error) {
	oauthClient := useCase.oauthClientRepository.FindClient(requestDto.ClientId, requestDto.ClientSecret)

	if oauthClient == nil {
		return nil, &response.Error{
			Code:  404,
			Error: errors.New("oauth client not found"),
		}
	}

	existingUser := useCase.userRepository.FindByEmail(requestDto.Email)

	if existingUser == nil {
		return nil, &response.Error{
			Code:  404,
			Error: errors.New("user not found"),
		}
	}

	expirationTime := time.Now().Add(time.Hour)

	userClaims := utils.AppClaims{
		Id:      existingUser.Id,
		Name:    existingUser.Name,
		Email:   existingUser.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	accessToken := utils.CreateJwtToken(userClaims)

	oauthAccessTokenEntity := oauth_access_token.OauthAccessToken{
		OauthClientId: oauthClient.Id,
		UserId:        existingUser.Id,
		Token:         accessToken,
		Scope:         "*",
		ExpiredAt:     &expirationTime,
	}

	createOauthTokenResult := useCase.oauthAccessTokenRepository.Create(oauthAccessTokenEntity)

	refreshTokenExpiredAt := expirationTime.Add(time.Minute * 5)

	oauthRefreshTokenEntity := oauth_refresh_token.OauthRefreshToken{
		OauthAccessTokenId: createOauthTokenResult,
		UserId:             existingUser.Id,
		Token:              utils.RandString(128),
		ExpiredAt:          &refreshTokenExpiredAt,
	}

	useCase.oauthRefreshTokenRepository.Create(oauthRefreshTokenEntity)

	return &dto.AuthResponseDto{
		AccessToken:  oauthAccessTokenEntity.Token,
		RefreshToken: oauthRefreshTokenEntity.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

func (useCase *authUseCase) Register(requestDto dto.RegisterRequestDto) response.Basic {
	existingUser := useCase.userRepository.FindByEmail(requestDto.Email)

	if existingUser != nil {
		return response.Basic{
			Code:  400,
			Error: errors.New("email already registered"),
		}
	}

	newUser := user.User{
		Name:            requestDto.Name,
		Email:           requestDto.Email,
		Password:        utils.EncryptPassword(requestDto.Password),
		CodeVerified:    utils.RandString(32),
		EmailVerifiedAt: nil,
	}

	_, errCreateUser := useCase.userRepository.Create(newUser)

	if errCreateUser != nil {
		return response.Basic{
			Code:  errCreateUser.Code,
			Error: errCreateUser.Error,
		}
	}

	return response.Success()
}
