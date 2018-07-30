package main

import (
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(nil) //	無いとPanicを起こす。
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("test")
	win.SetSizeRequest(400, 300)
	win.Connect("destroy", gtk.MainQuit)
	lab := gtk.NewLabel("このツール（ｒｙ")
	win.Add(lab)
	win.ShowAll()
	gtk.Main()
}
