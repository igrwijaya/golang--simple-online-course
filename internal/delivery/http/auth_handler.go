package http

import (
	"github.com/gin-gonic/gin"
	"golang-online-course/internal/usecase/auth"
	authDto "golang-online-course/internal/usecase/auth/dto"
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
	route.POST("/api/v1/auth/register", handler.Register)
}

func (handler *AuthHandler) Register(context *gin.Context) {
	var registerInput authDto.RegisterRequestDto

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
