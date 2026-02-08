package mercadopago

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Callback(ctx *gin.Context){
	code := ctx.Query("code")

	if code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "code não encontrado",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Mercado Pago conectado com sucesso!",
		"code":    code,
	})
}