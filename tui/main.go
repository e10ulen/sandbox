package main

import (
	"log"

	"github.com/comail/colog"
	tui "github.com/marcusolsson/tui-go"
)

var (
	player   = "vipper"
	playerLv = "1"
	playerHP = "100"
	playerMP = "50"
	pATK     = "10"
	pDEF     = "10"
	pINT     = "10"
	pAGI     = "10"
	pVIT     = "10"
	pLUC     = "10"
)

type getText struct {
	label *tui.Label
}

func main() {
	colog.Register()
	menuBar := tui.NewVBox()
	menuBar.SetBorder(true)
	addMenu()
	menuBar.Append(tui.NewLabel("test"))
	mainBar := tui.NewVBox()
	mainBar.Append(tui.NewLabel("test"))
	mainBar.SetBorder(true)
	root := tui.NewHBox(menuBar, mainBar)
	ui, err := tui.New(root)
	if err != nil {
		log.Print("w:", err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	if err := ui.Run(); err != nil {
		log.Print("e: RunError ", err)
	}
}
func addMenu(*getText) {
	menuBar.Append(label("addMenu"))
}
