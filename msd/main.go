package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/comail/colog"
	"github.com/fatih/color"
	m "github.com/mattn/go-mastodon"
	"github.com/spf13/viper"
	"golang.org/x/net/html"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	colog.Register()
	app := kingpin.New("md", "a Mastodon Application")
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
	timelineMastodon(app, cfg)
	tootMastodon(app, cfg)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func tootMastodon(app *kingpin.Application, cfg *m.Client) {
	cmd := app.Command("toot", "toot to mastodon")
	text := cmd.Arg("text", "text to toot").Strings()
	cmd.Action(func(c *kingpin.ParseContext) error {
		toot := strings.Join(*text, " ")
		cfg.PostStatus(context.Background(), &m.Toot{
			Status: toot,
		})
		return nil
	})
}
func timelineMastodon(app *kingpin.Application, cfg *m.Client) {
	cmd := app.Command("tl", "TimeLine for mastodon")
	cmd.Action(func(c *kingpin.ParseContext) error {
		timeline, err := cfg.GetTimelinePublic(context.Background(), true, nil)
		if err != nil {
			return err
		}
		for i := len(timeline) - 1; i >= 0; i-- {
			displayStatus(timeline[i])
		}
		return nil
	})
}

func acct(a string) string {
	return a
}
func displayStatus(t *m.Status) {
	if t == nil {
		return
	}
	if t.Reblog != nil {
		color.Set(color.FgHiRed)
		fmt.Printf(acct(t.Account.Acct))
		color.Set(color.Reset)
		fmt.Printf(" reblogged ")
		color.Set(color.FgHiBlue)
		fmt.Println(acct(t.Reblog.Account.Acct))
		fmt.Println(textContent(t.Reblog.Content))
		color.Set(color.Reset)
	} else {
		color.Set(color.FgHiRed)
		fmt.Printf(acct(t.Account.Acct))
		color.Set(color.Reset)
		color.Set(color.FgHiGreen)
		fmt.Printf(acct(t.Account.DisplayName))
		color.Set(color.Reset)
		fmt.Println(textContent(t.Content))
	}
}

func textContent(s string) string {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return s
	}
	var buf bytes.Buffer

	var extractText func(node *html.Node, w *bytes.Buffer)
	extractText = func(node *html.Node, w *bytes.Buffer) {
		if node.Type == html.TextNode {
			data := strings.Trim(node.Data, "\r\n")
			if data != "" {
				w.WriteString(data)
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			extractText(c, w)
		}
		if node.Type == html.ElementNode {
			name := strings.ToLower(node.Data)
			if name == "br" {
				w.WriteString("\n")
			}
		}
	}
	extractText(doc, &buf)
	return buf.String()
}
