package http

import (
	"github.com/gin-gonic/gin"
	"golang-online-course/internal/usecase/auth"
	"golang-online-course/pkg/response"
	"net/http"
)

type AuthHandler struct {
	authUseCase auth.UseCase
}

func NewAuthHandler(authUseCase auth.UseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (handler *AuthHandler) Route(route *gin.RouterGroup) {
	apiRouter := route.Group("/api/v1")

	apiRouter.POST("/auth/register", handler.Register)
	apiRouter.POST("/auth/login", handler.Login)
	apiRouter.POST("/auth/forgot-password", handler.SendForgotPasswordRequest)
	apiRouter.POST("/auth/reset-password", handler.ResetPassword)
	apiRouter.POST("/auth/refresh", handler.RefreshToken)
}

func (handler *AuthHandler) Register(context *gin.Context) {
	var registerInput auth.RegisterRequest

	errParsingJson := context.ShouldBindJSON(&registerInput)

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errParsingJson.Error()))
		context.Abort()
		return
	}

	registerResponse := handler.authUseCase.Register(registerInput)

	if registerResponse.Error != nil {
		context.JSON(int(registerResponse.Code), response.BadRequest(registerResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, "OK")
}

func (handler *AuthHandler) Login(context *gin.Context) {
	var loginInput auth.LoginRequest

	errParsingJson := context.ShouldBindJSON(&loginInput)

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errParsingJson.Error()))
		context.Abort()
		return
	}

	loginResponse, errLogin := handler.authUseCase.Login(loginInput)

	if errLogin != nil {
		context.JSON(int(errLogin.Code), response.BadRequest(errLogin.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, loginResponse)
}

func (handler *AuthHandler) SendForgotPasswordRequest(context *gin.Context) {
	var forgotPassInput auth.ForgotPasswordRequest

	errParsingJson := context.ShouldBindJSON(&forgotPassInput)

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errParsingJson.Error()))
		context.Abort()
		return
	}

	sendForgotPassResponse := handler.authUseCase.SendForgotPasswordRequest(forgotPassInput)

	if sendForgotPassResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(sendForgotPassResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, "OK")
}

func (handler *AuthHandler) ResetPassword(context *gin.Context) {
	var resetPassInput auth.ResetPasswordRequest

	errPassingJson := context.ShouldBindJSON(&resetPassInput)

	if errPassingJson != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errPassingJson.Error()))
		context.Abort()
		return
	}

	resetPassResponse := handler.authUseCase.ResetPassword(resetPassInput)

	if resetPassResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(resetPassResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, "OK")
}

func (handler *AuthHandler) RefreshToken(context *gin.Context) {
	var request auth.RefreshAuthTokenRequest

	errParsingJson := context.ShouldBindJSON(&request)

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errParsingJson.Error()))
		context.Abort()
		return
	}

	authToken, errResponse := handler.authUseCase.Refresh(request)

	if errResponse != nil || authToken == nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, authToken)
}
