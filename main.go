package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// realiza una solicitud HTTP GET a la página web
	doc, err := goquery.NewDocument("https://www.cotodigital3.com.ar/sitios/cdigi/browse/catalogo-almac%C3%A9n/")
	if err != nil {
		log.Fatal(err)
	}

	// abre el archivo .csv para escribir en él
	f, err := os.Create("products.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// crea un nuevo escritor de CSV
	w := csv.NewWriter(f)
	defer w.Flush()

	// escribe la cabecera del archivo .csv
	w.Write([]string{"name", "code", "price", "image_url"})

	// selecciona todos los elementos de la clase "product-description"
	doc.Find(".leftList").Each(func(i int, s *goquery.Selection) {
		// extrae el nombre del producto
		name := s.Find(".descrip_full").Text()

		// extrae el código del producto
		code := s.Find(".span_codigoplu").Text()

		// extrae el precio del producto
		price := s.Find(".atg_store_newPrice").Text()

		// extrae la url de la imagen del producto
		imageUrl, _ := s.Find(".atg_store_productImage img").Attr("src")

		// escribe los datos extraídos en el archivo .csv
		w.Write([]string{name, code, price, imageUrl})

	})

}
