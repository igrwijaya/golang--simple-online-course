package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang-online-course/pkg/entity/forgot_password"
	"golang-online-course/pkg/entity/oauth_access_token"
	"golang-online-course/pkg/entity/oauth_client"
	"golang-online-course/pkg/entity/oauth_refresh_token"
	"golang-online-course/pkg/entity/user"
	"golang-online-course/pkg/response"
	"golang-online-course/pkg/service/email_service"
	"golang-online-course/pkg/utils"
	"golang-online-course/pkg/utils/jwt_utils"
	"time"
)

type UseCase interface {
	Login(dto LoginRequest) (*LoginResponse, *response.Error)
	Register(dto RegisterRequest) response.Basic
	SendForgotPasswordRequest(request ForgotPasswordRequest) response.Basic
	ResetPassword(request ResetPasswordRequest) response.Basic
}

type authUseCase struct {
	userRepository              user.Repository
	oauthClientRepository       oauth_client.Repository
	oauthAccessTokenRepository  oauth_access_token.Repository
	oauthRefreshTokenRepository oauth_refresh_token.Repository
	emailService                email_service.Service
	forgotPasswordRepository    forgot_password.Repository
}

func NewUseCase(
	userRepository user.Repository,
	oauthClientRepository oauth_client.Repository,
	oauthAccessTokenRepository oauth_access_token.Repository,
	oauthRefreshTokenRepository oauth_refresh_token.Repository,
	mailService email_service.Service,
	forgotPasswordRepository forgot_password.Repository) UseCase {

	return &authUseCase{
		userRepository:              userRepository,
		oauthClientRepository:       oauthClientRepository,
		oauthAccessTokenRepository:  oauthAccessTokenRepository,
		oauthRefreshTokenRepository: oauthRefreshTokenRepository,
		emailService:                mailService,
		forgotPasswordRepository:    forgotPasswordRepository,
	}
}

func (useCase *authUseCase) Login(requestDto LoginRequest) (*LoginResponse, *response.Error) {
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

	return &LoginResponse{
		AccessToken:  oauthAccessTokenEntity.Token,
		RefreshToken: oauthRefreshTokenEntity.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

func (useCase *authUseCase) Register(requestDto RegisterRequest) response.Basic {

	if requestDto.Password != requestDto.ConfirmPassword {
		return response.Basic{
			Code:  400,
			Error: errors.New("password confirmation doesn't match"),
		}
	}

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

	emailVerificationDto := email_service.EmailVerificationRequest{
		Subject:          "Online Course Registration Verification",
		Email:            requestDto.Email,
		VerificationCode: newUser.CodeVerified,
	}

	sendEmailResponse := useCase.emailService.SendVerification(emailVerificationDto)

	if sendEmailResponse.Error != nil {
		return sendEmailResponse
	}

	return response.Success()
}

func (useCase *authUseCase) SendForgotPasswordRequest(request ForgotPasswordRequest) response.Basic {
	existingUser := useCase.userRepository.FindByEmail(request.Email)

	if existingUser == nil {
		return response.Success()
	}

	expiredRequest := time.Now().Add(time.Hour)

	forgotPasswordEntity := forgot_password.ForgotPassword{
		UserId:    existingUser.Id,
		Code:      utils.RandString(32),
		ExpiredAt: &expiredRequest,
	}

	useCase.forgotPasswordRepository.Create(forgotPasswordEntity)

	emailRequest := email_service.ForgotPasswordRequest{
		Subject: "Forgot Password",
		Email:   existingUser.Email,
		Code:    forgotPasswordEntity.Code,
	}

	sendForgotPassResponse := useCase.emailService.SendForgotPassword(emailRequest)

	if sendForgotPassResponse.Error != nil {
		return sendForgotPassResponse
	}

	return response.Success()
}

func (useCase *authUseCase) ResetPassword(request ResetPasswordRequest) response.Basic {
	if request.NewPassword != request.ConfirmNewPassword {
		return response.Basic{
			Code:  400,
			Error: errors.New("password confirmation doesn't match"),
		}
	}

	forgotPassEntity := useCase.forgotPasswordRepository.FindByCode(request.Code)

	if forgotPassEntity == nil {
		return response.Basic{
			Code:  400,
			Error: errors.New("invalid forgot password verification code"),
		}
	}

	if forgotPassEntity.ExpiredAt != nil {
		expiredAt := forgotPassEntity.ExpiredAt.UTC()

		if time.Now().UTC().After(expiredAt) {
			return response.Basic{
				Code:  400,
				Error: errors.New("invalid forgot password verification code"),
			}
		}
	}

	hashPassword := utils.EncryptPassword(request.NewPassword)

	useCase.userRepository.ChangePassword(forgotPassEntity.UserId, hashPassword)
	useCase.forgotPasswordRepository.Delete(forgotPassEntity.Id)

	return response.Success()
}
