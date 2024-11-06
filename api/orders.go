package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Función de ejemplo para manejar pedidos
func GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Listado de pedidos",
	})
}
