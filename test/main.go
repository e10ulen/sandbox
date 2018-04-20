package main

import (
    "fmt"
    "github.com/e10ulen/sandbox/lib"
)

func main(){
    fmt.Println("実装済みの物の使い方？")
    fmt.Println("Scanner使った一行読み込み")
    fmt.Println(lib.ScanLine())
    fmt.Println("Terminalかの判定処理")
	var foo bool
	var bar bool
	var baz string
	fmt.Println(lib.IsTerminal(foo, bar, baz))
}
