package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your pass"
	dbname   = "your db"
	dbtable = "productos_almacen"
	CATEGORY = "your category to put in the db"
	PARAM    = "The param to get the number of the limit iteration in the website to make scraping"
)

var (
	PSQLInfo     = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	documentPath = fmt.Sprintf(`C:/Temp/%v.csv`, CATEGORY)
	WEBSITE      = "https://www.thesupermarket.com.ar/thespecificcategory"
)

func PutDataInDB() {

	db, err := sql.Open("postgres", PSQLInfo)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf(`COPY %s FROM '%s' WITH (FORMAT csv, HEADER true)`, dbtable, documentPath))

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Data successfully saved in database!")
}
