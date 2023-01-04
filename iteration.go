package main

import (
	"encoding/csv"
	"net/url"
	"strconv"

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

func writeDocument(writer *csv.Writer, doc *goquery.Document, numPages uint8) {

	// selecciona todos los elementos de la clase "leftList"
	doc.Find(".leftList").Each(func(i int, s *goquery.Selection) {

		name := s.Find(".descrip_full").Text()

		code := s.Find(".span_codigoplu").Text()

		price := s.Find(".atg_store_newPrice").Text()

		imageUrl, _ := s.Find(".atg_store_productImage img").Attr("src")

		// escribe los datos extraídos en el archivo .csv
		writer.Write([]string{name, code, price, imageUrl})

	})
}
