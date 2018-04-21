package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//	オプション解析（デフォルトではしない）
	r.ParseForm()
	//	鯖側に出力される
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//	ここでwに入るのがクライアント側に出力される
	fmt.Fprintf(w, "hello hello")
}

func main() {
	//	アクセスルーティング
	http.HandleFunc("/", sayhelloName)
	//	監視するポート
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
