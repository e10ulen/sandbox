# これなに
[e10ulen][1]の習作プログラム集です。  
各フォルダで使ってるライブラリが違ったり、手動テストしてたり、後は自作プログラムのライブラリだったりします。  

以下は各フォルダについての説明です。  

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
  
### rot13フォルダ
なんとなく、kingpin使って、暗号化と復号化が実装できたらなと挑戦した。  
バグってます。  


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
済:現状はまだバグっているので対処する。  
[e10ulen/todo](https://github.com/e10ulen/todo)に移動しました。  
  

## msdフォルダ
- [e10ulen/sandbox](https://github.com/e10ulen/sandbox/)  
- [alecthomas/kingpin][11]  
- [mattn/go-mastodon][13]  
自作アプリ、todoのソースを見ながらkingpinを使い、mastodonに接続するツールです。  
現在、タイムライン（ローカルタイムライン）とトゥートに対応しています。  
時期にローカルタイムラインとホームタイムラインをコマンドで分け実装します。  
なおかつ、ローカルタイムラインはストリームに対応します。  
一応すでに使える状態であります。
msdフォルダにあるexample.yamlを書き換えてHOME直下に.mastodon.yamlとして置いてください。  
[e10ulen/msd](https://github.com/e10ulen/msd)に移動しました。  
  
## msd2
- [e10ulen/msd](https://github.com/e10ulen/msd)
- [mattn/go-mastodon][13]
- [mattn/go-gtk](https://github.com/mattn/go-gtk)

自作アプリ、[e10ulen/msd](https://github.com/e10ulen/msd) を元にmattnさんのgo-gtkでGUIを付加したものになります。  
現在、ローカルタイムラインを描画します（２０件）。  
[e10ulen/msd](https://github.com/msd)と同じ設定ファイルを使用するようになってます。  

以下のdisplayname内に `` disp = strings.Replace(disp, "FooBarBaz", "FooBaz", -1) `` を記載すれば置換されます。  
表示名そのままを描画すると長すぎるのに全部描画しようとするので自らの調整が必要になります。  

さらにその下にあるcontentはTootを取得した際、<p></p>タグや<br />タグを処理するための箇所です。  
お好みで書き換えてくださいますよう  

``

func func displayname(t *m.Status) string {
	disp := t.Account.DisplayName

	return disp
}

``

## upageフォルダ
-[PuerkitoBio/goquery][14]
goqueryを使い、避難所からスクレイピングしようとしています。

## ggフォルダ
- [alecthomas/kingpin.v2][11]  
git commit を適宣楽にしようとして作ってるgolang書き始めた時から色々書き換えたりしてるツールです。  
[e10ulen/gg](https://github.com/e10ulen/gg)に移動しました。

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
[13]:https://github.com/mattn/go-mastodon
[14]:https://github.com/PuerkitoBio/goquery