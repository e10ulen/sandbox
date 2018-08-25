package main

import (
	"context"
	"log"
	"strings"

	"github.com/comail/colog"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	m "github.com/mattn/go-mastodon"
	"github.com/spf13/viper"
)

func main() {
	colog.Register()
	gtk.Init(nil)
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("Mastodon TimeLine for SLCB")
	win.SetSizeRequest(400, 600)
	win.Connect("destroy", gtk.MainQuit)
	hbox := gtk.NewHBox(false, 1)
	entry := gtk.NewEntry()
	button := gtk.NewButtonWithLabel("toot")
	hbox.Add(entry)
	hbox.Add(button)

	swin := gtk.NewScrolledWindow(nil, nil)

	store := gtk.NewTreeStore(glib.G_TYPE_STRING, glib.G_TYPE_STRING)

	tree := gtk.NewTreeView()
	swin.Add(tree)

	tree.SetModel(store.ToTreeModel())
	tree.AppendColumn(gtk.NewTreeViewColumnWithAttributes("表示名", gtk.NewCellRendererText(), "text", 0))
	tree.AppendColumn(gtk.NewTreeViewColumnWithAttributes("トゥート", gtk.NewCellRendererText(), "text", 1))
	//	go-mastodon 関係
	viper.SetConfigName(".mastodon")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("w: ", err)
	}
	config := &m.Config{
		Server:       viper.GetString("server"),
		ClientID:     viper.GetString("clientid"),
		ClientSecret: viper.GetString("clientsecret"),
	}
	email := viper.GetString("emailaddress")
	pass := viper.GetString("password")

	cfg := m.NewClient(config)
	cfg.Authenticate(context.Background(), email, pass)

	//	Stream
	timeline, err := cfg.GetTimelinePublic(context.Background(), true, nil)
	if err != nil {
		log.Print("e: ", err)
	}
	for i := len(timeline) - 1; i >= 0; i-- {
		var iter gtk.TreeIter
		store.Append(&iter, nil)
		store.Set(&iter, displayname(timeline[i]), content(timeline[i]))
	}
	button.Clicked(func() {
		toot := entry.GetText()
		cfg.PostStatus(context.Background(), &m.Toot{
			Status: toot,
		})
		entry.SetText("")

		timeline, err := cfg.GetTimelinePublic(context.Background(), true, nil)
		if err != nil {
			log.Print("e: ", err)
		}
		store.Clear()
		for i := len(timeline) - 1; i >= 0; i-- {
			var iter gtk.TreeIter

			store.Append(&iter, nil)
			store.Set(&iter, displayname(timeline[i]), content(timeline[i]))
		}
	})

	vbox := gtk.NewVBox(false, 1)
	vbox.PackStart(hbox, false, false, 0)
	vbox.Add(swin)
	win.Add(vbox)
	win.ShowAll()
	gtk.Main()
}

func displayname(t *m.Status) string {
	disp := t.Account.DisplayName
	disp = strings.Replace(disp, "行城白雪/ Weiße Rosa FabriK", "行城白雪", -1)
	disp = strings.Replace(disp, "🔥藤堂傭兵🔥（避難用）", "藤堂傭兵", -1)
	disp = strings.Replace(disp, "🔥崎奈🔥@二人ぼっち時間", "崎奈", -1)
	disp = strings.Replace(disp, "🔥🔥ひと", "うらひと", -1)

	return disp
}

func content(t *m.Status) string {
	rep := t.Content
	rep = strings.Replace(rep, "<p>", "", -1)
	rep = strings.Replace(rep, "</p>", "", -1)
	rep = strings.Replace(rep, "<br />", "\n", -1)
	return rep
}
