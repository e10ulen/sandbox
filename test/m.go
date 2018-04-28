package main

import (
	"log"

	"github.com/comail/colog"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	colog.Register()
	home, _ := homedir.Dir()
	log.Print("w: ", home)
}
