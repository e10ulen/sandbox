package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	done_mark1    = "\u2610"
	done_mark2    = "\u2611"
	todo_filename = ".todo"
)

func main() {
	app := kingpin.New("todo", "a Todo Command.")
	//	OS	振り分け後、ファイルパス類を各コマンドに渡す。
	filename := ""
	existCurTodo := false
	curDir, err := homedir.Dir()
	if err == nil {
		filename = filepath.Join(curDir, todo_filename)
		_, err = os.Stat(filename)
		if err != nil {
			existCurTodo = true
		}
	}
	if !existCurTodo {
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		filename = filepath.Join(home, todo_filename)
	}
	listTask(app, filename)
	addTask(app, filename)
	doneTask(app, filename)
	undoneTask(app, filename)
	deleteTask(app, filename)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func listTask(app *kingpin.Application, filename string) {
	cmd := app.Command("list", "List Task")
	cmd.Action(func(c *kingpin.ParseContext) error {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		buf := bufio.NewReader(file)
		number := 1
		for {
			b, _, err := buf.ReadLine()
			if err != nil {
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

func addTask(app *kingpin.Application, filename string) {
	cmd := app.Command("add", "add Task")
	text := cmd.Arg("text", "text to task").Strings()
	cmd.Action(func(c *kingpin.ParseContext) error {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		defer file.Close()
		fmt.Fprintln(file, strings.Join(*text, " "))
		return nil
	})
}
func doneTask(app *kingpin.Application, filename string) {
	cmd := app.Command("done", "Done Task")
	text := cmd.Arg("number", "done task").Strings()
	cmd.Action(func(c *kingpin.ParseContext) error {

		ids := []int{}

		for _, no := range *text {
			id, err := strconv.Atoi(no)
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}

		//	仮のファイル作成（テンポラリファイル）
		w, err := os.Create(filename + "_")
		if err != nil {
			return err
		}
		defer w.Close()
		//	todoファイル本体の読み込み
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		buf := bufio.NewReader(f)
		number := 1
		for {
			buf, _, err := buf.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			match := false
			for _, id := range ids {
				if id == number {
					match = true
				}
			}
			line := strings.TrimSpace(string(buf))
			if match && !strings.HasPrefix(line, "-") {
				_, err = fmt.Fprintf(w, "-%s\n", line)
				if err != nil {
					return err
				}
			} else {
				_, err = fmt.Fprintf(w, "%s\n", line)
				if err != nil {
					return err
				}
			}
			number++
		}
		f.Close()
		w.Close()

		err = os.Remove(filename)
		if err != nil {
			return err
		}

		return os.Rename(filename+"_", filename)
	})
}

func undoneTask(app *kingpin.Application, filename string) {
	cmd := app.Command("undone", "Undone Task")
	text := cmd.Arg("number", "undone task").Strings()
	cmd.Action(func(c *kingpin.ParseContext) error {

		ids := []int{}

		for _, no := range *text {
			id, err := strconv.Atoi(no)
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}

		//	仮のファイル作成（テンポラリファイル）
		w, err := os.Create(filename + "_")
		if err != nil {
			return err
		}
		defer w.Close()
		//	todoファイル本体の読み込み
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		buf := bufio.NewReader(f)
		number := 1
		for {
			buf, _, err := buf.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			match := false
			for _, id := range ids {
				if id == number {
					match = true
				}
			}
			line := strings.TrimSpace(string(buf))
			if match && strings.HasPrefix(line, "-") {
				_, err = fmt.Fprintf(w, "%s\n", string(line[1:]))
				if err != nil {
					return err
				}
			} else {
				_, err = fmt.Fprintf(w, "%s\n", line)
				if err != nil {
					return err
				}
			}
			number++
		}
		f.Close()
		w.Close()

		err = os.Remove(filename)
		if err != nil {
			return err
		}

		return os.Rename(filename+"_", filename)
	})
}

//	Delete Task List
func deleteTask(app *kingpin.Application, filename string) {
	cmd := app.Command("delete", "Delete Task")
	text := cmd.Arg("number", "delete task").Strings()
	cmd.Action(func(c *kingpin.ParseContext) error {

		ids := []int{}

		for _, no := range *text {
			id, err := strconv.Atoi(no)
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}

		//	仮のファイル作成（テンポラリファイル）
		w, err := os.Create(filename + "_")
		if err != nil {
			return err
		}
		defer w.Close()
		//	todoファイル本体の読み込み
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		buf := bufio.NewReader(f)
		number := 1
		for {
			buf, _, err := buf.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			match := false
			for _, id := range ids {
				if id == number {
					match = true
				}
			}
			if !match {
				_, err = fmt.Fprintf(w, "%s\n", string(buf))
				if err != nil {
					return err
				}
			}
			number++
		}
		f.Close()
		w.Close()

		err = os.Remove(filename)
		if err != nil {
			return err
		}

		return os.Rename(filename+"_", filename)
	})
}
