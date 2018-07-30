package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/e10ulen/sandbox/lib"
	"github.com/fatih/color"
)

/*

["あ","い","う","え","お","か","き","く","け","こ","さ","し","す","せ","そ","た","ち","つ","て","と","な","に","ぬ","ね","の","は","ひ","ふ","へ","ほ","ま","み","む","め","も","や","ゆ","よ","ら","り","る","れ","ろ","わ","ん","きゃ","きゅ","きょ","しゃ","しゅ","しょ","ちゃ","ちゅ","ちょ","にゃ","にゅ","にょ","ひゃ","ひゅ","ひょ","みゃ","みゅ","みょ","りゃ","りゅ","りょ","ぎゃ","ぎゅ","ぎょ","じゃ","じゅ","じょ","びゃ","びゅ","びょ","ぴゃ","ぴゅ","ぴょ","が","ぎ","ぐ","げ","ご","ざ","じ","ず","ぜ","ぞ","だ","ぢ","づ","で","ど","ば","び","ぶ","べ","ぼ","ぱ","ぴ","ぷ","ぺ","ぽ",]

*/

func main() {
	arr := []string{"あ", "い", "う", "え", "お", "か", "き", "く", "け", "こ", "さ", "し", "す", "せ", "そ", "た", "ち", "つ", "て", "と", "な", "に", "ぬ", "ね", "の", "は", "ひ", "ふ", "へ", "ほ", "ま", "み", "む", "め", "も", "や", "ゆ", "よ", "ら", "り", "る", "れ", "ろ", "わ", "ん", "きゃ", "きゅ", "きょ", "しゃ", "しゅ", "しょ", "ちゃ", "ちゅ", "ちょ", "にゃ", "にゅ", "にょ", "ひゃ", "ひゅ", "ひょ", "みゃ", "みゅ", "みょ", "りゃ", "りゅ", "りょ", "ぎゃ", "ぎゅ", "ぎょ", "じゃ", "じゅ", "じょ", "びゃ", "びゅ", "びょ", "ぴゃ", "ぴゅ", "ぴょ", "が", "ぎ", "ぐ", "げ", "ご", "ざ", "じ", "ず", "ぜ", "ぞ", "だ", "ぢ", "づ", "で", "ど", "ば", "び", "ぶ", "べ", "ぼ", "ぱ", "ぴ", "ぷ", "ぺ", "ぽ"}
	arr2 := arr
	arr3 := arr
	arr4 := arr
	arr5 := arr
	arr6 := arr
	arr7 := arr
	arr8 := arr
	arr9 := arr
	lib.SliceShuffle(arr2)
	shuffle(arr3)
	lib.SliceShuffle(arr4)
	lib.SliceShuffle(arr5)
	lib.SliceShuffle(arr6)
	lib.SliceShuffle(arr7)
	lib.SliceShuffle(arr8)
	lib.SliceShuffle(arr9)
	lib.SliceShuffle(arr)
	arr = append(arr, arr2...)
	arr = append(arr, arr3...)
	arr = append(arr, arr4...)
	arr = append(arr, arr5...)
	arr = append(arr, arr6...)
	arr = append(arr, arr7...)
	arr = append(arr, arr8...)
	arr = append(arr, arr9...)
	shuffle(arr)
	res1 := strings.Join(arr, "")
	clear := color.Red("ぷみゃみ")

	fmt.Println(strings.Replace(res1, "ぷみゃみ", clear, 0))
}
func shuffle(data []string) {
	n := len(data)
	rand.Seed(time.Now().UnixNano() * time.Now().UnixNano())
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}
