package upage

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/comail/colog"
)

func Upage() {
	colog.Register()
	data, err := ioutil.ReadFile("thread.md")
	if err != nil {
		log.Print("e: ", err)
		return
	}
	fmt.Println(string(data))
}
