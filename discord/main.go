package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/comail/colog"
)

type config struct {
	Token   string `json:'Token'`
	BotName string `json:'BotName'`
}

var (
	StopBot = make(chan bool)

	msg = "!helloworld"
)

func loadConfig() (*config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Print("e: loadConfig os.Open err:", err)
		return nil, err
	}
	defer file.Close()
	var cfg config
	err = json.NewDecoder(file).Decode(&cfg)
	return &cfg, err
}

func main() {
	colog.Register()
	cfg, err := loadConfig()
	if err != nil {
		log.Println("e: Config Load miss")
	}
	dis, err := discordgo.New()
	dis.Token = cfg.Token
	if err != nil {
		log.Println("e: %d", err)
	}

	dis.AddHandler(onMessageCreate)
	//	Websocketを開いてlistening開始
	err = dis.Open()
	if err != nil {
		log.Print("e: %d", err)
		log.Print("えらーかしょ")
	}
	defer dis.Close()

	log.Println("d: Listening...")
	<-StopBot
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	fmt.Printf("%5s %5s %5s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

}

func sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)

	log.Print("d: >>>" + msg)
	if err != nil {
		log.Print("e: %d", err)
	}
}
