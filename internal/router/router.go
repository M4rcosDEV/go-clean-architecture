package router

import (
	mercadopago "github.com/M4rcosDEV/learningGO/internal/http/mercado_pago"
	"github.com/M4rcosDEV/learningGO/internal/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	handler := mercadopago.NewHandler()

	// Register all routes here version wise
	{
		v1 := router.Group("/api/v1")

		// User routes
		user.RegisterUserModule(v1, db)

		// Mercado Pago Routes
		mercadopago.RegisterMercadoPagoRoutes(v1, handler)
	}

	return router
}