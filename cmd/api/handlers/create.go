package handlers

import (
	"net/http"
	"time"

	"api-products/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.CreationTime = time.Now().UTC()

	id, err := h.Service.Create(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No disponible"})
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
