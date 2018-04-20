package main

import (
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	// アプリケーションの定義
	app := kingpin.New("chat", "A command-line chat application.")
	// register コマンドの定義
	register(app)
	// post コマンドの定義
	post(app)
	// コマンドライン引数のパースとコールバックの実行
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func register(app *kingpin.Application) {
	// コマンドの定義
	cmd := app.Command("register", "Register a new user.")
	// 引数の定義. 変更前は変数名が重複しないように "registerNick" のように接頭語を入れていたがローカル変数なのでそれも不要になった
	nick := cmd.Arg("nick", "Nickname for user.").Required().String()
	name := cmd.Arg("name", "Name of user.").Required().String()
	// コールバックの定義
	cmd.Action(func(c *kingpin.ParseContext) error {
		// 引数を使って処理を実行
		// *kingpin.ParseContext にはコマンド定義や引数定義も含まれているが簡便に取り出す手段はたぶん提供されていないので先に引数を定義してローカル変数に入れてから参照している
		println(*nick, *name)
		return nil
	})
}

func post(app *kingpin.Application) {
	// コマンドの定義
	cmd := app.Command("post", "Post a message to a channel.")
	// フラグの定義
	image := cmd.Flag("image", "Image to post.").File()
	// 引数の定義
	channel := cmd.Arg("channel", "Channel to post to.").Required().String()
	text := cmd.Arg("text", "Text to post.").Strings()
	// コールバックの定義
	cmd.Action(func(c *kingpin.ParseContext) error {
		text := strings.Join(*text, " ")
		println("Post:", text)
		println("Image:", *image)
		println("Channel:", *channel)
		return nil
	})
}
