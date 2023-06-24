package entity

import (
	"golang-online-course/pkg/db"
	"golang-online-course/pkg/entity/admin"
	"golang-online-course/pkg/entity/cart"
	"golang-online-course/pkg/entity/class_room"
	"golang-online-course/pkg/entity/discount"
	"golang-online-course/pkg/entity/forgot_password"
	"golang-online-course/pkg/entity/oauth_access_token"
	"golang-online-course/pkg/entity/oauth_client"
	"golang-online-course/pkg/entity/oauth_refresh_token"
	"golang-online-course/pkg/entity/order"
	"golang-online-course/pkg/entity/order_detail"
	"golang-online-course/pkg/entity/product"
	"golang-online-course/pkg/entity/product_category"
	"golang-online-course/pkg/entity/user"
)

func MigrateEntity() {
	appDb := db.NewAppDb()

	adminRepo := admin.NewRepository(appDb)
	cartRepo := cart.NewRepository(appDb)
	classRoomRepo := class_room.NewRepository(appDb)
	discountRepo := discount.NewRepository(appDb)
	forgotPasswordRepo := forgot_password.NewRepository(appDb)
	oauthAccessTokenRepo := oauth_access_token.NewRepository(appDb)
	oauthClientRepo := oauth_client.NewRepository(appDb)
	oauthRefreshTokenRepo := oauth_refresh_token.NewRepository(appDb)
	orderRepo := order.NewRepository(appDb)
	orderDetailRepo := order_detail.NewRepository(appDb)
	productRepo := product.NewRepository(appDb)
	productCategoryRepo := product_category.NewRepository(appDb)
	userRepo := user.NewRepository(appDb)

	adminRepo.Migrate()
	cartRepo.Migrate()
	classRoomRepo.Migrate()
	discountRepo.Migrate()
	forgotPasswordRepo.Migrate()
	oauthAccessTokenRepo.Migrate()
	oauthClientRepo.Migrate()
	oauthRefreshTokenRepo.Migrate()
	orderRepo.Migrate()
	orderDetailRepo.Migrate()
	productRepo.Migrate()
	productCategoryRepo.Migrate()
	userRepo.Migrate()
}
