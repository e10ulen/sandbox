package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fakeStatusChange()
}

func fakeStatusChange() {
	sentence := 50
	countup := 0
	var delay time.Duration
	lineout := 0
	rand.Seed(time.Now().UnixNano())
	for {
		if sentence != countup {
			fmt.Printf("#")
		} else {
			//	50文字になったら「done + 改行を出力する。
			//	何かが終わった感じを出力する。
			fmt.Printf(" done!!\n")
			//	countupをリセット
			countup = -1
			//	無限ループを脱出するためにカウントアップしておく。
			lineout++
		}
		//	カウントしていく
		countup++
		//	高速に出力するとそれっぽくないので、waitをかけて遅延をする。
		//	ランダム値にマッチすると待ち時間を更に遅くさせる。
		//	delayを作る。
		delay = time.Duration(100) * time.Millisecond
		if 5 < countup && countup <= 10*lineout {
			time.Sleep(delay)
		}
		//	10行書かれたらbreakして無限ループを終わらせる。
		if lineout == 10 {
			time.Sleep(1000)
			break
		}
	}
}
