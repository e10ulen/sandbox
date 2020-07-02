package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/comail/colog"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//	ロガー
	colog.Register()
	//	スレッドURL指定
	url := "http://dawnlight.ovh/test/read.cgi/viptext/1520663900/"
	res, err := http.Get(url)
	//	スレッド取得に失敗した際のエラー
	if err != nil {
		log.Print("e:")
	}
	defer res.Body.Close()
	var sliceURL []string
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print("e: url scrapping failed", err)
	}
	//一個ずつの投稿を取得する
	selection := doc.Find("dl.thread")
	innerselection := selection.Find("a")
	innerselection.Each(func(index int, s *goquery.Selection) {
		//	URL取得部分
		url, _ := s.Attr("href")
		//	置換で間違って入ってる、mailto:sageを除去。
		replaceURL := strings.Replace(url, "mailto:sage", "", -1)

		sliceURL = append(sliceURL, replaceURL)
		fmt.Printf("Nun:%d URL:%s\n", index, replaceURL)

	})
	fmt.Printf("Slice:%v", len(sliceURL))
}
