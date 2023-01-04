package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var WEBSITE = "https://www.cotodigital3.com.ar/sitios/cdigi/browse/catalogo-almac%C3%A9n/_/N-8pub5z?Nf=product.startDate%7CLTEQ+1.6726176E12%7C%7Cproduct.endDate%7CGTEQ+1.672704E12%7C%7Cproduct.endDate%7CGTEQ+1.6726176E12%7C%7Cproduct.startDate%7CLTEQ+1.672704E12&No=0&Nr=AND%28product.sDisp_200%3A1004%2Cproduct.language%3Aespa%C3%B1ol%2COR%28product.siteId%3ACotoDigital%29%29&Nrpp=48#"

const PARAM = "Nrpp"

func main() {

	// abre el archivo .csv para escribir en él
	file, err := os.Create("productos_almacen.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// crea un nuevo escritor de CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// escribe la cabecera del archivo .csv
	writer.Write([]string{"name", "code", "price", "image_url"})

	pages, err := getTotalUrlPages(WEBSITE, PARAM)

	for page := 0; uint8(page) <= pages; page++ {

		// obtiene la url con la nueva página
		WEBSITE = getUrlWithPageChanged(WEBSITE, strconv.Itoa(page))

		// realiza una solicitud HTTP GET a la página web obtenida
		doc, err := goquery.NewDocument(WEBSITE)
		if err != nil {
			log.Fatal(err)
		}

		// escribe en el documento los productos obtenidos
		writeDocument(writer, doc, pages)

	}

}
