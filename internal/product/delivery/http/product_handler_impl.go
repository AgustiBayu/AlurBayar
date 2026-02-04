package http

import (
	"AlurBayar/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandlerImpl struct {
	productUsecase domain.ProductUsecase
}

func NewProductHandler(u domain.ProductUsecase) ProductHandler {
	return &ProductHandlerImpl{
		productUsecase: u,
	}
}

func (h *ProductHandlerImpl) GetProduct(c *gin.Context) {
	idParam := c.Param("productID")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID produk tidak valid, harus berupa angka",
		})
		return
	}
	product, err := h.productUsecase.GetProduct(c.Request.Context(), productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Produk gagal diambil",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mendapatkan detail produk",
		"data":    product,
	})
}
