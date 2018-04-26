package main

import (
	"log"

	"github.com/comail/colog"
	"github.com/e10ulen/sandbox/lib"
)

const DEBUG = false

var (
	drawCheck bool
)

func main() {
	colog.Register()
	if DEBUG != true {
		colog.SetMinLevel(colog.LWarning)
	}
	//	カードの処理から実装を始めよう
	//	山札ジェネレート
	arp := lib.SliceGenerate(13, 1)
	log.Println("d:", arp)
	//	デッキ
	deck1 := append(arp, arp...)
	deck2 := append(arp, arp...)
	deck := append(deck1, deck2...)
	log.Print("d: デッキにしました。")
	lib.SliceShuffle(deck)
	log.Print("d: ", deck)
	playerHand := make([]string, 10)
	drawCards := make([]string, 10)
	drawCards, deck = stand(deck[:], 0)
	playerHand[0] = drawCards[0]
	drawCards, deck = stand(deck[:], 1)
	playerHand[1] = drawCards[0]
	log.Println("w: 二回ドロー結果")
	log.Println("d: ", deck)
	log.Print("w: ", playerHand)
	p := playerHand[0] + playerHand[1]
	log.Print("w: ", p)
	checkTime := gameCheck(playerHand[:])
	if checkTime == true {
		drawCards, deck = stand(deck[:], 2)
		playerHand[2] = drawCards[0]
		log.Print("w: ", playerHand)
	} else {
		battle()
	}
}
func stand(twodraw []string, i int) ([]string, []string) {
	user := draw(twodraw[:], i)

	lib.SliceUnset(twodraw[:], i)
	log.Println("d: ", twodraw)
	return user, twodraw
}
func gameCheck(hand []string) bool {
	gamemsg := "ドローする[y/n]\n"
	log.Println("w: 現在の手札です。どうしますか")
	log.Println("w: ", hand)
	check := lib.ScanLine(gamemsg)
	if check == "y" {
		drawCheck = true
		log.Print("w: ", drawCheck)
		return drawCheck

	}
	return false
}
func draw(cards []string, no int) []string {
	log.Println("i: ドローしてみます")
	user := make([]string, 10)
	user[0] = cards[no]
	log.Println("i: ", user)
	return user
}
func battle() {
}
