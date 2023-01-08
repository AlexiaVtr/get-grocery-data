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

	// this open the .csv and write in it
	file, err := os.Create(fmt.Sprintf(`C:/Temp/%s.csv`, database.CATEGORY))
	defer file.Close()

	// create a new csv writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write the header of the .csv document
	writer.Write([]string{"name", "code", "price", "image_url", "category"})

	//write all obtained data of the website
	writeDocumentAllPages(writer, database.CATEGORY)

	if err != nil {
		log.Fatal(err)
	}
}
