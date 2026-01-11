package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerUserRoutes(router *gin.RouterGroup, handler *Handler) {

	router.GET("/users", handler.GetUsers)

	{
		userGroup := router.Group("/user")
		
		userGroup.GET("/:id", handler.GetUserById)
		userGroup.POST("/", handler.CreateUser)
		userGroup.PUT("/", handler.UpdateUser)
		userGroup.DELETE("/", handler.DeleteUser)
	}
}

func RegisterUserModule(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepositoryImpl(db)
	service := NewService(repo)
	handler := NewHandler(service)

	registerUserRoutes(rg, handler)
}
