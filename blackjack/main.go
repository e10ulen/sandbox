package main

import (
	"log"

	"github.com/comail/colog"
	"github.com/e10ulen/sandbox/lib"
)

const DEBUG = true

func init() {
	colog.Register()
	if DEBUG != true {
		colog.SetMinLevel(colog.LWarning)
	}
	//	カードの処理から実装を始めよう
	//	山札ジェネレート
	arp := lib.SliceGenerate(13, 1)
	arx := lib.SliceGenerate(13, 1)
	ary := lib.SliceGenerate(13, 1)
	arz := lib.SliceGenerate(13, 1)
	log.Print("w: シャッフル前に山札を開示します。")
	log.Println("w:", arp)
	log.Println("w:", arx)
	log.Println("w:", ary)
	log.Println("w:", arz)

	//dealler := make([]string, 5)
	//	デッキ
	deck1 := append(arp, arx...)
	deck2 := append(ary, arz...)
	deck := append(deck1, deck2...)
	log.Print("d: デッキにしました。")
	lib.SliceShuffle(deck)
	log.Print("d: ", deck)
	cLen := copy(deck1, deck)
	log.Print("w: ", cLen)
}

func main() {
}
