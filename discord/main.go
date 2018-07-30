package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if m.Content == "お腹空いた" || m.Content == "おなかすいた" {
		s.ChannelMessageSend(m.ChannelID, "ご飯食べた？")
	}
	// If the message is "pong" reply with "Ping!"
	if m.Content == "村焼場" {
		s.ChannelMessageSend(m.ChannelID, "https://slcb.m.to/")
	}
	if m.Content == "いちくら" {
		s.ChannelMessageSend(m.ChannelID, "https://ichiji.social")
	}
	if m.Content == "幼女さん" || m.Content == "ロリ" || m.Content == "しょた" {
		s.ChannelMessageSend(m.ChannelID, "可愛い")
	}
	if m.Content == "ショタ" {
		s.ChannelMessageSend(m.ChannelID, "かっこいい")
	}
	if m.Content == "alhe" || m.Content == "あるひ" || m.Content == "アルヒ" {
		s.ChannelMessageSend(m.ChannelID, "活字ジャンキー")
	}
	if m.Content == "木村" {
		s.ChannelMessageSend(m.ChannelID, "ネットマナーを学んでから一昨日きやがれってんだ")
	}
	if m.Content == "崎奈さん" {
		s.ChannelMessageSend(m.ChannelID, "放課後スリーフィンガー")
	}
	if m.Content == "🔥" {
		s.ChannelMessageSend(m.ChannelID, "いいんですね？燃やしますよ！！！！！\n")
		time.Sleep(500 * time.Millisecond)
		s.ChannelMessageSend(m.ChannelID, "🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥")
	}
	if m.Content == "cmd" {
		s.ChannelMessageSend(m.ChannelID, "🔥/alhe/崎奈さん/インスタンス/いちくら/ping/お腹空いた/おなかすいた\n")
		s.ChannelMessageSend(m.ChannelID, "幼女さん/ロリ/ショタ/しょた/あるひ/アルヒ/")
	}
	if m.Content == "sena" {
		s.ChannelMessageSend(m.ChannelID, "http://slcb.xyz/images/gallery/sena_megane_no.kizunelaurant.png")
	}
	if m.Content == "はまちさん" {
		s.ChannelMessageSend(m.ChannelID, "バズりストはまち")
	}
	if m.Content == "ブレーメン" {
		s.ChannelMessageSend(m.ChannelID, "🐔")
		s.ChannelMessageSend(m.ChannelID, "🐈 ")
		s.ChannelMessageSend(m.ChannelID, "🐩")
		s.ChannelMessageSend(m.ChannelID, "🐎")
	}
}
