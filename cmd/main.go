package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configura la conexión
	dsn := "root:donato123@tcp(127.0.0.1:3306)/ecommerce"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa a la base de datos!")

}

func insertarProducto(db *sql.DB, nombre string, descripcion string, precio float64, stock int) {
	query := `INSERT INTO products (nombre, descripcion, precio, stock) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, nombre, descripcion, precio, stock)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Producto insertado con éxito")
}
