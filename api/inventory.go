// api/inventory.go
package api

import (
	"ecommerce-go/pkg/inventory"
	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"
)

// Obtener el inventario
func GetInventory(c *gin.Context) {
	items, err := inventory.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving inventory"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func Run() {
	fmt.Println("El servidor ha comenzado a ejecutarse...")
	// Aquí va la lógica de inicialización o inicio del servidor
}

// Agregar un artículo al inventario
func AddInventory(c *gin.Context) {
	var item inventory.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := inventory.AddItem(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding item"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Item added successfully"})
}
