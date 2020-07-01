package main

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
}
