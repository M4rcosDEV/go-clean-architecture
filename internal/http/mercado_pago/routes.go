package mercadopago

import (
	"github.com/gin-gonic/gin"
)

func RegisterMercadoPagoRoutes(router *gin.RouterGroup, handler *Handler) {

	router.GET("mercadopago/oauth/callback", handler.Callback)

}

// func RegisterUserModule(rg *gin.RouterGroup, db *gorm.DB) {
// 	repo := NewRepositoryImpl(db)
// 	service := NewService(repo)
// 	handler := NewHandler(service)

// 	registerUserRoutes(rg, handler)
// }