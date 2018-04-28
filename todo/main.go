package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	done_mark1 = "\u2610"
	done_mark2 = "\u2611"
)

func main() {
	app := kingpin.New("todo", "a Todo Command.")
	listTask(app)
	addTask(app)
	doneTask(app)
	undoneTask(app)
	deleteTask(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func listTask(app *kingpin.Application) {
	cmd := app.Command("list", "List Task")
	cmd.Action(func(c *kingpin.ParseContext) error {
		file, err := os.Open("todo.txt")
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

func addTask(app *kingpin.Application) {
	cmd := app.Command("add", "add Task")
	text := cmd.Arg("text", "text to task").Strings()
	cmd.Action(func(c *kingpin.ParseContext) error {
		file, err := os.OpenFile("todo.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		defer file.Close()
		fmt.Fprintln(file, strings.Join(*text, " "))
		return nil
	})
}
func doneTask(app *kingpin.Application) {
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
		w, err := os.Create("todo.txt_")
		if err != nil {
			return err
		}
		defer w.Close()
		//	todoファイル本体の読み込み
		f, err := os.Open("todo.txt")
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

		err = os.Remove("todo.txt")
		if err != nil {
			return err
		}

		return os.Rename("todo.txt_", "todo.txt")
	})
}

func undoneTask(app *kingpin.Application) {
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
		w, err := os.Create("todo.txt_")
		if err != nil {
			return err
		}
		defer w.Close()
		//	todoファイル本体の読み込み
		f, err := os.Open("todo.txt")
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

		err = os.Remove("todo.txt")
		if err != nil {
			return err
		}

		return os.Rename("todo.txt_", "todo.txt")
	})
}

//	Delete Task List
func deleteTask(app *kingpin.Application) {
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
		w, err := os.Create("todo.txt_")
		if err != nil {
			return err
		}
		defer w.Close()
		//	todoファイル本体の読み込み
		f, err := os.Open("todo.txt")
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

		err = os.Remove("todo.txt")
		if err != nil {
			return err
		}

		return os.Rename("todo.txt_", "todo.txt")
	})
}
