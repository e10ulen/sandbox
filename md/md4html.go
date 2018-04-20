package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	md, err := ioutil.ReadFile("readme.md")
	if err != nil {
		fmt.Println(err)
	}
	html := blackfriday.MarkdownBasic(md)
	dir, _ := os.Getwd()
	file, err := os.OpenFile(dir+"/readme.html", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	ioutil.WriteFile("readme.html", html, 0666)
}
