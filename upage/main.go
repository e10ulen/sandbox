package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/japanese"

	"github.com/comail/colog"
	"golang.org/x/text/transform"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//doc, err := goquery.NewDocument("http://dawnlight.ovh/test/read.cgi/viptext/1520663900/")
	//if err != nil {
	//	fmt.Print("url scrapping failed")
	//}
	//	URL取得部分
	//doc.Find("a").Each(func(_ int, s *goquery.Selection) {
	//	url, _ := s.Attr("href")
	//	fmt.Println(url)
	//})

	//一個ずつの投稿を取得する
	//selection := doc.Find("dl.thread")
	//innerselection := selection.Find("a")
	//innerselection.Each(func(index int, s *goquery.Selection) {
	// URLを取得する
	//	url, _ := s.Attr("href")
	//	fmt.Println(url)
	//})

	//File書き込み
	//doc.Find("a").Each(func(_ int, s *goquery.Selection) {
	//	url, _ := s.Attr("href")
	//	content := []byte(
	//	url
	//	)
	//})
	//ioutil.WriteFile("./log.txt", content, os.ModePerm)

	//	***
	//	Product
	//	***
	/*
		//	Get Request
		urls := "http://dawnlight.ovh/test/read.cgi/viptext/1520663900/"
		res, _ := http.Get(urls)
		defer res.Body.Close()

		//	Load Request
		buf, _ := ioutil.ReadAll(res.Body)

		//	Charset Request
		det := chardet.NewTextDetector()
		detRslt, _ := det.DetectBest(buf)
		fmt.Println(detRslt.Charset)
		bReader := bytes.NewReader(buf)
		reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

		//	HTML Parser
		docs, _ := goquery.NewDocumentFromReader(reader)

		//	Title Open
		rslt := docs.Find("title").Text()
		fmt.Println(rslt)
	*/
	///*/
	//	Requests
	//*/
	//	url := "http://dawnlight.ovh/test/read.cgi/viptext/1520663900/"
	//	res, _ := http.Get(url)
	//	defer res.Body.Close()
	//	buf, _ := ioutil.ReadAll(res.Body)
	//	det := chardet.NewTextDetector()
	//	detRslt, _ := det.DetectBest(buf)
	//	fmt.Println(detRslt.Charset)
	//	bReader := bytes.NewReader(buf)
	//	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)
	//	doc, _ := goquery.NewDocumentFromReader(reader)
	//	rslt := doc.Find("dl.thread")
	//	fmt.Println(rslt)
	colog.Register()
	url := "http://dawnlight.ovh/test/read.cgi/viptext/1520663900/"
	res, err := http.Get(url)
	if err != nil {
		log.Print("e:")
	}
	defer res.Body.Close()
	utfBody := transform.NewReader(bufio.NewReader(res.Body), japanese.ShiftJIS.NewDecoder())

	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		fmt.Print("url scrapping failed")
	}
	//一個ずつの投稿を取得する
	selection := doc.Find("dl.thread")
	innerselection := selection.Find("a")
	innerselection.Each(func(index int, s *goquery.Selection) {
		dd := s.Find("dd").Nodes
		// URLを取得する
		//	URL取得部分

		url, _ := s.Attr("href")
		//	置換で間違って入ってる、mailto:sageを除去。
		replaceURL := strings.Replace(url, "mailto:sage", "", -1)

		fmt.Println(replaceURL)
		fmt.Println(dd)

	})
}
