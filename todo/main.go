package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/comail/colog"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	done_mark1 = "\u2610"
	done_mark2 = "\u2611"
)

func main() {
	colog.Register()

	app := kingpin.New("todo", "a Todo Command.")
	listTask(app)
	addTask(app)
	doneTask(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func listTask(app *kingpin.Application) {
	cmd := app.Command("list", "List Task")
	cmd.Action(func(c *kingpin.ParseContext) error {
		file, err := os.Open("todo.txt")
		if err != nil {
			log.Print("w: ", err)
		}
		defer file.Close()
		buf := bufio.NewReader(file)
		number := 1
		for {
			b, _, err := buf.ReadLine()
			if err != nil {
				log.Print("w: ", err)
				if err != io.EOF {
					return err
				}
				break
			}
			line := string(b)
			if strings.HasPrefix(line, "-") {
				fmt.Printf("%s %03d: %s\n", done_mark2, number, strings.TrimSpace(string(line[1:])))
			} else {
				fmt.Printf("%s %03d: %s\n", done_mark1, number, strings.TrimSpace(line))
			}
			number++
		}
		return nil
	})
}

func addTask(app *kingpin.Application) {
	cmd := app.Command("add", "add Task")
	text := cmd.Arg("text", "text to task").Strings()
	cmd.Action(func(c *kingpin.ParseContext) error {
		file, err := os.OpenFile("todo.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			log.Print("w: ", err)
		}
		defer file.Close()
		fmt.Fprintln(file, strings.Join(*text, "\n"))
		return nil
	})
}

func doneTask(app *kingpin.Application) {
	cmd := app.Command("done", "Done Task")
	text := cmd.Arg("number", "done task").Int()
	cmd.Action(func(c *kingpin.ParseContext) error {
		log.Print("w: ", *text)
		return nil
	})
}
