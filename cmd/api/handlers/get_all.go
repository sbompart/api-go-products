package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No disponible"})
	}

	if products == nil {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusOK, gin.H{"products": products})
	}
}
