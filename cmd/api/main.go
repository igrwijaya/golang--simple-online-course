package main

import (
	"github.com/gin-gonic/gin"
	"golang-online-course/internal/delivery/http"
	"golang-online-course/internal/usecase/auth"
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity"
	"golang-online-course/pkg/entity/user"
)

func main() {
	entity.MigrateEntity()

	route := gin.Default()

	appDb := db.NewAppDb()

	userRepo := user.NewRepository(appDb)
	authUseCase := auth.NewUseCase(userRepo)
	authHandler := http.NewAuthHandler(authUseCase)

	authHandler.Route(&route.RouterGroup)

	errRun := route.Run()

	if errRun != nil {
		panic("Cannot start app. " + errRun.Error())
	}
}
