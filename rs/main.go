package main

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/comail/colog"
	"github.com/nfnt/resize"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	colog.Register()
	app := kingpin.New("res", "resize image")
	jpg_resize(app)
	png_resize(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func png_resize(app *kingpin.Application) {
	cmd := app.Command("png", "画像のリサイズを行います。\nwidthかheightの数値を0にすると縦横比を維持したままリサイズを行います。")
	fwidth := cmd.Arg("width", "リサイズ後の幅").Required().Int()
	fheight := cmd.Arg("height", "リサイズ後の高さ").Required().Int()
	ifile := cmd.Arg("png", "ファイル名を拡張子抜きで入力してください。").Strings()
	cmd.Action(func(c *kingpin.ParseContext) error {
		intHeight := uint(*fheight)
		intWidth := uint(*fwidth)
		files := strings.Join(*ifile, " ")
		file, err := os.Open(files + ".png")
		if err != nil {
			log.Print("e: ", err)
		}
		img, err := png.Decode(file)
		if err != nil {
			log.Print("e: ", err)
		}
		file.Close()
		m := resize.Resize(intWidth, intHeight, img, resize.Lanczos3)
		out, err := os.Create(files + ".test.png")
		if err != nil {
			log.Print("e: ", err)
		}
		defer file.Close()
		png.Encode(out, m)
		return nil
	})
}

func jpg_resize(app *kingpin.Application) {
	cmd := app.Command("jpg", "画像のリサイズを行います。widthかheightの数値を0にすると縦横比を維持したままリサイズを行います。")
	fwidth := cmd.Arg("width", "リサイズ後の幅").Required().Int()
	fheight := cmd.Arg("height", "リサイズ後の高さ").Required().Int()
	ifile := cmd.Arg("jpg", "ファイル名を拡張子抜きで入力してください。").Strings()
	cmd.Action(func(c *kingpin.ParseContext) error {
		intHeight := uint(*fheight)
		intWidth := uint(*fwidth)
		files := strings.Join(*ifile, " ")
		file, err := os.Open(files + ".jpg")
		if err != nil {
			log.Print("e: ", err)
		}
		//	decode jpg
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Print("e: ", err)
		}
		file.Close()
		m := resize.Resize(intWidth, intHeight, img, resize.Lanczos3)
		out, err := os.Create(files + ".test.jpg")
		if err != nil {
			log.Print("w: ", err)
		}
		defer file.Close()
		jpeg.Encode(out, m, nil)
		return nil
	})
}
