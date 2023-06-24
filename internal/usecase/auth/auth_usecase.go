package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang-online-course/pkg/entity/oauth_access_token"
	"golang-online-course/pkg/entity/oauth_client"
	"golang-online-course/pkg/entity/oauth_refresh_token"
	"golang-online-course/pkg/entity/user"
	"golang-online-course/pkg/response"
	"golang-online-course/pkg/service/mail_service"
	"golang-online-course/pkg/utils"
	"golang-online-course/pkg/utils/jwt_utils"
	"time"
)

type UseCase interface {
	Login(dto LoginRequest) (*AuthResponseDto, *response.Error)
	Register(dto RegisterRequest) response.Basic
}

type authUseCase struct {
	userRepository              user.Repository
	oauthClientRepository       oauth_client.Repository
	oauthAccessTokenRepository  oauth_access_token.Repository
	oauthRefreshTokenRepository oauth_refresh_token.Repository
	mailService                 mail_service.Service
}

func NewUseCase(
	userRepository user.Repository,
	oauthClientRepository oauth_client.Repository,
	oauthAccessTokenRepository oauth_access_token.Repository,
	oauthRefreshTokenRepository oauth_refresh_token.Repository,
	mailService mail_service.Service) UseCase {

	return &authUseCase{
		userRepository:              userRepository,
		oauthClientRepository:       oauthClientRepository,
		oauthAccessTokenRepository:  oauthAccessTokenRepository,
		oauthRefreshTokenRepository: oauthRefreshTokenRepository,
		mailService:                 mailService,
	}
}

func (useCase *authUseCase) Login(requestDto LoginRequest) (*AuthResponseDto, *response.Error) {
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
			Code:  400,
			Error: errors.New("email or password is invalid"),
		}
	}

	hasValidPassword := jwt_utils.HasValidPassword(existingUser.Password, requestDto.Password)

	if !hasValidPassword {
		return nil, &response.Error{
			Code:  400,
			Error: errors.New("email or password is invalid"),
		}
	}

	expirationTime := time.Now().Add(time.Hour)

	userClaims := jwt_utils.AppClaims{
		Id:      existingUser.Id,
		Name:    existingUser.Name,
		Email:   existingUser.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	accessToken := jwt_utils.CreateJwtToken(userClaims)

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

	return &AuthResponseDto{
		AccessToken:  oauthAccessTokenEntity.Token,
		RefreshToken: oauthRefreshTokenEntity.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

func (useCase *authUseCase) Register(requestDto RegisterRequest) response.Basic {
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

	emailVerificationDto := mail_service.EmailVerificationDto{
		Subject:          "Online Course Registration Verification",
		Email:            requestDto.Email,
		VerificationCode: newUser.CodeVerified,
	}

	sendEmailResponse := useCase.mailService.SendVerification(emailVerificationDto)

	if sendEmailResponse.Error != nil {
		return sendEmailResponse
	}

	return response.Success()
}
