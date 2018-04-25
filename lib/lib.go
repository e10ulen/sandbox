package lib

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/comail/colog"
	"github.com/mattn/go-isatty"
	"github.com/russross/blackfriday"
)

func logset() {
	colog.Register()
}

//	ScanLine()
//	1行読み込みを行います。
func ScanLine(why string) string {
	fmt.Print(why)
	fmt.Print("input :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scntext := scanner.Text()
	fmt.Println("input is", scntext)
	return scntext
}

/*	IsTerminal

	Terminalかどうかを判定します

*/
func IsTerminal(check bool, cyg bool, term string) (bool, bool, string) {

	if isatty.IsTerminal(os.Stdout.Fd()) {
		term = "Terminal"
		check = true
		cyg = false
		return check, cyg, term
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		term = "Cygwin"
		check = true
		cyg = true
		return check, cyg, term
	} else {
		term = "Pipe?"
		check = false
		cyg = false
		return check, cyg, term
	}
}

func Markdown4html() {
	logset()
	gfn := "ファイル名を入力してください\n"
	md, err := ioutil.ReadFile(ScanLine("Markdown" + gfn))
	if err != nil {
		log.Print("w: ", err)
	}
	html := blackfriday.MarkdownCommon(md)
	dir, _ := os.Getwd()
	fn := ScanLine(gfn)
	file, err := os.OpenFile(dir+"/"+fn, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Print("w: ", err)
	}
	defer file.Close()

	ioutil.WriteFile(fn, html, 0666)
}

//	簡易的なサーバー機能追加
func MiniServe(port, dir string) {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Print("e:ListenAndServe: ", err)
	}
}

//	スライス生成
func SliceGenerate(n, m int) []string {
	arr := make([]string, n)
	for i := range arr {
		arr[i] = fmt.Sprintf("%2d", i+m)
	}
	return arr
}

//	スライスシャッフル
func SliceShuffle(data []string) {
	n := len(data)
	rand.Seed(time.Now().UnixNano())
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

//	スライス削除
func SliceUnset(s []string, i int) []string {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
