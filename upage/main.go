package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/comail/colog"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-sqlite3"
)

//	スレッドURLv
var threadURL = "http://dawnlight.ovh/test/read.cgi/viptext/1520663900/"

//	DBPath
const dbPath = "./update.db"

//	データ格納
//	|書き込み番号|書き込まれたURL|
type Update struct {
	number int
	url    string
}

func main() {
	//	ロガー
	colog.Register()
	//	データベースOpen
	DBConnnect, _ := sql.Open("sqlite3", dbPath)
	//	データベースClose
	defer DBConnnect.Close()
	//	スレッドテーブル作成
	if file, err := os.Stat(dbPath); os.IsNotExist(err) || file.IsDir() {
		cmd := `CREATE TABLE IF NOT EXISTS "threadUP"(
			"number" INTEGER PRIMARY KEY,
			"url" VARCHAR(255))`
		//	cmdを実行する。リターンは受け取らない。
		_, err := DBConnnect.Exec(cmd)
		//	データベース用エラーハンドリング
		if err != nil {
			log.Fatal(err)
		}
	}

	res, err := http.Get(threadURL)
	//	スレッド取得に失敗した際のエラー
	if err != nil {
		log.Print("e:")
	}
	log.Print("d: %v", time.Now())
	//	通常処理
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print("e: url scrapping failed", err)
	}
	//	データベース挿入処理
	cmd := "UPDATE threadUP ( number, url) VALUES(?,?)"
	//	一個ずつの投稿を取得する
	selection := doc.Find("dl.thread")
	innerselection := selection.Find("a")
	innerselection.Each(func(index int, s *goquery.Selection) {
		//	URL取得部分
		updateurl, _ := s.Attr("href")

		//	置換で間違って入ってる、mailto:sageを除去。
		replaceURL := strings.Replace(updateurl, "mailto:sage", "", -1)
		fmt.Printf("Nun:%d URL:%s\n", index, replaceURL)
		_, err = DBConnnect.Exec(cmd, index, replaceURL)
		if err != nil {
			log.Fatal(err)
		}

	})
}
