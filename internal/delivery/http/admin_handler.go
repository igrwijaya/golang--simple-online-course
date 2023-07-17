package http

import (
	"github.com/gin-gonic/gin"
	"golang-online-course/internal/usecase/admin/create_admin"
	"golang-online-course/internal/usecase/admin/delete_admin"
	"golang-online-course/internal/usecase/admin/get_admin"
	"golang-online-course/internal/usecase/admin/login_admin"
	"golang-online-course/internal/usecase/admin/read_admin"
	"golang-online-course/internal/usecase/admin/update_admin"
	"golang-online-course/pkg/response"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	CreateAdminUseCase create_admin.UseCase
	ReadAdminUseCase   read_admin.UseCase
	UpdateAdminUseCase update_admin.UseCase
	DeleteAdminUseCase delete_admin.UseCase
	GetAdminUseCase    get_admin.UseCase
	LoginAdminUseCase  login_admin.UseCase
}

func NewAdminHandler(
	createAdminUseCase create_admin.UseCase,
	readAdminUseCase read_admin.UseCase,
	updateAdminUseCase update_admin.UseCase,
	deleteAdminUseCase delete_admin.UseCase,
	getAdminUseCase get_admin.UseCase,
	loginAdminUseCase login_admin.UseCase,
) *AdminHandler {
	return &AdminHandler{
		CreateAdminUseCase: createAdminUseCase,
		ReadAdminUseCase:   readAdminUseCase,
		UpdateAdminUseCase: updateAdminUseCase,
		DeleteAdminUseCase: deleteAdminUseCase,
		GetAdminUseCase:    getAdminUseCase,
		LoginAdminUseCase:  loginAdminUseCase,
	}
}

func (handler *AdminHandler) Route(route *gin.RouterGroup) {
	apiRouter := route.Group("/api/v1/admins")

	apiRouter.POST("", handler.Create)
	apiRouter.GET("/:id", handler.Read)
	apiRouter.PUT("/:id", handler.Update)
	apiRouter.DELETE("/:id", handler.Delete)
	apiRouter.GET("", handler.Get)

	apiRouter.POST("/login", handler.Login)
}

func (handler *AdminHandler) Create(context *gin.Context) {
	var createRequest create_admin.CreateAdminRequest

	errParsingJson := context.ShouldBindJSON(&createRequest)

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errParsingJson.Error()))
		context.Abort()
		return
	}

	createResponse := handler.CreateAdminUseCase.Execute(createRequest)

	if createResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(createResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, "Success Execute Admin")
}

func (handler *AdminHandler) Read(context *gin.Context) {
	id, errConvertParam := strconv.Atoi(context.Param("id"))

	if errConvertParam != nil || id <= 0 {
		context.JSON(http.StatusBadRequest, response.BadRequest("invalid path param"))
		context.Abort()
		return
	}

	readResponse := handler.ReadAdminUseCase.Execute(uint(id))

	if readResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(readResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, readResponse)
}

func (handler *AdminHandler) Update(context *gin.Context) {
	id, errConvertParam := strconv.Atoi(context.Param("id"))

	if errConvertParam != nil || id <= 0 {
		context.JSON(http.StatusBadRequest, response.BadRequest("invalid path param"))
		context.Abort()
		return
	}

	var updateRequest update_admin.UpdateAdminRequest

	errParsingJson := context.ShouldBindJSON(&updateRequest)

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errParsingJson.Error()))
		context.Abort()
		return
	}

	updateResponse := handler.UpdateAdminUseCase.Execute(uint(id), updateRequest)

	if updateResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(updateResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, "Success Execute Admin")
}

func (handler *AdminHandler) Delete(context *gin.Context) {
	id, errParsingParam := strconv.Atoi(context.Param("id"))

	if errParsingParam != nil || id <= 0 {
		context.JSON(http.StatusBadRequest, response.BadRequest("invalid path param"))
		context.Abort()
		return
	}

	deleteResponse := handler.DeleteAdminUseCase.Execute(uint(id))

	if deleteResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(deleteResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, "Success Execute Admin")
}

func (handler *AdminHandler) Get(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	size, _ := strconv.Atoi(context.Query("size"))

	if page <= 0 {
		page = 1
	}

	// apply -1 as unlimited
	if size == 0 || size < -1 {
		size = 10
	}

	getResponse := handler.GetAdminUseCase.Execute(page, size)

	if getResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(getResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, getResponse)
}

func (handler *AdminHandler) Login(context *gin.Context) {
	var loginRequest login_admin.LoginAdminRequest

	errParsingJson := context.ShouldBindJSON(&loginRequest)

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errParsingJson.Error()))
		context.Abort()
		return
	}

	loginResponse, errLogin := handler.LoginAdminUseCase.Login(loginRequest)

	if errLogin != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errLogin.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, loginResponse)
}
