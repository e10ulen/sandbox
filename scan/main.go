package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/spiegel-im-spiegel/logf"
)

func main() {
	logf.Println("Scanner使って一行読み込み")
	fmt.Print("input ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("input is", scanner.Text())
}
