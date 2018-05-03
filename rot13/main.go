package main

import (
	"bytes"
	"io"
	"os"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type rot13Reader struct {
	r bytes.Reader
}

//	復号
func rot13(c byte) byte {
	switch {
	case ('A' <= c && c <= 'Z'):
		return (c-'A'+13)%26 + 'A'
	case ('a' <= c && c <= 'z'):
		return (c-'a'+13)%26 + 'a'
	default:
		return c
	}
}
func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return 0, err
	}
	for i := range p {
		p[i] = rot13(p[i])
	}
	return
}

func main() {
	app := kingpin.New("rot13", "ROT13")
	encodeRot13(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func encodeRot13(app *kingpin.Application) {
	cmd := app.Command("e", "Encode Rot13")
	text := cmd.Arg("text", "rot13").Strings()
	cmd.Action(func(c *kingpin.ParseContext) error {
		rot := strings.Join(*text, " ")
		cod := rot13Reader{rot}
		io.Copy(os.Stdout, &cod)
		return nil
	})
}
