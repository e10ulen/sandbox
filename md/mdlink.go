package main

import (
	"github.com/e10ulen/sandbox/lib"
	"github.com/spiegel-im-spiegel/logf"
	"github.com/spiegel-im-spiegel/mklink"
)

func main() {
	text := "URLを入力してください\n"
	url := lib.ScanLine(text)
	link, err := mklink.New(url)
	if err != nil {
		logf.Warn(err)
		return
	}
	logf.Println(link.Encode(mklink.StyleMarkdown))
}
