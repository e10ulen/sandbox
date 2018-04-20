package main

import (
	"github.com/spiegel-im-spiegel/logf"
	"github.com/spiegel-im-spiegel/mklink"
	"github.com/atotto/clipboard"
)
func main() {
	text, _  := clipboard.ReadAll()
	link, err := mklink.New(text)
	if err != nil {
	    logf.Warn(err)
		clipboard.WriteAll("")
		return
	}
	logf.Println(link.Encode(mklink.StyleMarkdown))
}

