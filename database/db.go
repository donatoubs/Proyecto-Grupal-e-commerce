// database/db.go
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB es la conexión a la base de datos global
var DB *sql.DB

// InitDB inicializa la conexión a la base de datos
func InitDB() error {
	dsn := "root:donato123@tcp(localhost:3306)/ecommerce" // Cambia estos valores por los correctos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	// Verificamos la conexión
	if err = db.Ping(); err != nil {
		return fmt.Errorf("error al verificar la conexión: %w", err)
	}

	DB = db
	return nil
}
