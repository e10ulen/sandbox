package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/comail/colog"
)

func main() {
	colog.Register()
	res, err := http.Get("http://localhost:1313")
	if err != nil {
		log.Print("e: %s", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Print("e: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print("e: %s", err)
	}
	doc.Find(".c-heading__title__text").Each(func(i int, s *goquery.Selection) {
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s\n - %s", i, band, title)
	})
}
