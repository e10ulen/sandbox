# これなに
golang スクラップ（書きかけだったり失敗作置き場
一応libフォルダには使えるスクラップが突っ込まれている。
バージョンについてはなんとなく記載したかった。
尚、ゴールはない。

## libのバージョンについて
1.4

## testフォルダについて
libに関数追加するときに大体の動きを把握するために
試し呼びをするためのフォルダ

## mastodonフォルダについて
tui使って、[qqw](https://github.com/e10ulen/qqw)からtootとtimeline見る部分切り出して、  
ちょっとGUIっぽくしたかった（  

## mdフォルダについて
[mklink](https://github.com/spiegel-im-spiegel/mklink)を使って自分なりに使えたらなと思った失敗作、  
近日中には書き直す。  
### mdフォルダ更新
[blackfriday](https://github.com/russross/blackfriday)のv1を使って、htmlへ変換して、ファイルとして保存することに成功したので、  
libとして使えるように書いた  
内部的にlibのScanLineを使うので、ScanLineにも改修を施した。

## termフォルダについて
[go-isatty](https://github.com/mattn/go-isatty)を使ってターミナルかどうかの判定テストしてみたかった。

## scanフォルダについて
bufio.Scannerの一行読み込みを使いたかった。

## libフォルダについて
色々と自分で使い勝手のいい関数を作るためのフォルダ。  
termフォルダの判定とかも組み込む予定。  
mdフォルダよりMarkdownファイルをhtmlに変換する簡易的な物を実装  
ScanLine()関数を改修。  
ScanLine(args)と引数を要求するようにし、  
引数には適正な物、例えば今回の実装で使っている場面として、
このような形にしている。  
servフォルダより簡易サーバ実装しました。  

``

	gfn := "ファイル名を入力してください\n"
	fn := ScanLine(gfn)
	file, err := os.OpenFile(dir+"/"+fn, ...
    ...
    	defer file.Close()

	ioutil.WriteFile(fn, html, 0666)
``

## servフォルダについて
小さな簡易サーバ機能付加の練習

##  使い方一覧
argsについてはlibフォルダについてを参照 

##  blackjackフォルダについて
現在[プログラミング入門者からの卒業試験は『ブラックジャック』を開発すべし](https://bit.ly/2HtrQiC)を読み、golangで実装できないかとしています  
(items以下のURLが長いので直打ちするの怖い(pushするときだけテザリングでネット繋いでる))  


``  
package main  
  
import (  
    "os"
    "fmt"  
    "github.com/e10ulen/sandbox/lib"  
)  
  
func main(){  
    fmt.Println("実装済みの物の使い方？")  
    fmt.Println("Scanner使った一行読み込み")  
    lib.ScanLine(args)  
    fmt.Println("Terminalかの判定処理")  
    lib.IsTerminal()  
    fmt.Println("Markdown4html")  
    lib.Markdown4html()  
    fmt.Println("サーバー起動します")  
    port := "9090"  
    dir, _ := os.Getwd()  
    lib.MiniServe(port, dir)  
}  
``

