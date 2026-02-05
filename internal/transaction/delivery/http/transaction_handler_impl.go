package http

import (
	"AlurBayar/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandlerImpl struct {
	Usecase domain.TransactionUsecase
}

func NewTransactionHandler(u domain.TransactionUsecase) *TransactionHandlerImpl {
	return &TransactionHandlerImpl{Usecase: u}
}

func (h *TransactionHandlerImpl) Checkout(c *gin.Context) {
	var input struct {
		ProductID int `json:"product_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_id wajib diisi"})
		return
	}
	result, err := h.Usecase.CreateOrder(c.Request.Context(), input.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":     "Checkout berhasil",
		"data":        result,
		"payment_url": "https://app.sandbox.midtrans.com/snap/v2/vtweb/" + result.SnapToken,
	})
}
func (h *TransactionHandlerImpl) HandleNotification(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	err := h.Usecase.ProcessNotification(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "notification processed"})
}
