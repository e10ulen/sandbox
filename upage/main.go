package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://dawnlight.ovh/test/read.cgi/viptext/1520663900/")
	if err != nil {
		fmt.Print("url scrapping failed")
	}
	//	URL取得部分
	//doc.Find("a").Each(func(_ int, s *goquery.Selection) {
	//	url, _ := s.Attr("href")
	//	fmt.Println(url)
	//})

	//一個ずつの投稿を取得する
	selection := doc.Find("dl.thread")
	innerselection := selection.Find("a")
	innerselection.Each(func(index int, s *goquery.Selection) {
		// URLを取得する
		url, _ := s.Attr("href")

		fmt.Println(url)
	})

	//File書き込み
	//doc.Find("a").Each(func(_ int, s *goquery.Selection) {
	//	url, _ := s.Attr("href")
	//	content := []byte(
	//	url
	//	)
	//})
	//ioutil.WriteFile("./log.txt", content, os.ModePerm)
}
