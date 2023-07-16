package main

import (
	"github.com/gin-gonic/gin"
	"golang-online-course/internal/delivery/http"
	"golang-online-course/internal/usecase/admin/create_admin"
	"golang-online-course/internal/usecase/admin/delete_admin"
	"golang-online-course/internal/usecase/admin/get_admin"
	"golang-online-course/internal/usecase/admin/login_admin"
	"golang-online-course/internal/usecase/admin/read_admin"
	"golang-online-course/internal/usecase/admin/update_admin"
	"golang-online-course/internal/usecase/auth"
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity"
	"golang-online-course/pkg/entity/admin"
	"golang-online-course/pkg/entity/forgot_password"
	"golang-online-course/pkg/entity/oauth_access_token"
	"golang-online-course/pkg/entity/oauth_client"
	"golang-online-course/pkg/entity/oauth_refresh_token"
	"golang-online-course/pkg/entity/user"
	"golang-online-course/pkg/service/email_service"
)

func main() {
	entity.MigrateEntity()

	route := gin.Default()
	appDb := db.NewAppDb()

	oauthAccessTokenRepo := oauth_access_token.NewRepository(appDb)
	oauthClientRepo := oauth_client.NewRepository(appDb)
	oauthRefreshTokenRepo := oauth_refresh_token.NewRepository(appDb)
	userRepo := user.NewRepository(appDb)
	forgotPasswordRepo := forgot_password.NewRepository(appDb)
	adminRepo := admin.NewRepository(appDb)

	mailService := email_service.NewService()

	authUseCase := auth.NewUseCase(
		userRepo,
		oauthClientRepo,
		oauthAccessTokenRepo,
		oauthRefreshTokenRepo,
		mailService,
		forgotPasswordRepo)

	createAdminUseCase := create_admin.NewCreateAdminUseCase(adminRepo)
	readAdminUseCase := read_admin.NewReadAdminUseCase(adminRepo)
	updateAdminUseCase := update_admin.NewUpdateAdminUseCase(adminRepo)
	deleteAdminUseCase := delete_admin.NewDeleteAdminUseCase(adminRepo)
	getAdminUseCase := get_admin.NewGetAdminUseCase(adminRepo)
	loginAdminUseCase := login_admin.NewUseCase(
		adminRepo,
		oauthClientRepo,
		oauthAccessTokenRepo,
		oauthRefreshTokenRepo)

	authHandler := http.NewAuthHandler(authUseCase)
	authHandler.Route(&route.RouterGroup)

	adminHandler := http.NewAdminHandler(
		createAdminUseCase,
		readAdminUseCase,
		updateAdminUseCase,
		deleteAdminUseCase,
		getAdminUseCase,
		loginAdminUseCase)
	adminHandler.Route(&route.RouterGroup)

	errRun := route.Run()

	if errRun != nil {
		panic("Cannot start app. " + errRun.Error())
	}
}
