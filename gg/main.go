package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/e10ulen/sandbox/lib"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	date_format = "2006/01/02 15:04"
)

// ggCmd represents the gg command
func main() {
	app := kingpin.New("gg", "a git commit & push")
	autoCommit(app) //	git add
	getMessage(app) //	git commit -m
	pushRemote(app) //	git push -u
	getCommit(app)  //	git log --date=short --no-merges --pretty=format:"%cd (@%cn) %h %s"
	fullCommand(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
func getCommit(app *kingpin.Application) {
	cmd := app.Command("list", "CommitList")
	cmd.Action(func(c *kingpin.ParseContext) error {
		list, _ := exec.Command("git", "log", "--date=short", "--no-merges", "--pretty=format:%cd (@%cn) %h %s").CombinedOutput()
		fmt.Print(string(list))
		return nil
	})
}

func autoCommit(app *kingpin.Application) {
	cmd := app.Command("a", "Commit")
	cmd.Action(func(c *kingpin.ParseContext) error {
		add, _ := exec.Command("git", "add", "--all").CombinedOutput()
		log.Print(string(add))
		return nil
	})
}
func getMessage(app *kingpin.Application) {
	cmd := app.Command("c", "Get CommitMessage")
	cmd.Action(func(c *kingpin.ParseContext) error {
		tm := time.Now().Format(date_format)
		get := "コミットメッセージを入力してください\n"
		cmt, _ := exec.Command("git", "commit", "-m", "[Commit]"+tm+" "+lib.ScanLine(get)).CombinedOutput()
		log.Print(string(cmt))
		return nil
	})
}

func pushRemote(app *kingpin.Application) {
	cmd := app.Command("p", "Push Remote")
	cmd.Action(func(c *kingpin.ParseContext) error {
		push, _ := exec.Command("git", "push", "-u").CombinedOutput()
		log.Print(string(push))
		return nil
	})
}

func fullCommand(app *kingpin.Application) {
	cmd := app.Command("f", "add &commit &push to one command.")
	text := cmd.Arg("text", "text to commit message").Strings()
	cmd.Action(func(c *kingpin.ParseContext) error {
		tm := time.Now().Format(date_format)
		addf, _ := exec.Command("git", "add", "--all").CombinedOutput()
		msg := strings.Join(*text, " ")
		cmmt, _ := exec.Command("git", "commit", "-m", "[Commit]"+tm+" "+msg).CombinedOutput()
		push, _ := exec.Command("git", "push", "-u").CombinedOutput()
		log.Print(string(addf), string(cmmt), string(push))
		return nil
	})
}
