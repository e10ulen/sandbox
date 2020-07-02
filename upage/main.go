package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/comail/colog"
)

//	スレッドURLv
var threadURL = "http://dawnlight.ovh/test/read.cgi/viptext/1520663900/"

func main() {
	//	ロガー
	colog.Register()
	res, err := http.Get(threadURL)
	//	スレッド取得に失敗した際のエラー
	if err != nil {
		log.Print("e:")
	}
	//	通常処理
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print("e: url scrapping failed", err)
	}
	//	一個ずつの投稿を取得する
	selection := doc.Find("dl.thread")
	innerselection := selection.Find("a")
	innerselection.Each(func(index int, s *goquery.Selection) {
		//	URL取得部分
		updateurl, _ := s.Attr("href")
		//	置換で間違って入ってる、mailto:sageを除去。
		replaceURL := strings.Replace(updateurl, "mailto:sage", "", -1)
		fmt.Printf("Nun:%d URL:%s\n", index, replaceURL)
	})
}
