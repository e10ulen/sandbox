package main

import (
	"time"
	"context"
	"fmt"
	"github.com/spf13/viper"
	m "github.com/mattn/go-mastodon"
	"github.com/spiegel-im-spiegel/logf"
	"github.com/marcusolsson/tui-go"
)

type post struct {
	username string
	message string
	time string
}

var posts = []post{
	{username: "e10ulen", message: "ここにタイムライン描画", time:"15:15"},
	{username: "e10ulen", message: "終了するときはESCキー", time:"15:15"},
}

func main(){
	//	コンフィグをロードする
	viper.SetConfigName(".qqw")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logf.Warn(err)
	}
	config := &m.Config{
		Server: viper.GetString("server"),
		ClientID: viper.GetString("clientid"),
		ClientSecret: viper.GetString("clientsecret"),
	}
	email := viper.GetString("email")
	pass := viper.GetString("pass")
	c := m.NewClient(config)
	c.Authenticate(context.Background(), email, pass)
	his := tui.NewVBox()

	for _, m := range posts {
		his.Append(tui.NewHBox(
			tui.NewLabel(m.time),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", m.username))),
			tui.NewLabel(m.message),
			tui.NewSpacer(),

		))
	}
	hisScroll := tui.NewScrollArea(his)
	hisScroll.SetAutoscrollToBottom(true)

	hisBox := tui.NewVBox(hisScroll)
	hisBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(hisBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	//	ここからマストドン要素入れていく？？？？
	input.OnSubmit(func(e *tui.Entry){
		his.Append(tui.NewHBox(
			tui.NewLabel(time.Now().Format("15:04")),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", "e10ulen"))),
			tui.NewLabel(e.Text()),
			//
			logf.Debugf("ここまできてる")
			c.PostStatus(context.Background(), &m.Toot{
				Status: e.Text(),
			}),
			//
			tui.NewSpacer(),
		))
		input.SetText("")
	})

	ui, err := tui.New(chat)
	if err != nil {
		logf.Warn(err)
	}

	ui.SetKeybinding("Esc", func(){ ui.Quit() })
	if err := ui.Run(); err != nil {
		logf.Warn(err)
	}
}
