package main

import (
	"encoding/csv"
	"fmt"
	"get-product-data/database"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getTotalUrlPages(website string, param string) (uint8, error) {
	var numPages uint64
	// parse the url
	urlParsed, err := url.Parse(website)

	// this get the valor of the website the param
	pages := urlParsed.Query().Get(param)
	numPages, err = strconv.ParseUint(pages, 10, 32)
	fmt.Println(pages)

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

	// select all the elements in "leftList" and iterate in those
	doc.Find(".leftList").Each(func(i int, s *goquery.Selection) {

		name := s.Find(".descrip_full").Text()

		code := strings.Trim(s.Find(".span_codigoplu").Text(), " () ")

		price := strings.TrimSpace(s.Find(".atg_store_newPrice").Text())

		imageUrl, _ := s.Find(".atg_store_productImage img").Attr("src")

		// write the extracted data in the .csv
		writer.Write([]string{name, code, price, imageUrl, category})

	})
}

func writeDocumentAllPages(writer *csv.Writer, category string) {

	// get the total of the iterations to make in the website
	pages, err := getTotalUrlPages(database.WEBSITE, database.PARAM)

	for page := 0; uint8(page) <= pages; page++ {

		// get the url ti iterate
		database.WEBSITE = getUrlWithPageChanged(database.WEBSITE, strconv.Itoa(page))

		// parse and save the content of the url
		pageContent, err := goquery.NewDocument(database.WEBSITE)
		if err != nil {
			log.Fatal(err)
		}

		// write in the document the products obtained
		writeDocumentCurrentPage(writer, pageContent, category)

	}
}
