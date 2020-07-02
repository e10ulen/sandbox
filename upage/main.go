package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/comail/colog"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-sqlite3"
)

func execDB(db *sql.DB, q string) {
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err)
	}
}

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
	//	DataBase接続設定
	db, err := sql.Open("sqlite3", "./test.db")
	//	DB接続時エラー
	if err != nil {
		log.Fatal(err)
	}
	//	DB作成
	q := `
		CREATE TABLE update (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			body VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))
		)
	`
	//	DB呼び出し
	execDB(db, q)

	//	DBテスト書き込み
	q = `
        INSERT INTO update (body)
        VALUES ('body1'), ('body2')
    `
	//	DBテスト書き込みのための呼び出し
	execDB(db, q)
	//	DBを閉じる
	db.Close()

	//	通常処理
	defer res.Body.Close()
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
		fmt.Printf("Nun:%d URL:%s\n", index, replaceURL)
	})
}
