package route

import (
	"AlurBayar/internal/product/delivery/http"

	"github.com/gin-gonic/gin"
)

// MapProductRoutes mendaftarkan semua endpoint produk
func MapProductRoutes(r *gin.Engine, h http.ProductHandler) {
	productGroup := r.Group("/api/v1")
	{
		productGroup.GET("/products/:productID", h.GetProduct)
	}
}
