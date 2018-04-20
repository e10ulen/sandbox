package lib

import (
	"fmt"
	"bufio"
	"os"
	"github.com/spiegel-im-spiegel/logf"
	"github.com/mattn/go-isatty"
)
//	ScanLine()
//	1行読み込みを行います。
func ScanLine() string {
	logf.Println("Scanner使って一行読み込み")
	fmt.Print("input ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scntext := scanner.Text()
	fmt.Println("input is", scntext)
	return scntext
}
/*	IsTerminal
	
	Terminalかどうかを判定します
	
*/
func IsTerminal(check bool, cyg bool, term string) (bool , bool, string) {

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
