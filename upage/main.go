package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/comail/colog"
)

//	スレッドURLv
var threadURL = "http://dawnlight.ovh/test/read.cgi/viptext/1520663900/"

func main() {
	//	ロガー
	colog.Register()
	//	Call Of & Write & Read
	writeAge()
	readAge()
	//readAgeTitle()
}

//	書き込み処理
func writeAge() {
	file, err := os.OpenFile("thread.md", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Print("e:", err)
	}
	//
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
	defer file.Close()
	//	一個ずつの投稿を取得する
	selection := doc.Find("dl.thread")
	innerselection := selection.Find("a")
	innerselection.Each(func(_ int, s *goquery.Selection) {

		//	URL取得部分
		updateurl, _ := s.Attr("href")

		//	置換で間違って入ってる、mailto:sageを除去。
		replaceURL3 := strings.Replace(updateurl, "mailto:sage", "", -1)
		replaceURL2 := strings.Replace(replaceURL3, "http://tinyscenery.in/test/read.cgi/bbs/1491121427/", "", -1)
		replaceURL1 := strings.Replace(replaceURL2, "http://example.com/", "", -1)
		replaceURL := strings.TrimSpace(replaceURL1)
		file.WriteString(replaceURL + "\n")

	})
}

//	読み込み処理
func readAge() {
	const sleepInterval = 10
	data, err := os.Open("thread.md")
	if err != nil {
		log.Print("e: ", err)
		return
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		logging := fmt.Sprint(scanner.Text())
		fmt.Println(logging)
		log.Print("d: Debug")
		res, err := http.Get(logging)
		if err != nil {
			log.Print("e: ", err)
		}
		defer res.Body.Close()
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Print("e: ", err)
		}
		log.Print("d: ", doc)
	}
}
