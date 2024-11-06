// pkg/user/user.go
package user

import (
	"ecommerce-go/database"
	"fmt"
)

// Ejemplo de uso de database.DB
func GetUserByID(id int) error {
	var user User // Define tu estructura de usuario

	err := database.DB.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return fmt.Errorf("error al obtener el usuario: %w", err)
	}

	// Resto de la lógica para manejar el usuario
	return nil
}
