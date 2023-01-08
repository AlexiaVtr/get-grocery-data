package main

import (
	"encoding/csv"
	"fmt"
	"get-product-data/database"
	"log"
	"os"
)

func main() {

	defer database.PutDataInDB()

	// abre el archivo .csv para escribir en él
	file, err := os.Create(fmt.Sprintf(`C:/Temp/%s.csv`, database.CATEGORY))
	defer file.Close()

	// crea un nuevo escritor de CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// escribe la cabecera del archivo .csv
	writer.Write([]string{"name", "code", "price", "image_url", "category"})

	//escribe todos los datos obtenidos del website
	writeDocumentAllPages(writer, database.CATEGORY)

	if err != nil {
		log.Fatal(err)
	}
}
