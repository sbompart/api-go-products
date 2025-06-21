package handlers

import (
	service "api-products/internal/service/products"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *service.ProductService
}

func NewProductHandler(s *service.ProductService, r *gin.Engine) *ProductHandler {
	h := &ProductHandler{Service: s}
	group := r.Group("/v1/products")
	{
		group.POST("/", h.CreateProduct)
		group.GET("/", h.GetAll)
		// add PUT, DELETE...
	}
	return h
}
