package main

import (
	"encoding/csv"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getTotalUrlPages(website string, param string) (uint8, error) {
	var numPages uint64

	// parseamos la URL
	urlParsed, err := url.Parse(website)

	// obtenemos el valor del parámetro Nrpp
	pages := urlParsed.Query().Get(param)
	numPages, err = strconv.ParseUint(pages, 10, 32)

	// control de errores
	if err != nil {
		return 0, err
	}
	return uint8(numPages), err
}

func getUrlWithPageChanged(website string, page string) string {

	newUrl, err := url.Parse(website)

	if err != nil {
		panic(err)
	}

	urlValues := newUrl.Query()

	urlValues.Set("No", page)

	newUrl.RawQuery = urlValues.Encode()

	return newUrl.String()
}

func writeDocumentCurrentPage(writer *csv.Writer, doc *goquery.Document, category string) {

	// selecciona todos los elementos de la clase "leftList"
	doc.Find(".leftList").Each(func(i int, s *goquery.Selection) {

		name := s.Find(".descrip_full").Text()

		code := strings.Trim(s.Find(".span_codigoplu").Text(), " () ")

		price := strings.TrimSpace(s.Find(".atg_store_newPrice").Text())

		imageUrl, _ := s.Find(".atg_store_productImage img").Attr("src")

		// escribe los datos extraídos en el archivo .csv
		writer.Write([]string{name, code, price, imageUrl, category})

	})
}

func writeDocumentAllPages(writer *csv.Writer, category string) error {

	// obtiene el total de iteraciones a realizar en la url
	pages, err := getTotalUrlPages(WEBSITE, PARAM)

	for page := 0; uint8(page) <= pages; page++ {

		// obtiene la url con la página a iterar.
		WEBSITE = getUrlWithPageChanged(WEBSITE, strconv.Itoa(page))

		// parsea y guarda el contenido de la url
		pageContent, err := goquery.NewDocument(WEBSITE)
		if err != nil {
			log.Fatal(err)
		}

		// escribe en el documento los productos obtenidos
		writeDocumentCurrentPage(writer, pageContent, category)

	}

	if err != nil {
		log.Println(err)
	}
	return err
}
