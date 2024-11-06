// pkg/inventory/inventory.go
package inventory

import (
	"ecommerce-go/database"
	"log"
)

// Estructura del artículo del inventario
type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Obtener todos los artículos del inventario
func GetAllItems() ([]Item, error) {
	rows, err := database.DB.Query("SELECT id, name, quantity, price FROM inventory")
	if err != nil {
		log.Println("Error retrieving inventory:", err)
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Name, &i.Quantity, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

// Agregar un artículo al inventario
func AddItem(i Item) error {
	_, err := database.DB.Exec("INSERT INTO inventory (name, quantity, price) VALUES (?, ?, ?)", i.Name, i.Quantity, i.Price)
	if err != nil {
		log.Println("Error adding item:", err)
	}
	return err
}
