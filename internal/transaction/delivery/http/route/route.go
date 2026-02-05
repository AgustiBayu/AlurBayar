package route

import (
	"AlurBayar/internal/transaction/delivery/http"

	"github.com/gin-gonic/gin"
)

func MapTransactionRoutes(r *gin.Engine, txHandler *http.TransactionHandlerImpl) {
	api := r.Group("/api/v1")
	{
		api.POST("/checkout", txHandler.Checkout)
		api.POST("/notification", txHandler.HandleNotification)
	}
}
