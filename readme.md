# これなに
[e10ulen][1]の習作プログラム集です。  
各フォルダで使ってるライブラリが違ったり、手動テストしてたり、後は自作プログラムのライブラリだったりします。  

以下は各フォルダについての説明と、現在一番まともに改修を繰り返しているtodoアプリの説明です。  

### blackjackフォルダ
現在[プログラミング入門者からの卒業試験は『ブラックジャック』を開発すべし](https://bit.ly/2HtrQiC)を読み、golangで実装できないかとしています  
(items以下のURLが長いので直打ちするの怖い(pushするときだけテザリングでネット繋いでる))  
### kingpinフォルダ
todoアプリの前段階、色々とテストした場所。  
yamlを読んでみたり、jsonを読んでみたりした。  
結局todoは標準ライブラリを使ってファイル読み書きすることにした。  

### tuiフォルダ
[marcusolsson/tui-go][9]を使ってテキストユーザインタフェースの実装をしてみたかった  

### mdフォルダ
- [russross/blackfriday][7]
- [spiegel-im-spiegel/mklink][8]
  
以上二点を使ったテスト、blackfridayの方はlibフォルダにてMarkdown4html関数に使っています。  
  
### servフォルダ
特筆すべきことはないが、net/httpのサーバ立ち上げについてのテスト。  
lib.MiniServe関数に組み込まれているのはこれ  

### termフォルダ
[mattn/go-isatty][2]
### scrapフォルダ
- [PuerkitoBio/goquery][5]
- [nakabonne/netsurfer][6]
  
以上二点の使い方テスト？を行ったフォルダ  
  
### scanフォルダ
特筆すべきことはないが、bufioとかの使い方をテストしたかっただけ、  
lib.ScanLineに組み込まれているので必要なし。 

### testフォルダ
libの関数追加や、libの関数をすぐ使いたい時に  
  
### homedirフォルダ
[mitchellh/go-homedir][12]を使って[mattn/todo][3]の処理だとバグる  
`` os.Getwd() ``の代わりに使用するためのテスト  
なんかググってもまともに情報出てこないし、  
とりあえず、valueがいくつひつようなのかとかのテストしたフォルダ
## libフォルダ
lib.xxxxという形でライブラリとし、  

- 簡易サーバ  
- Markdownをhtml化  
- 標準入力の簡素化  
- スライスの生成  
- スライスシャッフル  
- スライスの削除  

以上六点を提供する自作ライブラリである。  


## todoフォルダ
- [mattn/todo][3]  
- [chooyan/todo][4]  

以上二点を参考にkingpinで作成したtodoアプリである。  
現状はまだバグっているので対処する。  
## よく使うライブラリ
- [comail/colog][10]  
- [alecthomas/kingpin.v2][11]  
  
[1]:https://github.com/e10ulen
[2]:https://github.com/mattn/go-isatty
[3]:https://github.com/mattn/todo
[4]:https://github.com/chooyan/todo
[5]:https://github.com/PuerkitoBio/goquery
[6]:https://github.com/nakabonne/netsurfer
[7]:https://github.com/russross/blackfriday
[8]:https://github.com/spiegel-im-spiegel/mklink
[9]:https://github.com/marcusolsson/tui-go
[10]:https://github.com/comail/colog
[11]:https://gopkg.in/alecthomas/kingpin.v2
[12]:https://github.com/mitchellh/go-homedir
