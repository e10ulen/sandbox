package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/comail/colog"
)

func main() {
	colog.Register()
	res, err := http.Get("http://shindanmaker.com/a/67048")
	if err != nil {
		log.Print("e: ", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Print("e: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print("e: %s", err)
	}
	doc.Find(".shindanform")
}
