package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/comail/colog"
)

func main() {
	colog.Register()
	res, err := goquery.NewDocument("http://shindanmaker.com/a/67048")
	if err != nil {
		log.Print("e: ", err)
	}
	selection := res.Find("div#main2")
	html := selection.Text()
	log.Print("w: ", html)

}
