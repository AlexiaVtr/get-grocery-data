package database

import (
	"database/sql"
	"fmt"
)

func main() {
	// Conectarse a la base de datos
	db, err := sql.Open("postgres")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
}

func Connect() {
	fmt.Println("Example")
}
