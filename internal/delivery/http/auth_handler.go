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
}

func (handler *AuthHandler) Register(context *gin.Context) {
	var registerInput auth.RegisterRequest

	errParsingJson := context.ShouldBindJSON(&registerInput)

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.ErrorHttp(http.StatusBadRequest, errParsingJson.Error()))
		context.Abort()
		return
	}

	registerResponse := handler.authUseCase.Register(registerInput)

	if registerResponse.Error != nil {
		context.JSON(int(registerResponse.Code), response.ErrorHttp(registerResponse.Code, registerResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, "OK")
}

func (handler *AuthHandler) Login(context *gin.Context) {
	var loginInput auth.LoginRequest

	errParsingJson := context.ShouldBindJSON(&loginInput)

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.ErrorHttp(http.StatusBadRequest, errParsingJson.Error()))
		context.Abort()
		return
	}

	loginResponse, errLogin := handler.authUseCase.Login(loginInput)

	if errLogin != nil {
		context.JSON(int(errLogin.Code), response.ErrorHttp(errLogin.Code, errLogin.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, loginResponse)
}

func (handler *AuthHandler) SendForgotPasswordRequest(context *gin.Context) {
	var forgotPassInput auth.ForgotPasswordRequest

	errParsingJson := context.ShouldBindJSON(&forgotPassInput)

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.ErrorHttp(http.StatusBadRequest, errParsingJson.Error()))
		context.Abort()
		return
	}

	sendForgotPassResponse := handler.authUseCase.SendForgotPasswordRequest(forgotPassInput)

	if sendForgotPassResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.ErrorHttp(http.StatusBadRequest, sendForgotPassResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, "OK")
}
