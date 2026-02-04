package http

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	GetProduct(c *gin.Context)
}
