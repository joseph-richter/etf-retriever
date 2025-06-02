package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func webContentGet(link string) float32 {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("Error Status Code : %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var yieldToWorst float32

	if strings.Contains(link, "blackrock") {
		doc.Find(".product-data-item.col-yieldToWorst").Each(func(i int, s *goquery.Selection) {
			valueNotFormatted := strings.Fields(s.Text())

			valueString := strings.Replace(valueNotFormatted[6], "%", "", -1)
			valueString = strings.Replace(valueString, ",", ".", -1)

			value, err := strconv.ParseFloat(valueString, 32)
			if err != nil {
				log.Fatal(err)
			}
			yieldToWorst = float32(value)
		})
	}
	return yieldToWorst
}
