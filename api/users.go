// api/users.go
package api

import (
	"ecommerce-go/pkg/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Obtener todos los usuarios
func GetUsers(c *gin.Context) {
	users, err := user.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Crear un nuevo usuario
func CreateUser(c *gin.Context) {
	var newUser user.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := user.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
