package main

import (
	"fmt"
	"os"

	"github.com/e10ulen/sandbox/lib"
)

func main() {
	fmt.Println("実装済みの物のテスト")
	//fmt.Println("Scanner使った一行読み込み")
	//fmt.Println(lib.ScanLine())

	//fmt.Println("Terminalかの判定処理")
	//var foo bool
	//var bar bool
	//var baz string
	//fmt.Println(lib.IsTerminal(foo, bar, baz))
	//fmt.Println("md4html")
	lib.Markdown4html()
	port := "9090"
	dir, _ := os.Getwd()
	lib.MiniServe(port, dir)
}
