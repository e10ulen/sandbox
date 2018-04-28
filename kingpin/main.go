package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/comail/colog"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	yaml "gopkg.in/yaml.v2"
)

type Todo struct {
	Id    int
	Text  string
	Check bool
}

type Config struct {
	Day  string
	Todo []Todo
}

func main() {
	colog.Register()
	log.Print("w: ymlを読み込む")
	buf, err := ioutil.ReadFile("todo.yml")
	if err != nil {
		log.Print("w: ", err)
	}
	t := Config{}
	err = yaml.Unmarshal(buf, &t)
	if err != nil {
		log.Print("w: ", err)
	}
	fmt.Println("%v", t)

	app := kingpin.New("todo", "a Todo Command.")
	addTask(app)
	//listTask(app)
	//deleteTask(app)
	//doneTask(app)
	//undoneTask(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func addTask(app *kingpin.Application) {
	cmd := app.Command("add", "add Tasklist.")
	text := cmd.Arg("text", "text to Tasklist.").Strings()
	day := time.Now().Format("06/01/02")
	cmd.Action(func(c *kingpin.ParseContext) error {
		text := strings.Join(*text, " ")
		println("day:", day)
		println("todo:", text)
		return nil
	})
}
