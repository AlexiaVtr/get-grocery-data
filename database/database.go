package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	user         = "postgres"
	password     = "2022"
	dbname       = "GroceryWebsite"
	documentPath = `C:\Temp\Almacen.csv`
)

var PSQLInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func PutDataInDB() {

	db, err := sql.Open("postgres", PSQLInfo)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	if err = db.Ping(); err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf(`COPY productos_almacen FROM '%s' WITH (FORMAT csv, HEADER true)`, documentPath))

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
}
