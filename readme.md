# これなに
golang スクラップ（書きかけだったり失敗作置き場

## mastodonフォルダについて
tui使って、[qqw](https://github.com/e10ulen/qqw)からtootとtimeline見る部分切り出して、  
ちょっとGUIっぽくしたかった（  

## mdフォルダについて
[mklink](https://github.com/spiegel-im-spiegel/mklink)を使って自分なりに使えたらなと思った失敗作、  
近日中には書き直す。  
### mdフォルダ更新
[blackfriday](https://github.com/russross/blackfriday)のv1を使って、htmlへ変換して、ファイルとして保存することに成功したので、  
libとして使えるように書く  


## termフォルダについて
[go-isatty](https://github.com/mattn/go-isatty)を使ってターミナルかどうかの判定テストしてみたかった。

## scanフォルダについて
bufio.Scannerの一行読み込みを使いたかった。

## libフォルダについて
色々と自分で使い勝手のいい関数を作るためのフォルダ。
termフォルダの判定とかも組み込む予定。

``
package main

import (
    "fmt"
    "github.com/e10ulen/sandbox/lib"
)

func main(){
    fmt.Println("実装済みの物の使い方？")
    fmt.Println("Scanner使った一行読み込み")
    lib.ScanLine()
    fmt.Println("Terminalかの判定処理")
    lib.IsTerminal()
}
``

