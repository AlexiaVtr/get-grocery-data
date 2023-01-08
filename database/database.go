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
	password = "2022"
	dbname   = "GroceryWebsite"
	CATEGORY = "Perfumeria"
	PARAM    = "Nrpp"
)

var (
	PSQLInfo     = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	documentPath = fmt.Sprintf(`C:/Temp/%v.csv`, CATEGORY)
	WEBSITE      = "https://www.cotodigital3.com.ar/sitios/cdigi/browse/catalogo-perfumer%C3%ADa/_/N-cblpjz?Nf=product.endDate%7CGTEQ+1.673136E12%7C%7Cproduct.startDate%7CLTEQ+1.673136E12&Nr=AND%28product.sDisp_200%3A1004%2Cproduct.language%3Aespa%C3%B1ol%2COR%28product.siteId%3ACotoDigital%29%29"
)

func PutDataInDB() {

	db, err := sql.Open("postgres", PSQLInfo)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf(`COPY productos_almacen FROM '%s' WITH (FORMAT csv, HEADER true)`, documentPath))

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Data successfully saved in database!")
}
